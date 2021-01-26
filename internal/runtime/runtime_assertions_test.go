package runtime

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RunTimeAssertionsTestSuite struct {
	suite.Suite
}

func (s *RunTimeAssertionsTestSuite) TestAssertTestifyPrinterStmts() {
	printer := NewTestifySuitePrinter("s")
	s.Run(printer.String(), func() {
		tests := []struct {
			Name   string
			Input  AssertStmt
			Output string
		}{
			{
				Name: "Error assertion",
				Input: AssertStmt{
					Type:     AssertStmtTypeError,
					Value:    "",
					Expected: "err",
				},
				Output: `s.Error(err)`,
			},
			{
				Name: "No Error assertion",
				Input: AssertStmt{
					Type:     AssertStmtTypeNoError,
					Value:    "",
					Expected: "err",
				},
				Output: `s.NoError(err)`,
			},
			{
				Name: "bool true",
				Input: AssertStmt{
					Type:     AssertStmtTypeTrue,
					Value:    "",
					Expected: "bool",
				},
				Output: `s.True(bool)`,
			},
			{
				Name: "bool false",
				Input: AssertStmt{
					Type:     AssertStmtTypeFalse,
					Value:    "",
					Expected: "bool",
				},
				Output: `s.False(bool)`,
			},
			{
				Name: "nil",
				Input: AssertStmt{
					Type:     AssertStmtTypeNil,
					Value:    "",
					Expected: "var",
				},
				Output: `s.Nil(var)`,
			},
			{
				Name: "equal vals",
				Input: AssertStmt{
					Type:     AssertStmtTypeEqualValues,
					Value:    "val",
					Expected: "exp",
				},
				Output: `s.EqualValues(exp,val)`,
			},
			{
				Name: "unknown type",
				Input: AssertStmt{
					Type:     AssertStmtType("unknown"),
					Value:    "val",
					Expected: "exp",
				},
				Output: `// FIXME: unknown assertion s.unknown(exp,val)`,
			},
		}

		for _, testCase := range tests {
			s.Run(testCase.Name, func() {
				printed := printer.PrintAssertStmt(testCase.Input)
				s.Equal(testCase.Output, printed)
			})
		}
	})
}

func TestRunTimeAssertionsTestSuite(t *testing.T) {
	suite.Run(t, new(RunTimeAssertionsTestSuite))
}
