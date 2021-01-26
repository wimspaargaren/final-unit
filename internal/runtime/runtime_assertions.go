package runtime

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// AssertStmtType Type of assert statement
type AssertStmtType string

// Used assert statement types
const (
	AssertStmtTypeEqualValues AssertStmtType = "EqualValues"
	AssertStmtTypeNil         AssertStmtType = "Nil"
	AssertStmtTypeNoError     AssertStmtType = "NoError"
	AssertStmtTypeError       AssertStmtType = "Error"
	AssertStmtTypeFalse       AssertStmtType = "False"
	AssertStmtTypeTrue        AssertStmtType = "True"
)

// AssertStmt an assert statement
type AssertStmt struct {
	Type     AssertStmtType
	Expected string
	Value    string
}

// StmtPrinter printer for assert statements
type StmtPrinter interface {
	fmt.Stringer
	PrintAssertStmt(astmt AssertStmt) string
}

// TestifySuitePrinter printer for testify suites
type TestifySuitePrinter struct {
	Receiver string
}

// NewTestifySuitePrinter new testify suite
func NewTestifySuitePrinter(receiver string) StmtPrinter {
	return &TestifySuitePrinter{
		Receiver: receiver,
	}
}

// PrintAssertStmt prints an assert statement for a testcase in a testify suite
func (t *TestifySuitePrinter) PrintAssertStmt(astmt AssertStmt) string {
	switch astmt.Type {
	case AssertStmtTypeEqualValues:
		return fmt.Sprintf("%s.%s(%s,%s)", t.Receiver, astmt.Type, astmt.Expected, astmt.Value)
	case AssertStmtTypeNil,
		AssertStmtTypeNoError,
		AssertStmtTypeError,
		AssertStmtTypeFalse,
		AssertStmtTypeTrue:
		return fmt.Sprintf("%s.%s(%s)", t.Receiver, astmt.Type, astmt.Expected)
	default:
		log.Warningf("unexpected assert stmt type")
		return fmt.Sprintf("// FIXME: unknown assertion %s.%s(%s,%s)", t.Receiver, astmt.Type, astmt.Expected, astmt.Value)
	}
}

func (t *TestifySuitePrinter) String() string {
	return "testify suite printer"
}
