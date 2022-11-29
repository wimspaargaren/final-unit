//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/wimspaargaren/final-unit/internal/evo"
	"github.com/wimspaargaren/final-unit/internal/gen"
	"github.com/wimspaargaren/final-unit/pkg/seed"
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
	dir := "./examples/simple"

	genOpts := &gen.Options{
		OrganismAmount:   1,
		MaxRecursion:     3,
		TestCasesPerFunc: 1,
	}
	seed.SetRandomSeed(time.Now().Unix())
	generator, err := gen.New(dir, genOpts)
	s.Require().NoError(err)
	popOpts := evo.PopulationOpts{
		OverrideTestCases: true,
		Target:            0,
		MaxNoImprovGens:   1,
	}
	population, err := evo.NewPopulation(dir, generator, popOpts)
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
	files, err := os.ReadDir("../../test/data/inputs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		s.Run(f.Name(), func() {
			dir := "../../test/data/inputs/" + f.Name()

			genOpts := &gen.Options{
				OrganismAmount:   1,
				MaxRecursion:     3,
				TestCasesPerFunc: 1,
			}
			seed.SetRandomSeed(time.Now().Unix())
			generator, err := gen.New(dir, genOpts)
			s.Require().NoError(err)
			popOpts := evo.PopulationOpts{
				OverrideTestCases: true,
				Target:            0,
				MaxNoImprovGens:   1,
			}
			population, err := evo.NewPopulation(dir, generator, popOpts)
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
	files, err := os.ReadDir("../../test/data/outputs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		s.Run(f.Name(), func() {
			dir := "../../test/data/outputs/" + f.Name()
			genOpts := &gen.Options{
				OrganismAmount:   1,
				MaxRecursion:     3,
				TestCasesPerFunc: 1,
			}
			genOpts.OrganismAmount = 1
			seed.SetRandomSeed(time.Now().Unix())
			generator, err := gen.New(dir, genOpts)
			s.Require().NoError(err)
			popOpts := evo.PopulationOpts{
				OverrideTestCases: true,
				Target:            0,
				MaxNoImprovGens:   1,
			}
			population, err := evo.NewPopulation(dir, generator, popOpts)
			s.Require().NoError(err)
			err = population.Evolve()
			s.Require().NoError(err)
		})
	}
}

func TestEvoTestSuite(t *testing.T) {
	suite.Run(t, new(EvoTestSuite))
}
