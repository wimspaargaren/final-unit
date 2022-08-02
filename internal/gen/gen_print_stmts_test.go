package gen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wimspaargaren/final-unit/internal/testcase"
	"github.com/wimspaargaren/final-unit/pkg/seed"
)

type OutputTestResult struct {
	Func     string
	ResStmts []string
	FuncStmt string
}

type AssignStmtGeneratorSuite struct {
	suite.Suite
}

func (s *AssignStmtGeneratorSuite) TestExampleOutputs() {
	tests := []struct {
		Name        string
		Path        string
		TestResults []OutputTestResult
	}{
		{
			Name: "basic types",
			Path: "../../test/data/outputs/example_base_types",
			TestResults: []OutputTestResult{
				{
					Func: "IntFunc",
					ResStmts: []string{
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `out`, out)",
						"fmt.Println(\"\")",
					},
					FuncStmt: "out := IntFunc()",
				},
				{
					Func:     "BoolFunc",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `bool`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := BoolFunc()",
				},
				{
					Func:     "StringFunc",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": %#v}`, `string`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := StringFunc()",
				},
				{
					Func:     "Float32Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `float32`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := Float32Func()",
				},
				{
					Func:     "Float64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `float64`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := Float64Func()",
				},
				{
					Func:     "UIntFunc",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := UIntFunc()",
				},
				{
					Func:     "UInt8Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint8`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := UInt8Func()",
				},
				{
					Func:     "UInt16Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint16`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := UInt16Func()",
				},
				{
					Func:     "UInt32Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint32`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := UInt32Func()",
				},
				{
					Func:     "UInt64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint64`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := UInt64Func()",
				},
				{
					Func:     "UIntPtr64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uintptr`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := UIntPtr64Func()",
				},
				{
					Func:     "Int8Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int8`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := Int8Func()",
				},
				{
					Func:     "Int16Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int16`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := Int16Func()",
				},
				{
					Func:     "Int32Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int32`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := Int32Func()",
				},
				{
					Func:     "Int64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int64`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := Int64Func()",
				},
				{
					Func:     "RuneFunc",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `rune`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := RuneFunc()",
				},
				{
					Func:     "Complex64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `complex64`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := Complex64Func()",
				},
				{
					Func:     "Complex128Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `complex128`, `out`, out)", "fmt.Println(\"\")"},
					FuncStmt: "out := Complex128Func()",
				},
			},
		},
		{
			Name: "ignore types",
			Path: "../../test/data/outputs/ignore_type",
			TestResults: []OutputTestResult{
				{
					Func:     "IgnoreChan",
					ResStmts: []string{},
					FuncStmt: "_ = IgnoreChan()",
				},
				{
					Func:     "IgnoreFunc",
					ResStmts: []string{},
					FuncStmt: "_ = IgnoreFunc()",
				},
				{
					Func:     "IgnoreInterfaceEmpty",
					ResStmts: []string{},
					FuncStmt: "_ = IgnoreInterfaceEmpty()",
				},
				{
					Func:     "IgnoreInterface",
					ResStmts: []string{},
					FuncStmt: "_ = IgnoreInterface()",
				},
				{
					Func:     "FuncReturnFunc",
					ResStmts: []string{},
					FuncStmt: "_ = FuncReturnFunc()",
				},
			},
		},
		{
			Name: "array types",
			Path: "../../test/data/outputs/example_array",
			TestResults: []OutputTestResult{
				{
					Func:     "ArrayFunc",
					ResStmts: []string{"for xVlBz := 0; xVlBz < len(out); xVlBz++ {\n\tfmt.Printf(`{ \"type\": \"%s\", \"arr_ident\": \"%s\", \"var_name\": \"%s\", \"val\": \"%+v\", \"child\": `, `arr`, `xVlBz`, `out`, xVlBz)\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `out[xVlBz]`, out[xVlBz])\n\tfmt.Printf(`}`)\n\tfmt.Println(\"\")\n}"},
					FuncStmt: "out := ArrayFunc()",
				},
			},
		},
		{
			Name: "pointer types",
			Path: "../../test/data/outputs/pointer",
			TestResults: []OutputTestResult{
				{
					Func: "PointerFunc",
					ResStmts: []string{
						"if out == nil {\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"nil\" } `, `pointer`, `out`)\n\tfmt.Println(\"\")\n} else {\n\tpointerOut := *out\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `pointer`, `out`)\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `pointerOut`, pointerOut)\n\tfmt.Printf(`}`)\n\tfmt.Println(\"\")\n}",
					},
					FuncStmt: "out := PointerFunc()",
				},
			},
		},
		{
			Name: "error types",
			Path: "../../test/data/outputs/errors",
			TestResults: []OutputTestResult{
				{
					Func: "ErrorFunc",
					ResStmts: []string{
						"if out == nil {\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"nil\" } `, `error`, `out`)\n\tfmt.Println(\"\")\n} else {\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"notnil\" } `, `error`, `out`)\n\tfmt.Println(\"\")\n}",
					},
					FuncStmt: "out := ErrorFunc()",
				},
			},
		},
		{
			Name: "map types",
			Path: "../../test/data/outputs/maps",
			TestResults: []OutputTestResult{
				{
					Func: "MapFunc",
					ResStmts: []string{
						"for vlBzg := range out {\n\tfmt.Printf(`{ \"type\": \"%s\", \"arr_ident\": \"%s\", \"map_key_type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%+v\", \"child\": `, `map`, `vlBzg`, `int`, `out`, vlBzg)\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": %#v}`, `string`, `out[vlBzg]`, out[vlBzg])\n\tfmt.Printf(`}`)\n\tfmt.Println(\"\")\n}",
					},
					FuncStmt: "out := MapFunc(x)",
				},
				{
					Func:     "MapUnSupportedKeyFunc",
					ResStmts: []string{},
					FuncStmt: "_ = MapUnSupportedKeyFunc(x)",
				},
				{
					Func:     "MapUnSupportedValFunc",
					ResStmts: []string{},
					FuncStmt: "_ = MapUnSupportedValFunc(x)",
				},
			},
		},
		{
			Name: "custom types",
			Path: "../../test/data/outputs/custom_types",
			TestResults: []OutputTestResult{
				{
					Func: "CustomTypeFunc",
					ResStmts: []string{
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `custom`, `X`)",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `out`, out)",
						"fmt.Printf(`}`)",
						"fmt.Println(\"\")",
					},
					FuncStmt: "out := CustomTypeFunc()",
				},
			},
		},
		{
			Name: "struct types",
			Path: "../../test/data/outputs/struct",
			TestResults: []OutputTestResult{
				{
					Func: "StructFunc",
					ResStmts: []string{
						"_ = out",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `out`)",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `out.X`, out.X)",
						"fmt.Printf(`}`)",
						"fmt.Println(\"\")",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `out`)",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": %#v}`, `string`, `out.y`, out.y)",
						"fmt.Printf(`}`)",
						"fmt.Println(\"\")",
					},
					FuncStmt: "out := StructFunc()",
				},
				{
					Func: "StructCustomTypeDef",
					ResStmts: []string{
						"if out == nil {\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"nil\" } `, `pointer`, `out`)\n\tfmt.Println(\"\")\n} else {\n\tpointerOut := *out\n\t_ = pointerOut\n\tfor xVlBz := 0; xVlBz < len(pointerOut.X); xVlBz++ {\n\t\t_ = pointerOut.X[xVlBz]\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `pointer`, `out`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `pointerOut`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"arr_ident\": \"%s\", \"var_name\": \"%s\", \"val\": \"%+v\", \"child\": `, `arr`, `xVlBz`, `pointerOut.X`, xVlBz)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `pointerOut.X[xVlBz]`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `pointerOut.X[xVlBz].X`, pointerOut.X[xVlBz].X)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Println(\"\")\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `pointer`, `out`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `pointerOut`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"arr_ident\": \"%s\", \"var_name\": \"%s\", \"val\": \"%+v\", \"child\": `, `arr`, `xVlBz`, `pointerOut.X`, xVlBz)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `pointerOut.X[xVlBz]`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": %#v}`, `string`, `pointerOut.X[xVlBz].y`, pointerOut.X[xVlBz].y)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Println(\"\")\n\t}\n}",
					},
					FuncStmt: "out := StructCustomTypeDef()",
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			opts := &Options{
				OrganismAmount:   1,
				MaxRecursion:     3,
				TestCasesPerFunc: 1,
			}
			seed.SetRandomSeed(1)
			generator, err := New(test.Path, opts)
			s.Require().NoError(err)
			organisms := generator.GetTestCases()
			s.Require().Equal(1, len(organisms))
			files := organisms[0].Files
			s.Require().Equal(1, len(files))
			res := files[0].TestCases

			for _, testResult := range test.TestResults {
				s.Run(testResult.Func, func() {
					funcTestCases, ok := res[testResult.Func]
					s.Require().Equal(1, len(funcTestCases))
					funcTestCase := funcTestCases[0]
					s.Require().True(ok)
					s.Require().Equal(len(testResult.ResStmts), len(funcTestCase.ResultStmts))
					for i, stmt := range funcTestCase.ResultStmts {
						s.Equal(testResult.ResStmts[i], stmt, fmt.Sprintf("line: %d", i))
					}
					s.Equal(testResult.FuncStmt, funcTestCase.FuncPrintStmt)
				})
			}
		})
	}
}

func TestAssignStmtGeneratorSuite(t *testing.T) {
	suite.Run(t, new(AssignStmtGeneratorSuite))
}

func (s *AssignStmtGeneratorSuite) GetTestCase(files []*File, funcName string) []*testcase.TestCase {
	for _, file := range files {
		if x, ok := file.TestCases[funcName]; ok {
			return x
		}
	}
	s.Fail("test not found")
	return nil
}
