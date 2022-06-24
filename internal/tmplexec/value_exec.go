package tmplexec

import (
	"github.com/asherascout/final-unit/internal/gen"
	"gopkg.in/pipe.v2"
)

// NewValueExecutor creates new value template executor
// this is the second step to read values to assert
func NewValueExecutor(opts Opts) IExecutor {
	return &ValueExecutor{
		Opts: opts,
	}
}

// ValueExecutor implementation of IValueExecutor
type ValueExecutor struct {
	Opts Opts
}

// Execute executes organism on value template
func (v *ValueExecutor) Execute(organism *gen.Organism) (string, error) {
	err := generateFileFromTemplate(organism, valueTemplate)
	if err != nil {
		return "", err
	}
	script := pipe.Script(
		pipe.Exec("goimports", "-w", v.Opts.Dir),
		pipe.Exec("go", "clean", "-testcache"),
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
