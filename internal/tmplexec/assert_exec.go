package tmplexec

import (
	"github.com/asherascout/final-unit/internal/gen"
	"gopkg.in/pipe.v2"
)

// NewAssertExecutor creates new assert template executor
// the final generator which creates inputs & verifies output
func NewAssertExecutor(opts Opts) IExecutor {
	return &AssertExecutor{
		Opts: opts,
	}
}

// AssertExecutor implementation of IAssertExecutor
type AssertExecutor struct {
	Opts Opts
}

// Execute executes organism on assert template
func (v *AssertExecutor) Execute(organism *gen.Organism) (string, error) {
	err := generateFileFromTemplate(organism, assertTemplate)
	if err != nil {
		return "", err
	}
	script := pipe.Script(
		pipe.Exec("goimports", "-w", v.Opts.Dir),
		pipe.Exec("go", "test", v.Opts.Dir, "-v"),
	)
	p := pipe.Line(
		script,
	)
	out, err := pipe.Output(p)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
