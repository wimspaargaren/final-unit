package runtime

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

// StmtType indicates the type of statement
type StmtType string

// Different statement types
const (
	StmtTypeAssert StmtType = "assert"
	StmtTypeAssign StmtType = "assign"
)

// Stmt statement interface
type Stmt interface {
	Type() StmtType
	Replace(key, val string)
}

// AssignStmtType is the type of assign statement
type AssignStmtType string

// Different assign statement types
const (
	AssignStmtTypeDefine AssignStmtType = ":="
	AssignSTmtTypeAssign AssignStmtType = "="
)

// AssignStmt an assign statement
type AssignStmt struct {
	AssignStmtType AssignStmtType
	LeftHand       string
	RightHand      string
}

// Type retrieves the type of assert stmt
func (a *AssignStmt) Type() StmtType {
	return StmtTypeAssign
}

// Replace replaces key with value used to resolve map and array index values
func (a *AssignStmt) Replace(key, val string) {
	if strings.Contains(a.RightHand, key) {
		a.RightHand = strings.ReplaceAll(a.RightHand, key, val)
	}
}

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
	AssertStmtType AssertStmtType
	Expected       string
	Value          string
}

// Type retrieves the type of assert stmt
func (a *AssertStmt) Type() StmtType {
	return StmtTypeAssert
}

// Replace replaces key with value
func (a *AssertStmt) Replace(key, val string) {
}

// StmtPrinter printer for assert statements
type StmtPrinter interface {
	fmt.Stringer
	PrintStmt(stmt Stmt) string
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

// PrintStmt prints a statement
func (t *TestifySuitePrinter) PrintStmt(stmt Stmt) string {
	switch tp := stmt.(type) {
	case *AssertStmt:
		return t.PrintAssertStmt(tp)
	case *AssignStmt:
		return t.PrintAssignStmt(tp)
	default:
		log.Warningf("unexpected stmt type")
		return ""
	}
}

// PrintAssertStmt prints an assert statement for a testcase in a testify suite
func (t *TestifySuitePrinter) PrintAssertStmt(astmt *AssertStmt) string {
	switch astmt.AssertStmtType {
	case AssertStmtTypeEqualValues:
		return fmt.Sprintf("%s.%s(%s,%s)", t.Receiver, astmt.AssertStmtType, astmt.Expected, astmt.Value)
	case AssertStmtTypeNil,
		AssertStmtTypeNoError,
		AssertStmtTypeError,
		AssertStmtTypeFalse,
		AssertStmtTypeTrue:
		return fmt.Sprintf("%s.%s(%s)", t.Receiver, astmt.AssertStmtType, astmt.Expected)
	default:
		log.Warningf("unexpected assert stmt type")
		return fmt.Sprintf("// FIXME: unknown assertion %s.%s(%s,%s)", t.Receiver, astmt.AssertStmtType, astmt.Expected, astmt.Value)
	}
}

// PrintAssignStmt prints an assign stmt
func (t *TestifySuitePrinter) PrintAssignStmt(astmt *AssignStmt) string {
	switch astmt.AssignStmtType {
	case AssignSTmtTypeAssign,
		AssignStmtTypeDefine:
		return fmt.Sprintf("%s %s %s", astmt.LeftHand, astmt.AssignStmtType, astmt.RightHand)
	default:
		log.Warningf("unexpected assert stmt type")
		return fmt.Sprintf("// FIXME: unknown assign %s %s %s", astmt.LeftHand, astmt.AssignStmtType, astmt.RightHand)
	}
}

func (t *TestifySuitePrinter) String() string {
	return "testify suite printer"
}
