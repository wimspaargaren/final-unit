//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wimspaargaren/final-unit/internal/evo"
	"github.com/wimspaargaren/final-unit/internal/gen"
)

type EvoTestSuite struct {
	suite.Suite
}

func (s *EvoTestSuite) TestExamplesSingleFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestPanicFunc0", r)
		}
	}()
	dir := "examples/simple"

	genOpts := gen.DefaultOpts()
	genOpts.OrganismAmount = 1
	generator, err := gen.New(dir, "e2e", genOpts)
	s.Require().NoError(err)
	popOpts := evo.DefaultPopOpts(generator)
	popOpts.OverrideTestCases = true
	popOpts.Target = 0
	popOpts.MaxNoImprovGens = 1
	population, err := evo.NewPopulation(dir, popOpts)
	s.Require().NoError(err)
	err = population.Evolve()
	s.Require().NoError(err)
}

func (s *EvoTestSuite) TestInputExamples() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestPanicFunc0", r)
		}
	}()
	files, err := ioutil.ReadDir("../../test/data/inputs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		s.Run(f.Name(), func() {
			dir := "../../test/data/inputs/" + f.Name()

			genOpts := gen.DefaultOpts()
			genOpts.OrganismAmount = 1
			generator, err := gen.New(dir, "e2e", genOpts)
			s.Require().NoError(err)
			popOpts := evo.DefaultPopOpts(generator)
			popOpts.OverrideTestCases = true
			popOpts.Target = 0
			popOpts.MaxNoImprovGens = 1
			population, err := evo.NewPopulation(dir, popOpts)
			s.Require().NoError(err)
			err = population.Evolve()
			s.Require().NoError(err)
		})
	}
}

func (s *EvoTestSuite) TestOutputExamples() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestPanicFunc0", r)
		}
	}()
	files, err := ioutil.ReadDir("../../test/data/outputs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		s.Run(f.Name(), func() {
			dir := "../../test/data/outputs/" + f.Name()

			genOpts := gen.DefaultOpts()
			genOpts.OrganismAmount = 1
			generator, err := gen.New(dir, "e2e", genOpts)
			s.Require().NoError(err)
			popOpts := evo.DefaultPopOpts(generator)
			popOpts.OverrideTestCases = true
			popOpts.Target = 0
			popOpts.MaxNoImprovGens = 1
			population, err := evo.NewPopulation(dir, popOpts)
			s.Require().NoError(err)
			err = population.Evolve()
			s.Require().NoError(err)
		})
	}
}

func TestEvoTestSuite(t *testing.T) {
	suite.Run(t, new(EvoTestSuite))
}
