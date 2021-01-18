package gen

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RuntTimeTestSuite struct {
	suite.Suite
}

func (s *RuntTimeTestSuite) TestExampleOutputs() {
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
	}

	for _, testCase := range tests {
		s.Run(testCase.Name, func() {
			x := ParseLine(testCase.Input, &[]string{})
			s.EqualValues(testCase.Output, x)
		})
	}
}

func TestRuntTimeTestSuite(t *testing.T) {
	suite.Run(t, new(RuntTimeTestSuite))
}
