// Package tmplexec provides functionality for generating test cases and executing them
package tmplexec

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/gen"
	"gopkg.in/pipe.v2"
)

// NewCoverageExecutor creates new assert template executor
// this is the first step to measure coverage
func NewCoverageExecutor(opts Opts) IExecutor {
	return &CoverageExecutor{
		Opts: opts,
	}
}

// CoverageExecutor implementation of ICoverageExecutor
type CoverageExecutor struct {
	Opts Opts
}

// Execute executes organism on assert template
func (v *CoverageExecutor) Execute(organism *gen.Organism) (string, error) {
	for _, f := range organism.Files {
		ext := filepath.Ext(f.FileName)
		filePath := filepath.Join(f.Dir, f.FileName)
		executionPath := strings.TrimSuffix(filePath, ext) + "_test.go"
		tmpl, err := template.New("").Funcs(template.FuncMap{
			"add": func(x int) int {
				return x + 1
			},
		}).Parse(coverageTemplate)
		if err != nil {
			return "", err
		}
		file, err := os.Create(filepath.Clean(executionPath))
		if err != nil {
			return "", err
		}
		defer func() {
			err := file.Close()
			if err != nil {
				log.WithError(err).Error("unable to close file")
			}
		}()

		err = tmpl.Execute(file, f)
		if err != nil {
			return "", err
		}
	}
	script := pipe.Script(
		pipe.Exec("goimports", "-w", v.Opts.Dir),
		pipe.Exec("go", "test", "./"+v.Opts.Dir, "-cover"),
	)
	p := pipe.Line(
		script,
	)
	out, err := pipe.Output(p)
	if err != nil {
		return string(out), err
	}
	if strings.Contains(string(out), "no statements") {
		organism.Fitness = maxTestScore
		return fmt.Sprintf("%f", maxTestScore), err
	}
	splitted := strings.Split(string(out), "coverage: ")

	// We don't have coverage just return
	if len(splitted) == 1 {
		organism.Fitness = 0
		return "0", nil
	}
	result := strings.Split(splitted[1], "%")[0]
	const bitSize = 64
	f, err := strconv.ParseFloat(result, bitSize)
	if err != nil {
		return "", err
	}
	organism.Fitness = f
	return fmt.Sprintf("%f", f), nil
}

const maxTestScore float64 = 100
