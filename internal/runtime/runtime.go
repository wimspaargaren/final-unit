// Package runtime analyses runtime output and converts it into assert statements
package runtime

// Info information about values on runtime
type Info struct {
	Panics      bool
	AssertStmts []Stmt
	SecondRun   []Stmt
	Printer     StmtPrinter
}

// NewInfo creates new runtime info for given printer
func NewInfo(printer StmtPrinter) *Info {
	return &Info{
		Printer: printer,
	}
}

// GetAssertStmts retrieve the assert statements
func (info *Info) GetAssertStmts() []string {
	res := []string{}
	for _, stmt := range info.AssertStmts {
		res = append(res, info.Printer.PrintStmt(stmt))
	}
	return res
}

// IsValid verifies that created runtime info is valid
// used when generating end result
func (info *Info) IsValid() bool {
	if len(info.AssertStmts) != len(info.SecondRun) {
		return false
	}

	for i := 0; i < len(info.AssertStmts); i++ {
		if info.AssertStmts[i].Type() != info.SecondRun[i].Type() {
			return false
		}
		assertStmt, ok := info.AssertStmts[i].(*AssertStmt)
		assertStmt2, ok2 := info.SecondRun[i].(*AssertStmt)
		if ok && ok2 {
			if *assertStmt != *assertStmt2 {
				return false
			}
			continue
		}
		assignStmt, ok := info.AssertStmts[i].(*AssertStmt)
		assignStmt2, ok2 := info.SecondRun[i].(*AssertStmt)
		if ok && ok2 {
			if *assignStmt != *assignStmt2 {
				return false
			}
			continue
		}
	}
	return true
}

// AssertStmtsForTestCase creates assert statements for a testcase
func (info *Info) AssertStmtsForTestCase(printed string, firstRun bool, funcName string, index int) {
	outputParser := NewOutputParser()
	stmts, panics := outputParser.Parse(printed, funcName, index)
	if panics {
		info.Panics = true
		return
	}
	if firstRun {
		info.AssertStmts = append(info.AssertStmts, stmts...)
	} else {
		info.SecondRun = append(info.SecondRun, stmts...)
	}
}
