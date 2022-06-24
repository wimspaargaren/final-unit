package gen

import (
	"fmt"
	"testing"

	"github.com/asherascout/final-unit/internal/ident"
	"github.com/asherascout/final-unit/internal/testcase"
	"github.com/asherascout/final-unit/pkg/values"
	"github.com/asherascout/final-unit/pkg/variables"
	"github.com/stretchr/testify/suite"
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
		Name            string
		Path            string
		ValuesMockSetup func() *values.GenMock
		TestResults     []OutputTestResult
	}{
		{
			Name: "basic types",
			Path: "../../test/data/outputs/example_base_types",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				return genMock
			},
			TestResults: []OutputTestResult{
				{
					Func: "IntFunc",
					ResStmts: []string{
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `alfa`, alfa)",
						"fmt.Println(\"\")",
					},
					FuncStmt: "alfa := IntFunc()",
				},
				{
					Func:     "BoolFunc",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `bool`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := BoolFunc()",
				},
				{
					Func:     "StringFunc",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": %#v}`, `string`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := StringFunc()",
				},
				{
					Func:     "Float32Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `float32`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := Float32Func()",
				},
				{
					Func:     "Float64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `float64`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := Float64Func()",
				},
				{
					Func:     "UIntFunc",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := UIntFunc()",
				},
				{
					Func:     "UInt8Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint8`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := UInt8Func()",
				},
				{
					Func:     "UInt16Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint16`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := UInt16Func()",
				},
				{
					Func:     "UInt32Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint32`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := UInt32Func()",
				},
				{
					Func:     "UInt64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uint64`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := UInt64Func()",
				},
				{
					Func:     "UIntPtr64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `uintptr`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := UIntPtr64Func()",
				},
				{
					Func:     "Int8Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int8`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := Int8Func()",
				},
				{
					Func:     "Int16Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int16`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := Int16Func()",
				},
				{
					Func:     "Int32Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int32`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := Int32Func()",
				},
				{
					Func:     "Int64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int64`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := Int64Func()",
				},
				{
					Func:     "RuneFunc",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `rune`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := RuneFunc()",
				},
				{
					Func:     "Complex64Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `complex64`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := Complex64Func()",
				},
				{
					Func:     "Complex128Func",
					ResStmts: []string{"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `complex128`, `alfa`, alfa)", "fmt.Println(\"\")"},
					FuncStmt: "alfa := Complex128Func()",
				},
			},
		},
		{
			Name: "ignore types",
			Path: "../../test/data/outputs/ignore_type",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				return genMock
			},
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
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				return genMock
			},
			TestResults: []OutputTestResult{
				{
					Func:     "ArrayFunc",
					ResStmts: []string{"for alfa := 0; alfa < len(alfa); alfa++ {\n\tfmt.Printf(`{ \"type\": \"%s\", \"arr_ident\": \"%s\", \"var_name\": \"%s\", \"val\": \"%+v\", \"child\": `, `arr`, `alfa`, `alfa`, alfa)\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `alfa[alfa]`, alfa[alfa])\n\tfmt.Printf(`}`)\n\tfmt.Println(\"\")\n}"},
					FuncStmt: "alfa := ArrayFunc()",
				},
			},
		},
		{
			Name: "pointer types",
			Path: "../../test/data/outputs/pointer",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				return genMock
			},
			TestResults: []OutputTestResult{
				{
					Func: "PointerFunc",
					ResStmts: []string{
						"if alfa == nil {\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"nil\" } `, `pointer`, `alfa`)\n\tfmt.Println(\"\")\n} else {\n\talfa := *alfa\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `pointer`, `alfa`)\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `alfa`, alfa)\n\tfmt.Printf(`}`)\n\tfmt.Println(\"\")\n}",
					},
					FuncStmt: "alfa := PointerFunc()",
				},
			},
		},
		{
			Name: "error types",
			Path: "../../test/data/outputs/errors",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				return genMock
			},
			TestResults: []OutputTestResult{
				{
					Func: "ErrorFunc",
					ResStmts: []string{
						"if alfa == nil {\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"nil\" } `, `error`, `alfa`)\n\tfmt.Println(\"\")\n} else {\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"notnil\" } `, `error`, `alfa`)\n\tfmt.Println(\"\")\n}",
					},
					FuncStmt: "alfa := ErrorFunc()",
				},
			},
		},
		{
			Name: "map types",
			Path: "../../test/data/outputs/maps",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("MapLen").Return(2)
				genMock.On("Int").Return("2")
				genMock.On("String").Return("2")
				return genMock
			},
			TestResults: []OutputTestResult{
				{
					Func: "MapFunc",
					ResStmts: []string{
						"for alfa := range alfa {\n\tfmt.Printf(`{ \"type\": \"%s\", \"arr_ident\": \"%s\", \"map_key_type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%+v\", \"child\": `, `map`, `alfa`, `int`, `alfa`, alfa)\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": %#v}`, `string`, `alfa[alfa]`, alfa[alfa])\n\tfmt.Printf(`}`)\n\tfmt.Println(\"\")\n}",
					},
					FuncStmt: "alfa := MapFunc(alfa)",
				},
				{
					Func:     "MapUnSupportedKeyFunc",
					ResStmts: []string{},
					FuncStmt: "_ = MapUnSupportedKeyFunc(alfa)",
				},
				{
					Func:     "MapUnSupportedValFunc",
					ResStmts: []string{},
					FuncStmt: "_ = MapUnSupportedValFunc(alfa)",
				},
			},
		},
		{
			Name: "custom types",
			Path: "../../test/data/outputs/custom_types",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				return genMock
			},
			TestResults: []OutputTestResult{
				{
					Func: "CustomTypeFunc",
					ResStmts: []string{
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `custom`, `X`)",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `alfa`, alfa)",
						"fmt.Printf(`}`)",
						"fmt.Println(\"\")",
					},
					FuncStmt: "alfa := CustomTypeFunc()",
				},
			},
		},
		{
			Name: "struct types",
			Path: "../../test/data/outputs/struct",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				return genMock
			},
			TestResults: []OutputTestResult{
				{
					Func: "StructFunc",
					ResStmts: []string{
						"_ = alfa",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `alfa`)",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `alfa.X`, alfa.X)",
						"fmt.Printf(`}`)",
						"fmt.Println(\"\")",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `alfa`)",
						"fmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": %#v}`, `string`, `alfa.y`, alfa.y)",
						"fmt.Printf(`}`)",
						"fmt.Println(\"\")",
					},
					FuncStmt: "alfa := StructFunc()",
				},
				{
					Func: "StructCustomTypeDef",
					ResStmts: []string{
						"if alfa == nil {\n\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"nil\" } `, `pointer`, `alfa`)\n\tfmt.Println(\"\")\n} else {\n\talfa := *alfa\n\t_ = alfa\n\tfor alfa := 0; alfa < len(alfa.X); alfa++ {\n\t\t_ = alfa.X[alfa]\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `pointer`, `alfa`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `alfa`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"arr_ident\": \"%s\", \"var_name\": \"%s\", \"val\": \"%+v\", \"child\": `, `arr`, `alfa`, `alfa.X`, alfa)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `alfa.X[alfa]`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": \"%#v\"}`, `int`, `alfa.X[alfa].X`, alfa.X[alfa].X)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Println(\"\")\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `pointer`, `alfa`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `alfa`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"arr_ident\": \"%s\", \"var_name\": \"%s\", \"val\": \"%+v\", \"child\": `, `arr`, `alfa`, `alfa.X`, alfa)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"child\": `, `struct`, `alfa.X[alfa]`)\n\t\tfmt.Printf(`{ \"type\": \"%s\", \"var_name\": \"%s\", \"val\": %#v}`, `string`, `alfa.X[alfa].y`, alfa.X[alfa].y)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Printf(`}`)\n\t\tfmt.Println(\"\")\n\t}\n}",
					},
					FuncStmt: "alfa := StructCustomTypeDef()",
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			iGen := test.ValuesMockSetup()
			opts := &Options{
				OrganismAmount:   1,
				MaxRecursion:     3,
				TestCasesPerFunc: 1,
				ValGenerator:     iGen,
				VarGenerator:     variables.NewMock(),
				IdentGen:         ident.NewMock(),
			}
			generator, err := New(test.Path, "", opts)
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
