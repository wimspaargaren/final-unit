package tmplexec

const assertTemplate = `// Code generated by finalunit {{.Version}}, visit us at https://github.com/wimspaargaren/final-unit
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
{{ if $testCase.HasChan }}
wg := sync.WaitGroup{}
wg.Add(1)
{{ end }}
{{range  $testCase.Stmts}}	{{ . }}
{{end}}
{{/* If run time info reported that a function may panic wrap it in a Panics func */}}
{{ if $testCase.RunTimeInfo.Panics }}
s.Panics(func(){
	{{ $testCase.FuncStmt }}
})
{{/* If run time detected valid use normal assert */}}
{{ else if $testCase.RunTimeInfo.IsValid }}
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

{{range  $testCase.RunTimeInfo.AssertStmts}}{{ . }}
{{end}}
{{/* Ensure values are always used */}}
{{range  $testCase.ResultUsageStmts}}{{ . }}
{{end}}
{{ if $testCase.HasChan }}
}()
{{range  $testCase.ChanIdents}}	close({{ . }})
{{end}}
// Wait until function is executed
wg.Wait()
{{ end }}
{{/* If not valid add FIXME comment */}}
{{ else }}
// FIXME: non deterministic results detected, please add assert statements manually
{{ $testCase.FuncStmt }}
{{end}}
}

{{end}}
{{ end }}

func Test{{.SuiteName}}Suite(t *testing.T) {
	suite.Run(t, new({{.SuiteName}}Suite))
}

`