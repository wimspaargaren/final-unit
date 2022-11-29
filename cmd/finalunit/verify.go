package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/pipe.v2"
)

// error definitions
var (
	ErrMissingGoImports = fmt.Errorf("go imports not found please install using: go get golang.org/x/tools/cmd/goimports")
	ErrNoGoFiles        = fmt.Errorf("no go files found in given directory")
)

// UnreadableDir error indicating that directory is unreadable
type UnreadableDir struct {
	Dir string
}

func (e *UnreadableDir) Error() string {
	return fmt.Sprintf("unable to read files at dir: %s", e.Dir)
}

// Verify perform additional verification before running
func Verify(opts *Opts) error {
	err := GoImportsInstalled()
	if err != nil {
		return err
	}
	return CheckGoFiles(opts)
}

// CheckGoFiles check if go files exist in current given directory
func CheckGoFiles(opts *Opts) error {
	files, err := os.ReadDir(opts.Dir)
	if err != nil {
		return &UnreadableDir{Dir: opts.Dir}
	}

	for _, f := range files {
		extension := filepath.Ext(f.Name())
		if extension == ".go" && !strings.HasSuffix(f.Name(), "_test.go") {
			return nil
		}
	}
	return ErrNoGoFiles
}

// GoImportsInstalled check if goimports is installed
func GoImportsInstalled() error {
	script := pipe.Script(
		pipe.Exec("goimports", "-w", "."),
	)
	p := pipe.Line(
		script,
	)
	_, err := pipe.Output(p)
	if err != nil {
		log.WithError(err).Debug("couldn't execute goimports --help")
		return ErrMissingGoImports
	}
	return nil
}
