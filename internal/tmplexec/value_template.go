package tmplexec

const valueTemplate = `// Value template
package {{.PackageName}}

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type {{.SuiteName}}Suite struct {
	suite.Suite
}

{{/* assign file to var for usage inside loop  */}}
{{ $test := .}}

{{ range $funcName, $testCases := .TestCases }}
{{/* range test cases */}}
{{range $index, $testCase := $testCases}}
{{/* Print declarations */}}
{{range  $testCase.Decls}}
{{ . }}
{{end}}
{{/* Print functions */}}

func (s *{{$test.SuiteName}}Suite) Test{{ $funcName }}{{  $index }}(){
defer func() {
		if r := recover(); r != nil {
		fmt.Println("<START;{{ $funcName }}{{ $index }}>")
		fmt.Println("Recovered in Test{{ $funcName }}{{  $index }}", r)
		fmt.Println("<END;{{ $funcName }}{{ $index }}>")
	}
}()
{{ if $testCase.HasChan }}
wg := sync.WaitGroup{}
wg.Add(1)
{{ end }}
{{range  $testCase.Stmts}}	{{ . }}
{{end}}
{{ if $testCase.HasChan }}
go func(){
	defer func() {
		if r := recover(); r != nil {
		fmt.Println("Recovered in Test{{ $funcName }}{{  $index }}", r)
	}
	defer wg.Done()
	}()
{{ end }}
{{ if $testCase.HasPrintStmts }}
{{ $testCase.FuncPrintStmt }}
{{ else }}
{{ $testCase.FuncStmt }}
{{ end }}
fmt.Println("<START;{{ $funcName }}{{  $index }}>")
{{range  $testCase.ResultStmts}}	 {{ . }}
{{end}}
fmt.Println("<END;{{ $funcName }}{{  $index }}>")
{{ if $testCase.HasChan }}
}()
{{range  $testCase.ChanIdents}}	close({{ . }})
{{end}}
// Wait until function is executed
wg.Wait()
{{ end }}
}

{{end}}
{{ end }}

func Test{{.SuiteName}}Suite(t *testing.T) {
	suite.Run(t, new({{.SuiteName}}Suite))
}

`
