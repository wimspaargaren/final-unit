package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/evo"
	"github.com/wimspaargaren/final-unit/internal/gen"
	"github.com/wimspaargaren/final-unit/pkg/seed"
)

// Version current version
var Version = "dev"

func setLogger(level log.Level) {
	log.SetLevel(level)

	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
}

func main() {
	globalOpts := Opts{}
	if err := initCmd(&globalOpts).Execute(); err != nil {
		return
	}
	setLogger(globalOpts.LogLevel())
	if err := Verify(&globalOpts); err != nil {
		log.Fatalln(err.Error())
	}
	if globalOpts.OverrideTestCases {
		log.Warningf("running in override mode, existing tests will be replaced!")
	}
	// Start generating
	seed.SetRandomSeed(time.Now().Unix())
	Generate(globalOpts.Dir, &globalOpts)
	log.Infof("generating test cases complete")
}

// Generate starts the generator
func Generate(dir string, opts *Opts) {
	if dir == "." {
		log.Infof("analysing current directory")
	} else {
		log.Infof("analysing directory: %s", dir)
	}
	generator, err := gen.New(dir, &opts.Options)
	if err != nil {
		log.WithError(err).Debug("unable to create new generator")
		log.Fatal("something unexpected went wrong trying to generate the test cases")
	}

	log.Infof("creating first generation")
	population, err := evo.NewPopulation(dir, generator, opts.PopulationOpts)
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
