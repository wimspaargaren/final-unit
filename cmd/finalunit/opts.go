package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/evo"
	"github.com/wimspaargaren/final-unit/internal/gen"
)

// Opts current options
type Opts struct {
	Verbose bool
	Version bool
	Dir     string
	Debug   bool

	gen.Options
	evo.PopulationOpts
}

// LogLevel get the log level base on options
func (o *Opts) LogLevel() log.Level {
	lvl := log.WarnLevel
	if o.Verbose || o.Version {
		lvl = log.InfoLevel
	}
	if o.Debug {
		lvl = log.DebugLevel
	}
	return lvl
}
