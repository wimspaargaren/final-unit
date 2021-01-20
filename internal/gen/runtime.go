package gen

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

// UpdateAssertStmts sets an os assert statements based on printed runtime result
func (o *Organism) UpdateAssertStmts(printed string, firstRun bool) {
	for _, f := range o.Files {
		for funcName, testCases := range f.TestCases {
			for i, testCase := range testCases {
				mem := []string{}
				// Regex output for current organism
				re := regexp.MustCompile(fmt.Sprintf(`%s\n((.*)\n)*%s`, StartName(funcName, i), EndName(funcName, i)))
				curFuncOutput := re.FindString(printed)
				// Check if function paniced
				if strings.Contains(curFuncOutput, fmt.Sprintf("Recovered in Test%s%d", funcName, i)) {
					testCase.RunTimeInfo.Panics = true
					continue
				}
				// Otherwise split lines
				lines := strings.Split(curFuncOutput, "\n")
				for _, line := range lines {
					// Check if line starts with expected JSON
					if strings.HasPrefix(line, `{ "type":`) {
						// Create assert statements from JSON line
						assertStmts := ParseLine(line, &mem)
						if firstRun {
							testCase.RunTimeInfo.AssertStmts = append(testCase.RunTimeInfo.AssertStmts, assertStmts...)
						} else {
							testCase.RunTimeInfo.SecondRun = append(testCase.RunTimeInfo.SecondRun, assertStmts...)
						}
					}
				}
				if !firstRun {
					testCase.RunTimeInfo.SetIsValid()
				}
			}
		}
	}
}

// RuntimeOutput JSON structure for runtime output
type RuntimeOutput struct {
	Type       string         `json:"type"`
	VarName    string         `json:"var_name"`
	MapKeyType string         `json:"map_key_type"`
	Val        string         `json:"val"`
	ArrIdent   string         `json:"arr_ident"`
	Child      *RuntimeOutput `json:"child"`
}

// ParseLine parses a line of output
func ParseLine(jsonString string, mem *[]string) []string {
	data := RuntimeOutput{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.WithError(err).WithField("line", jsonString).Errorf("unable to parse runtime output")
	}
	return AssertStmts(&data, []Replacement{}, TypeCorrections{}, []string{}, mem)
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
func AssertStmts(data *RuntimeOutput, replacements []Replacement, typeCorrection TypeCorrections, resStmts []string, mem *[]string) []string {
	if data.Child == nil {
		return CreateAssertStmts(data, replacements, typeCorrection, resStmts)
	}

	switch data.Type {
	case "arr":
		// In case we have an arr type we need to replace the key of the loop at runtime with the value of the loop identifier
		replacements = append(replacements, Replacement{
			Key: data.ArrIdent,
			Val: data.Val,
		})
		return AssertStmts(data.Child, replacements, typeCorrection, resStmts, mem)
	case "custom":
		// Ignore custom values for now, since we test on equal values
		return AssertStmts(data.Child, replacements, typeCorrection, resStmts, mem)
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
		return AssertStmts(data.Child, replacements, typeCorrection, resStmts, mem)
	case "pointer":
		// In case pointer we check if value nil, if not
		// we need to use the value of the identifier using the start operator
		if data.Val != "nil" {
			// sanity check
			if data.Child == nil {
				log.Warningf("unable to create assert stmts, expected pointer to have child")
				return []string{}
			}
			pointerStmt := fmt.Sprintf("%s := *%s", data.Child.VarName, data.VarName)
			if !Contains(*mem, pointerStmt) {
				*mem = append(*mem, fmt.Sprintf("%s :=", data.Child.VarName))
				resStmts = append(resStmts, fmt.Sprintf("%s := *%s", data.Child.VarName, data.VarName))
			} else {
				resStmts = append(resStmts, fmt.Sprintf("%s = *%s", data.Child.VarName, data.VarName))
			}
		}
		// If nil just continue recursion
		return AssertStmts(data.Child, replacements, typeCorrection, resStmts, mem)
	default:
		return AssertStmts(data.Child, replacements, typeCorrection, resStmts, mem)
	}
}

// Contains check if memory contains key
func Contains(mem []string, key string) bool {
	for _, x := range mem {
		if strings.HasPrefix(key, x) {
			return true
		}
	}
	return false
}

// CreateAssertStmts creates the eventual assert statement based on runtime output and corrections
func CreateAssertStmts(x *RuntimeOutput, replacements []Replacement, typeCorrection TypeCorrections, resStmts []string) []string {
	for i := 0; i < len(replacements); i++ {
		x.VarName = strings.ReplaceAll(x.VarName, replacements[i].Key, replacements[i].Val)
	}
	for j := 0; j < len(resStmts); j++ {
		for i := 0; i < len(replacements); i++ {
			resStmts[j] = strings.ReplaceAll(resStmts[j], replacements[i].Key, replacements[i].Val)
		}
	}
	switch x.Type {
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
		return append(resStmts, fmt.Sprintf("s.EqualValues(%s%s(%s)%s,%s)", typeCorrection.Prefix, x.Type, x.Val, typeCorrection.Suffix, x.VarName))
	case "string":
		return append(resStmts, fmt.Sprintf("s.EqualValues(%s%s(`%s`)%s,%s)", typeCorrection.Prefix, x.Type, x.Val, typeCorrection.Suffix, x.VarName))
	case "bool":
		if x.Val == "true" {
			return append(resStmts, fmt.Sprintf("s.True(%s)", x.VarName))
		}
		return append(resStmts, fmt.Sprintf("s.False(%s)", x.VarName))
	case "complex64", "complex128":
		return append(resStmts, fmt.Sprintf("s.EqualValues(%s%s%s%s,%s)", typeCorrection.Prefix, x.Type, x.Val, typeCorrection.Suffix, x.VarName))
	// Only nil pointers will reach this point
	case "pointer":
		return append(resStmts, fmt.Sprintf("s.Nil(%s)", x.VarName))
	case "error":
		if x.Val == "nil" {
			return append(resStmts, fmt.Sprintf("s.NoError(%s)", x.VarName))
		}
		return append(resStmts, fmt.Sprintf("s.Error(%s)", x.VarName))
	default:
		log.Warningf("unknown type: %s, value: %s", x.Type, x.Val)
		return []string{}
	}
}

// StartName expected start tag for test case output
func StartName(funcName string, index int) string {
	return fmt.Sprintf("<START;%s%d>", funcName, index)
}

// EndName expected end tag for test case output
func EndName(funcName string, index int) string {
	return fmt.Sprintf("<END;%s%d>", funcName, index)
}
