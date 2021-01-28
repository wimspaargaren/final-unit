package runtime

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RunTimeTestSuite struct {
	suite.Suite
}

func (s *RunTimeTestSuite) TestAssertStmtsForTestCase() {
	info := &Info{
		Printer: NewTestifySuitePrinter("s"),
	}
	info.AssertStmtsForTestCase(output, true, "DoubleArray", 1)
	s.Equal(0, len(info.GetAssertStmts()))
	info.AssertStmtsForTestCase(output, true, "DoubleArrayx", 0)
	s.Equal(0, len(info.GetAssertStmts()))
	info.AssertStmtsForTestCase(output, true, "DoubleArray", 0)
	s.Equal(4, len(info.GetAssertStmts()))
	info.AssertStmtsForTestCase(output, false, "DoubleArray", 0)
	s.Equal(4, len(info.SecondRun))
}

func (s *RunTimeTestSuite) TestAssertStmtsForPanicTestCase() {
	info := &Info{
		Printer: NewTestifySuitePrinter("s"),
	}
	info.AssertStmtsForTestCase(panicOutput, true, "DoubleArray", 0)
	s.True(info.Panics)
}

func (s *RunTimeTestSuite) TestIsValid() {
	tests := []struct {
		Name     string
		Input    *Info
		Expected bool
	}{
		{
			Name: "length not equal",
			Input: &Info{
				AssertStmts: []Stmt{},
				SecondRun:   []Stmt{&AssertStmt{Expected: "some val"}},
			},
			Expected: false,
		},
		{
			Name: "not equal val",
			Input: &Info{
				AssertStmts: []Stmt{&AssertStmt{Expected: "not equal val"}},
				SecondRun:   []Stmt{&AssertStmt{Expected: "other val"}},
			},
			Expected: false,
		},
		{
			Name: "not equal type",
			Input: &Info{
				AssertStmts: []Stmt{&AssignStmt{LeftHand: "not equal val"}},
				SecondRun:   []Stmt{&AssertStmt{Expected: "other val"}},
			},
			Expected: false,
		},
		{
			Name: "equal",
			Input: &Info{
				AssertStmts: []Stmt{&AssertStmt{Expected: "equal val"}},
				SecondRun:   []Stmt{&AssertStmt{Expected: "equal val"}},
			},
			Expected: true,
		},
	}

	for _, testCase := range tests {
		s.Run(testCase.Name, func() {
			testCase.Input.IsValid()
			s.Equal(testCase.Expected, testCase.Input.IsValid())
		})
	}
}

func TestRunTimeTestSuite(t *testing.T) {
	suite.Run(t, new(RunTimeTestSuite))
}

const panicOutput = `
<START;DoubleArray0>
Recovered in TestDoubleArray0
<END;DoubleArray0>`

const output = `=== RUN   TestArraysSuite/TestDoubleArray0
<START;DoubleArray0>
{ "type": "arr", "arr_ident": "mxcRp", "var_name": "out", "val": "0", "child": { "type": "arr", "arr_ident": "nzxxp", "var_name": "out[mxcRp]", "val": "0", "child": { "type": "int", "var_name": "out[mxcRp][nzxxp]", "val": "3"}}}
{ "type": "arr", "arr_ident": "mxcRp", "var_name": "out", "val": "0", "child": { "type": "arr", "arr_ident": "nzxxp", "var_name": "out[mxcRp]", "val": "1", "child": { "type": "int", "var_name": "out[mxcRp][nzxxp]", "val": "4"}}}
{ "type": "arr", "arr_ident": "mxcRp", "var_name": "out", "val": "1", "child": { "type": "arr", "arr_ident": "nzxxp", "var_name": "out[mxcRp]", "val": "0", "child": { "type": "int", "var_name": "out[mxcRp][nzxxp]", "val": "5"}}}
{ "type": "arr", "arr_ident": "mxcRp", "var_name": "out", "val": "1", "child": { "type": "arr", "arr_ident": "nzxxp", "var_name": "out[mxcRp]", "val": "1", "child": { "type": "int", "var_name": "out[mxcRp][nzxxp]", "val": "6"}}}
<END;DoubleArray0>
`
