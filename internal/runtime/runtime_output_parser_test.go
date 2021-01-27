package runtime

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RunTimeOutputParserTestSuite struct {
	suite.Suite
}

func (s *RunTimeOutputParserTestSuite) TestExampleOutputs() {
	tests := []struct {
		Name   string
		Input  string
		Output []string
	}{
		{
			Name:   "int",
			Input:  `{ "type": "int", "var_name": "podrt", "val": "0"}`,
			Output: []string{"s.EqualValues(int(0),podrt)"},
		},
		{
			Name:   "string",
			Input:  `{ "type": "string", "var_name": "mtjaz", "val": "Hello World"}`,
			Output: []string{"s.EqualValues(string(`Hello World`),mtjaz)"},
		},
		{
			Name:   "array",
			Input:  `{ "type": "arr", "arr_ident": "zFmET", "var_name": "out", "val": "4", "child": { "type": "pointer", "var_name": "out[zFmET]", "child": { "type": "struct", "var_name": "pointerOut", "child": { "type": "arr", "arr_ident": "hhQHG", "var_name": "pointerOut.Y", "val": "1", "child": { "type": "pointer", "var_name": "pointerOut.Y[hhQHG]", "child": { "type": "int", "var_name": "pointerOut2", "val": "2"}}}}}}`,
			Output: []string{"pointerOut := *out[4]", "pointerOut2 := *pointerOut.Y[1]", "s.EqualValues(int(2),pointerOut2)"},
		},
		{
			Name:   "map nil",
			Input:  `{ "type": "map", "arr_ident": "retWA", "map_key_type": "int", "var_name": "out", "val": "3", "child": { "type": "pointer", "var_name": "out[retWA]", "val": "nil" } }`,
			Output: []string{"s.Nil(out[3])"},
		},
		{
			Name:   "map not nil",
			Input:  `{ "type": "map", "arr_ident": "qKejm", "map_key_type": "int", "var_name": "out", "val": "3", "child": { "type": "pointer", "var_name": "out[qKejm]", "child": { "type": "map", "arr_ident": "wncbs", "map_key_type": "string", "var_name": "pointerOut", "val": "asdf", "child": { "type": "struct", "var_name": "pointerOut[wncbs]", "child": { "type": "pointer", "var_name": "pointerOut[wncbs].Y", "child": { "type": "string", "var_name": "pointerOut2", "val": "asdf"}}}}}}`,
			Output: []string{"pointerOut := *out[3]", "pointerOut2 := *pointerOut[\"asdf\"].Y", "s.EqualValues(string(`asdf`),pointerOut2)"},
		},
		{
			Name:   "struct",
			Input:  `{ "type": "pointer", "var_name": "out", "child": { "type": "struct", "var_name": "pointerOut", "child": { "type": "pointer", "var_name": "pointerOut.C", "child": { "type": "struct", "var_name": "pointerOut2", "child": { "type": "int", "var_name": "pointerOut2.X", "val": "2"}}}}}`,
			Output: []string{"pointerOut := *out", "pointerOut2 := *pointerOut.C", "s.EqualValues(int(2),pointerOut2.X)"},
		},
		{
			Name:   "custom",
			Input:  `{ "type": "custom", "var_name": "Something", "child": { "type": "byte", "var_name": "out", "val": "0x4"}}`,
			Output: []string{"s.EqualValues(byte(0x4),out)"},
		},
		{
			Name:   "bool true",
			Input:  `{ "type": "bool", "var_name": "out", "val": "true"}`,
			Output: []string{"s.True(out)"},
		},
		{
			Name:   "bool false",
			Input:  `{ "type": "bool", "var_name": "out", "val": "false"}`,
			Output: []string{"s.False(out)"},
		},
		{
			Name:   "err nil",
			Input:  `{ "type": "error", "var_name": "out2", "val": "nil" } `,
			Output: []string{"s.NoError(out2)"},
		},
		{
			Name:   "err not nil",
			Input:  `{ "type": "error", "var_name": "out", "val": "notnil" } `,
			Output: []string{"s.Error(out)"},
		},
		{
			Name:   "complex128",
			Input:  `{ "type": "complex128", "var_name": "out6", "val": "(234.33333333333334+0i)"}`,
			Output: []string{"s.EqualValues(complex128(234.33333333333334+0i),out6)"},
		},
	}

	for _, testCase := range tests {
		s.Run(testCase.Name, func() {
			printer := NewTestifySuitePrinter("s")
			info := NewInfo(printer)
			x := info.ParseLine(testCase.Input, &[]string{})
			s.EqualValues(testCase.Output, x)
		})
	}
}

func TestRuntTimeTestSuite(t *testing.T) {
	suite.Run(t, new(RunTimeOutputParserTestSuite))
}
