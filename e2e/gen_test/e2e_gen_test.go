//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"testing"

	"github.com/asherascout/final-unit/internal/gen"
	"github.com/asherascout/final-unit/internal/tmplexec"
	"github.com/stretchr/testify/suite"
)

type E2ESuite struct {
	suite.Suite
}

func (s *E2ESuite) TestExamples() {
	tests := []struct {
		CasesPerFunc   int
		OrganismAmount int
		Dir            string
	}{
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/interface_other",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/named_multi_param_return",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/example_directly_nested_struct",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/context",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/other_import",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/chan",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/chan_other",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/ignore_main",
		},
		{
			CasesPerFunc:   10,
			OrganismAmount: 1,
			Dir:            "examples/directlynestedinterface",
		},
		{
			CasesPerFunc:   10,
			OrganismAmount: 1,
			Dir:            "examples/simplepackage",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/packagewithimports",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/panic",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/interface_cycle",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/importadditional",
		},
	}
	for _, test := range tests {
		executor := tmplexec.NewCoverageExecutor(tmplexec.Opts{Dir: test.Dir, Override: true})
		s.Run(test.Dir, func() {
			opts := gen.DefaultOpts()
			opts.OrganismAmount = test.OrganismAmount
			opts.TestCasesPerFunc = test.CasesPerFunc
			g, err := gen.New(test.Dir, "e2e", opts)
			if err != nil {
				panic(err)
			}
			organisms := g.GetTestCases()
			s.Require().Equal(1, len(organisms))

			coverage, err := executor.Execute(organisms[0])
			s.Require().NoError(err)
			fmt.Println("Coverage: ", coverage)
		})
	}
}

func TestAssignStmtGeneratorSuite(t *testing.T) {
	suite.Run(t, new(E2ESuite))
}
