package runtime

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Output JSON structure for runtime output
type Output struct {
	Type       string  `json:"type"`
	VarName    string  `json:"var_name"`
	MapKeyType string  `json:"map_key_type"`
	Val        string  `json:"val"`
	ArrIdent   string  `json:"arr_ident"`
	Child      *Output `json:"child"`
}

// OutputParser parses runtime output strings
type OutputParser struct {
	mem *[]string
}

// NewOutputParser creates a new output paraser
func NewOutputParser() *OutputParser {
	return &OutputParser{
		mem: &[]string{},
	}
}

// Parse parses printed runtime output to statements
func (o *OutputParser) Parse(printed, funcName string, index int) ([]Stmt, bool) {
	result := []Stmt{}
	// Regex output for current organism
	re := regexp.MustCompile(fmt.Sprintf(`%s\n((.*)\n)*%s`, StartName(funcName, index), EndName(funcName, index)))
	curFuncOutput := re.FindString(printed)
	// Check if function paniced
	if strings.Contains(curFuncOutput, fmt.Sprintf("Recovered in Test%s%d", funcName, index)) {
		return result, true
	}
	// Otherwise split lines
	lines := strings.Split(curFuncOutput, "\n")
	for _, line := range lines {
		// Check if line starts with expected JSON
		if strings.HasPrefix(line, `{ "type":`) {
			// Create assert statements from JSON line
			result = append(result, o.ParseLine(line)...)
		}
	}
	return result, false
}

// ParseLine parses a line of output
func (o *OutputParser) ParseLine(jsonString string) []Stmt {
	data := Output{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.WithError(err).WithField("line", jsonString).Errorf("unable to parse runtime output")
		return []Stmt{}
	}
	return o.AssertStmts(&data, []Replacement{}, TypeCorrections{}, []Stmt{})
}

// Replacement struct which contains keys and replacement values
type Replacement struct {
	Key string
	Val string
}

// TypeCorrections struct containing information for printing corrections in types
// e.g. custom values X(int(3)) in order to make correct assert statements
type TypeCorrections struct {
	Prefix string
	Suffix string
}

// AssertStmts new assert statement from runtime output
func (o *OutputParser) AssertStmts(data *Output, replacements []Replacement, typeCorrection TypeCorrections, resStmts []Stmt) []Stmt {
	if data.Child == nil {
		return o.CreateAssertStmts(data, replacements, typeCorrection, resStmts)
	}

	switch data.Type {
	case "arr":
		// In case we have an arr type we need to replace the key of the loop at runtime with the value of the loop identifier
		replacements = append(replacements, Replacement{
			Key: data.ArrIdent,
			Val: data.Val,
		})
		return o.AssertStmts(data.Child, replacements, typeCorrection, resStmts)
	case "custom":
		// Ignore custom values for now, since we test on equal values
		return o.AssertStmts(data.Child, replacements, typeCorrection, resStmts)
	case "map":
		// In case of map, we need to replace the map key in the loop with the value of the map key at runtime in the identifier
		x := data.Val
		if data.MapKeyType == "string" {
			x = fmt.Sprintf("\"%s\"", x)
		}
		replacements = append(replacements, Replacement{
			Key: data.ArrIdent,
			Val: x,
		})
		return o.AssertStmts(data.Child, replacements, typeCorrection, resStmts)
	case "pointer":
		// In case pointer we check if value nil, if not
		// we need to use the value of the identifier using the start operator
		if data.Val != "nil" {
			// sanity check
			if data.Child == nil {
				log.Warningf("unable to create assert stmts, expected pointer to have child")
				return []Stmt{}
			}
			pointerStmt := fmt.Sprintf("%s := *%s", data.Child.VarName, data.VarName)
			if !Contains(*o.mem, pointerStmt) {
				*o.mem = append(*o.mem, fmt.Sprintf("%s :=", data.Child.VarName))
				resStmts = append(resStmts, &AssignStmt{
					AssignStmtType: AssignStmtTypeDefine,
					LeftHand:       data.Child.VarName,
					RightHand:      "*" + data.VarName,
				})
			} else {
				resStmts = append(resStmts, &AssignStmt{
					AssignStmtType: AssignSTmtTypeAssign,
					LeftHand:       data.Child.VarName,
					RightHand:      "*" + data.VarName,
				})
			}
		}
		// If nil just continue recursion
		return o.AssertStmts(data.Child, replacements, typeCorrection, resStmts)
	default:
		return o.AssertStmts(data.Child, replacements, typeCorrection, resStmts)
	}
}

// CreateAssertStmts creates the eventual assert statement based on runtime output and corrections
func (o *OutputParser) CreateAssertStmts(runtimeOutput *Output, replacements []Replacement, typeCorrection TypeCorrections, resStmts []Stmt) []Stmt {
	for i := 0; i < len(replacements); i++ {
		runtimeOutput.VarName = strings.ReplaceAll(runtimeOutput.VarName, replacements[i].Key, replacements[i].Val)
	}
	for j := 0; j < len(resStmts); j++ {
		for i := 0; i < len(replacements); i++ {
			resStmts[j].Replace(replacements[i].Key, replacements[i].Val)
		}
	}
	switch runtimeOutput.Type {
	case "int",
		"float32",
		"float64",
		"byte",
		"rune",
		"uintptr",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"int8",
		"int16",
		"int32",
		"int64":
		return append(resStmts, &AssertStmt{
			AssertStmtType: AssertStmtTypeEqualValues,
			Expected:       fmt.Sprintf("%s%s(%s)%s", typeCorrection.Prefix, runtimeOutput.Type, runtimeOutput.Val, typeCorrection.Suffix),
			Value:          runtimeOutput.VarName,
		})
	case "string":
		return append(resStmts, &AssertStmt{
			AssertStmtType: AssertStmtTypeEqualValues,
			Expected:       fmt.Sprintf("%s%s(`%s`)%s", typeCorrection.Prefix, runtimeOutput.Type, runtimeOutput.Val, typeCorrection.Suffix),
			Value:          runtimeOutput.VarName,
		})
	case "bool":
		if runtimeOutput.Val == "true" {
			return append(resStmts, &AssertStmt{
				AssertStmtType: AssertStmtTypeTrue,
				Expected:       runtimeOutput.VarName,
			})
		}
		return append(resStmts, &AssertStmt{
			AssertStmtType: AssertStmtTypeFalse,
			Expected:       runtimeOutput.VarName,
		})
	case "complex64", "complex128":
		return append(resStmts, &AssertStmt{
			AssertStmtType: AssertStmtTypeEqualValues,
			Expected:       fmt.Sprintf("%s%s%s%s", typeCorrection.Prefix, runtimeOutput.Type, runtimeOutput.Val, typeCorrection.Suffix),
			Value:          runtimeOutput.VarName,
		})
	// Only nil pointers will reach this point
	case "pointer":
		return append(resStmts, &AssertStmt{
			AssertStmtType: AssertStmtTypeNil,
			Expected:       runtimeOutput.VarName,
		})
	case "error":
		if runtimeOutput.Val == "nil" {
			return append(resStmts, &AssertStmt{
				AssertStmtType: AssertStmtTypeNoError,
				Expected:       runtimeOutput.VarName,
			})
		}
		return append(resStmts, &AssertStmt{
			AssertStmtType: AssertStmtTypeError,
			Expected:       runtimeOutput.VarName,
		})
	default:
		log.Warningf("unknown type: %s, value: %s", runtimeOutput.Type, runtimeOutput.Val)
		return []Stmt{}
	}
}
