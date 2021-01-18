package tmplexec

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/gen"
)

// Opts opts for template generation and execution
type Opts struct {
	Dir      string
	Override bool
}

// IExecutor interface for executor of template
type IExecutor interface {
	Execute(organism *gen.Organism) (string, error)
}

func generateFileFromTemplate(organism *gen.Organism, templateString string) error {
	for _, f := range organism.Files {
		ext := filepath.Ext(f.FileName)
		filePath := filepath.Join(f.Dir, f.FileName)
		executionPath := strings.TrimSuffix(filePath, ext) + "_test.go"
		tmpl, err := template.New("").Funcs(template.FuncMap{
			"add": func(x int) int {
				return x + 1
			},
		}).Parse(templateString)
		if err != nil {
			return err
		}
		file, err := os.Create(executionPath)
		if err != nil {
			return err
		}
		defer func() {
			err := file.Close()
			if err != nil {
				log.WithError(err).Error("unable to close file")
			}
		}()
		err = tmpl.Execute(file, f)
		if err != nil {
			return err
		}
	}
	return nil
}
