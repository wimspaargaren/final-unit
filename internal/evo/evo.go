// Package evo provides genetic evolution algorithm for evolving test cases
// nolint: gosec
package evo

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/gen"
	"github.com/wimspaargaren/final-unit/internal/tmplexec"
	"github.com/wimspaargaren/final-unit/pkg/chance"
)

// error definitions
var (
	ErrOrganismCantBreed error = fmt.Errorf("organisms dont match, expected equal amount of files")
)

// nolint: gochecknoinits
func init() {
	rand.Seed(time.Now().Unix())
}

// Default values for population
const (
	DefaultMutationRate float64 = 5
	DefaultTarget       float64 = 95.0
	DefaultNoImprovGens int     = 10
)

// PopulationStats struct representing the stats of the population
type PopulationStats struct {
	AverageFit     float64
	NoImprovedGens int
	BestFit        float64
	Generation     int
}

func (s *PopulationStats) logStats() {
	log.Infof("bestFit: %.2f, average fit: %.2f, no improved gens; %d, generation: %d", s.BestFit, s.AverageFit, s.NoImprovedGens, s.Generation)
}

// PopulationOpts struct representing options of the population
type PopulationOpts struct {
	MutationRate      float64
	Target            float64
	MaxNoImprovGens   int
	OrgGenerator      *gen.Generator
	OverrideTestCases bool
}

// DefaultPopOpts create some default options for the population
func DefaultPopOpts(gen *gen.Generator) PopulationOpts {
	return PopulationOpts{
		OrgGenerator:      gen,
		MutationRate:      DefaultMutationRate,
		Target:            DefaultTarget,
		MaxNoImprovGens:   DefaultNoImprovGens,
		OverrideTestCases: false,
	}
}

// Population a population which is evolving
type Population struct {
	Organisms []*gen.Organism
	BestFit   *gen.Organism
	Opts      PopulationOpts
	Stats     PopulationStats
	Fitness   float64

	Executor tmplexec.IExecutor
	// StatChan  chan PopulationStats
}

// NewPopulation creates a new population for specified options
func NewPopulation(dir string, opts PopulationOpts) (*Population, error) {
	// Create first generation
	organisms := opts.OrgGenerator.GetTestCases()
	p := &Population{
		Organisms: organisms,
		Opts:      opts,
		Executor:  tmplexec.NewCoverageExecutor(tmplexec.Opts{Dir: dir, Override: opts.OverrideTestCases}),
	}
	err := p.GetFitnessForOrganisms()
	if err != nil {
		return nil, err
	}
	return p, err
}

// GetFitnessForOrganisms retrieve the fitness for the current generation
func (p *Population) GetFitnessForOrganisms() error {
	newBest := false
	totalFitness := 0.0
	log.Debugf("GetFitnessForOrganisms")
	for _, o := range p.Organisms {
		_, err := p.Executor.Execute(o)
		if err != nil {
			return err
		}
	}

	for _, org := range p.Organisms {
		fitness := org.Fitness
		totalFitness += fitness
		org.Fitness = fitness
		if p.BestFit == nil {
			newBest = true
			p.BestFit = org
			p.Stats.BestFit = fitness
		}
		if fitness > p.BestFit.Fitness {
			newBest = true
			p.BestFit = org
			p.Stats.BestFit = fitness
		}
	}
	p.Stats.AverageFit = totalFitness / float64(len(p.Organisms))
	if newBest {
		p.Stats.NoImprovedGens = 0
	} else {
		p.Stats.NoImprovedGens++
	}

	return nil
}

// Evolve evolve a population using natural selection
func (p *Population) Evolve() error {
	for {
		p.Stats.logStats()
		if p.BestFit.Fitness >= p.Opts.Target || p.Stats.NoImprovedGens >= p.Opts.MaxNoImprovGens {
			log.Infof("organism which meets target criteria found, generating result")
			return p.CreateBestFitResult()
		}
		t := time.Now()
		// Perform natural selection for creating a newe generation
		err := p.NaturalSelection()
		if err != nil {
			return err
		}
		log.Infof("natural selection took: %.2fs", time.Since(t).Seconds())
		// Target reached we can stop evolving
	}
}

// CreateBestFitResult creates result for best fit
func (p *Population) CreateBestFitResult() error {
	path := p.Opts.OrgGenerator.Dir
	valueExecutor := tmplexec.NewValueExecutor(tmplexec.Opts{Dir: path, Override: p.Opts.OverrideTestCases})

	// First run
	res, err := valueExecutor.Execute(p.BestFit)
	if err != nil {
		return err
	}
	p.BestFit.UpdateAssertStmts(res, true)

	// Second run
	res, err = valueExecutor.Execute(p.BestFit)
	if err != nil {
		return err
	}
	p.BestFit.UpdateAssertStmts(res, false)

	// Assert executor
	assertExecutor := tmplexec.NewAssertExecutor(tmplexec.Opts{Dir: path, Override: p.Opts.OverrideTestCases})
	_, err = assertExecutor.Execute(p.BestFit)
	if err != nil {
		return err
	}
	return nil
}

// NaturalSelection creates a new generator using natural selection
func (p *Population) NaturalSelection() error {
	log.Debugf("Performing natural selection")
	// Increment generation
	p.Stats.Generation++

	pool := p.createSelectionPool()
	nextGen := []*gen.Organism{}

	for i := 0; i < len(p.Organisms); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := p.Organisms[pool[r1]]
		b := p.Organisms[pool[r2]]
		child, err := p.crossover(a, b)
		if err != nil {
			return err
		}
		nextGen = append(nextGen, child)
	}
	p.Organisms = nextGen

	return p.GetFitnessForOrganisms()
}

func (p *Population) crossover(a, b *gen.Organism) (*gen.Organism, error) {
	res := &gen.Organism{}
	if len(a.Files) != len(b.Files) {
		return nil, ErrOrganismCantBreed
	}
	const crossOverRate = 50
	for i, af := range a.Files {
		bf := b.Files[i]
		x := &gen.File{
			Dir:         af.Dir,
			FileName:    af.FileName,
			PackageName: af.PackageName,
			TestCases:   make(map[string][]*gen.TestCase),
		}
		for funcName, testCaseList := range af.TestCases {
			for j, testCase := range testCaseList {
				if chance.IsChance(p.Opts.MutationRate) {
					x.TestCases[funcName] = append(x.TestCases[funcName], p.Opts.OrgGenerator.FuncDeclToTestCase(testCase.FuncDecl, testCase.Pointer))
					continue
				}
				if chance.IsChance(crossOverRate) {
					x.TestCases[funcName] = append(x.TestCases[funcName], testCase)
				} else {
					x.TestCases[funcName] = append(x.TestCases[funcName], bf.TestCases[funcName][j])
				}
			}
		}

		res.Files = append(res.Files, x)
	}
	return res, nil
}

func (p *Population) createSelectionPool() []int {
	pool := []int{}
	const fullPercent = 100
	for i := range p.Organisms {
		pool = append(pool, i)
	}
	for i, o := range p.Organisms {
		num := int((o.Fitness) * fullPercent)
		for j := 0; j < num; j++ {
			pool = append(pool, i)
		}
	}
	return pool
}
