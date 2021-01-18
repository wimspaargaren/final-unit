package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/evo"
	"github.com/wimspaargaren/final-unit/internal/gen"
)

// Opts current options
type Opts struct {
	Verbose  bool
	Version  bool
	Dir      string
	LogLevel log.Level
}

// InitOpts init generator opts
func InitOpts() (*Opts, *gen.Options, evo.PopulationOpts) {
	// Run opts
	version := flag.Bool("version", false, "current version")
	verbose := flag.Bool("v", false, "run generator in verbose mode")
	var debug *bool
	if Version == "dev" {
		debug = flag.Bool("debug", false, "run generator in debug mode")
	}
	dir := flag.String("d", ".", "dir for which to execute the generator")

	// Pop Opts
	noImprovedGens := flag.Int("no-improve-gens", 10, "max amount of generations without improvements before the generator halts")
	targetFitness := flag.Int("target-fitness", 95, "number between 0 and 100 indicating the target coverage we try to hit")
	// override := flag.Bool("override-test-cases", false, "if set overrides already existing test cases")

	// Gen opts
	organismAmount := flag.Int("org-amount", 10, "amount of organisms in the population")
	testCasesPerFunc := flag.Int("test-cases-func", 10, "amount of test cases created for every function")

	// Parse flags
	flag.Parse()

	// Normal options
	opts := &Opts{
		Verbose:  *verbose,
		Dir:      *dir,
		LogLevel: log.WarnLevel,
		Version:  *version,
	}
	if opts.Verbose || opts.Version {
		opts.LogLevel = log.InfoLevel
	}
	if debug != nil && *debug {
		opts.LogLevel = log.DebugLevel
	}

	// Initialise the logger
	setLogger(opts.LogLevel)

	// Population options
	popOpts := evo.DefaultPopOpts(nil)
	popOpts.MaxNoImprovGens = *noImprovedGens
	// popOpts.OverrideTestCases = *override
	customFitness := float64(*targetFitness)
	if customFitness >= 0 && customFitness <= 100 {
		popOpts.Target = float64(*targetFitness)
	} else {
		log.Warningf("incorrect fitness specified, should be in range 0 and 100, using default")
	}

	// Generation options

	genOpts := gen.DefaultOpts()
	genOpts.OrganismAmount = *organismAmount
	genOpts.TestCasesPerFunc = *testCasesPerFunc

	return opts, genOpts, popOpts
}
