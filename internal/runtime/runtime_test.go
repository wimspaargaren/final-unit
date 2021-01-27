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
	s.Equal(0, len(info.AssertStmts))
	info.AssertStmtsForTestCase(output, true, "DoubleArrayx", 0)
	s.Equal(0, len(info.AssertStmts))
	info.AssertStmtsForTestCase(output, true, "DoubleArray", 0)
	s.Equal(4, len(info.AssertStmts))
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
	info := &Info{
		AssertStmts: []string{},
		SecondRun:   []string{"hi"},
		Printer:     NewTestifySuitePrinter("s"),
	}
	info.SetIsValid()
	s.False(info.IsValid)

	tests := []struct {
		Name     string
		Input    *Info
		Expected bool
	}{
		{
			Name: "length not equal",
			Input: &Info{
				AssertStmts: []string{},
				SecondRun:   []string{"hi"},
			},
			Expected: false,
		},
		{
			Name: "not equal",
			Input: &Info{
				AssertStmts: []string{"other val"},
				SecondRun:   []string{"hi"},
			},
			Expected: false,
		},
		{
			Name: "equal",
			Input: &Info{
				AssertStmts: []string{"equal val"},
				SecondRun:   []string{"equal val"},
			},
			Expected: true,
		},
	}

	for _, testCase := range tests {
		s.Run(testCase.Name, func() {
			testCase.Input.SetIsValid()
			s.Equal(testCase.Expected, testCase.Input.IsValid)
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
