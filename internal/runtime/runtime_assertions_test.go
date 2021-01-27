package runtime

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RunTimeAssertionsTestSuite struct {
	suite.Suite
}

func (s *RunTimeAssertionsTestSuite) TestAssertTestifyPrinterAssertStmts() {
	printer := NewTestifySuitePrinter("s")
	s.Run(printer.String(), func() {
		tests := []struct {
			Name   string
			Input  Stmt
			Output string
		}{
			{
				Name: "Error assertion",
				Input: &AssertStmt{
					AssertStmtType: AssertStmtTypeError,
					Value:          "",
					Expected:       "err",
				},
				Output: `s.Error(err)`,
			},
			{
				Name: "No Error assertion",
				Input: &AssertStmt{
					AssertStmtType: AssertStmtTypeNoError,
					Value:          "",
					Expected:       "err",
				},
				Output: `s.NoError(err)`,
			},
			{
				Name: "bool true",
				Input: &AssertStmt{
					AssertStmtType: AssertStmtTypeTrue,
					Value:          "",
					Expected:       "bool",
				},
				Output: `s.True(bool)`,
			},
			{
				Name: "bool false",
				Input: &AssertStmt{
					AssertStmtType: AssertStmtTypeFalse,
					Value:          "",
					Expected:       "bool",
				},
				Output: `s.False(bool)`,
			},
			{
				Name: "nil",
				Input: &AssertStmt{
					AssertStmtType: AssertStmtTypeNil,
					Value:          "",
					Expected:       "var",
				},
				Output: `s.Nil(var)`,
			},
			{
				Name: "equal vals",
				Input: &AssertStmt{
					AssertStmtType: AssertStmtTypeEqualValues,
					Value:          "val",
					Expected:       "exp",
				},
				Output: `s.EqualValues(exp,val)`,
			},
			{
				Name: "unknown type",
				Input: &AssertStmt{
					AssertStmtType: AssertStmtType("unknown"),
					Value:          "val",
					Expected:       "exp",
				},
				Output: `// FIXME: unknown assertion s.unknown(exp,val)`,
			},
		}

		for _, testCase := range tests {
			s.Run(testCase.Name, func() {
				printed := printer.PrintStmt(testCase.Input)
				s.Equal(testCase.Output, printed)
			})
		}
	})
}

func (s *RunTimeAssertionsTestSuite) TestAssertTestifyPrinterAssignStmts() {
	printer := NewTestifySuitePrinter("s")
	s.Run(printer.String(), func() {
		tests := []struct {
			Name   string
			Input  Stmt
			Output string
		}{
			{
				Name: "Assign statement defineassign",
				Input: &AssignStmt{
					AssignStmtType: AssignStmtTypeDefine,
					LeftHand:       "x",
					RightHand:      "y",
				},
				Output: `x := y`,
			},
			{
				Name: "Assign statement assign",
				Input: &AssignStmt{
					AssignStmtType: AssignSTmtTypeAssign,
					LeftHand:       "x",
					RightHand:      "y",
				},
				Output: `x = y`,
			},
			{
				Name: "Assign statement assign",
				Input: &AssignStmt{
					AssignStmtType: AssignSTmtTypeAssign,
					LeftHand:       "x",
					RightHand:      "y",
				},
				Output: `x = y`,
			},
			{
				Name: "Assign statement assign",
				Input: &AssignStmt{
					AssignStmtType: AssignStmtType(""),
					LeftHand:       "x",
					RightHand:      "y",
				},
				Output: `// FIXME: unknown assign x  y`,
			},
		}

		for _, testCase := range tests {
			s.Run(testCase.Name, func() {
				printed := printer.PrintStmt(testCase.Input)
				s.Equal(testCase.Output, printed)
			})
		}
	})
}

func TestRunTimeAssertionsTestSuite(t *testing.T) {
	suite.Run(t, new(RunTimeAssertionsTestSuite))
}
