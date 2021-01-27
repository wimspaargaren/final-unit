// Package runtime analyses runtime output and converts it into assert statements
package runtime

import (
	"fmt"
	"regexp"
	"strings"
)

// Info information about values on runtime
type Info struct {
	IsValid     bool
	Panics      bool
	AssertStmts []string
	SecondRun   []string
	Printer     StmtPrinter
}

// NewInfo creates new runtime info for given printer
func NewInfo(printer StmtPrinter) *Info {
	return &Info{
		Printer: printer,
	}
}

// SetIsValid verifies that created runtime info is valid
// used when generating end result
func (info *Info) SetIsValid() bool {
	if len(info.AssertStmts) != len(info.SecondRun) {
		info.IsValid = false
		return false
	}

	for i := 0; i < len(info.AssertStmts); i++ {
		if info.AssertStmts[i] != info.SecondRun[i] {
			info.IsValid = false
			return false
		}
	}
	info.IsValid = true
	return true
}

// AssertStmtsForTestCase creates assert statements for a testcase
func (info *Info) AssertStmtsForTestCase(printed string, firstRun bool, funcName string, index int) {
	mem := []string{}
	// Regex output for current organism
	re := regexp.MustCompile(fmt.Sprintf(`%s\n((.*)\n)*%s`, StartName(funcName, index), EndName(funcName, index)))
	curFuncOutput := re.FindString(printed)
	// Check if function paniced
	if strings.Contains(curFuncOutput, fmt.Sprintf("Recovered in Test%s%d", funcName, index)) {
		info.Panics = true
		return
	}
	// Otherwise split lines
	lines := strings.Split(curFuncOutput, "\n")
	for _, line := range lines {
		// Check if line starts with expected JSON
		if strings.HasPrefix(line, `{ "type":`) {
			// Create assert statements from JSON line
			assertStmts := info.ParseLine(line, &mem)
			if firstRun {
				info.AssertStmts = append(info.AssertStmts, assertStmts...)
			} else {
				info.SecondRun = append(info.SecondRun, assertStmts...)
			}
		}
	}
	if !firstRun {
		info.SetIsValid()
	}
}
