// +build e2e

package e2e

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wimspaargaren/final-unit/internal/gen"
	"github.com/wimspaargaren/final-unit/internal/tmplexec"
)

type E2EResultSuite struct {
	suite.Suite
}

func (s *E2EResultSuite) TestOutputs() {
	tests := []struct {
		CasesPerFunc   int
		OrganismAmount int
		Dir            string
	}{
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/import_custom_type",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/other_panics",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/weird_interface",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/example_directly_nested_struct",
		},

		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/other_import",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/arrays",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/basic_types",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/chan",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/import_local",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/imports",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/maps",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/multi_file",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/named_returns",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/output_complex",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/pointers",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/structs",
		},
		{
			CasesPerFunc:   1,
			OrganismAmount: 1,
			Dir:            "examples/nondeterministic",
		},
	}
	for _, test := range tests {
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
			path, err := filepath.Abs(test.Dir)
			s.Require().NoError(err)

			organism := organisms[0]

			valueExecutor := tmplexec.NewValueExecutor(tmplexec.Opts{Dir: path})

			// First run
			res, err := valueExecutor.Execute(organism)
			s.Require().NoError(err)
			organism.UpdateAssertStmts(res, true)

			// Second run
			res, err = valueExecutor.Execute(organism)
			s.Require().NoError(err)
			organism.UpdateAssertStmts(res, false)

			// Assert executor
			assertExecutor := tmplexec.NewAssertExecutor(tmplexec.Opts{Dir: path, Override: true})
			_, err = assertExecutor.Execute(organism)
			s.Require().NoError(err)
		})
	}
}

func TestE2EResultSuite(t *testing.T) {
	suite.Run(t, new(E2EResultSuite))
}
