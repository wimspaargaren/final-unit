package main

import (
	"github.com/asherascout/final-unit/internal/evo"
	"github.com/asherascout/final-unit/internal/gen"
	log "github.com/sirupsen/logrus"
)

// Version current version
var Version string = "dev"

func setLogger(level log.Level) {
	log.SetLevel(level)

	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
}

func main() {
	// Initialise options
	opts, genOpts, evoOpts := InitOpts()
	if opts.Version {
		log.Infof("final unit version %s", Version)
		return
	}

	err := Verify(opts)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if evoOpts.OverrideTestCases {
		log.Warningf("running in override mode, existing tests will be replaced!")
	}

	// Start generating
	Generate(opts.Dir, genOpts, evoOpts)
	log.Infof("generating test cases complete")
}

// Generate starts the generator
func Generate(dir string, genOpts *gen.Options, popOpts evo.PopulationOpts) {
	if dir == "." {
		log.Infof("analysing current directory")
	} else {
		log.Infof("analysing directory: %s", dir)
	}
	generator, err := gen.New(dir, Version, genOpts)
	if err != nil {
		log.WithError(err).Debug("unable to create new generator")
		log.Fatal("something unexpected went wrong trying to generate the test cases")
	}

	popOpts.OrgGenerator = generator

	log.Infof("creating first generation")
	population, err := evo.NewPopulation(dir, popOpts)
	if err != nil {
		log.WithError(err).Debug("unable to create new population")
		log.Fatal("something unexpected went wrong trying to generate the test cases")
	}
	log.Infof("start to evolve population")
	err = population.Evolve()
	if err != nil {
		log.WithError(err).Debug("unable to evolve the population")
		log.Fatal("something unexpected went wrong trying to generate the test cases")
	}
}
