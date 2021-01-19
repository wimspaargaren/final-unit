package gen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/wimspaargaren/final-unit/internal/ident"
	"github.com/wimspaargaren/final-unit/internal/testcase"
	"github.com/wimspaargaren/final-unit/pkg/values"
	"github.com/wimspaargaren/final-unit/pkg/variables"
)

type TestResult struct {
	Func     string
	ResStmts []string
	ResDecls []string
}

type PrintStmtTestSuite struct {
	suite.Suite
}

func (s *PrintStmtTestSuite) TestExamplesSingleFile() {
	tests := []struct {
		Name            string
		Path            string
		ValuesMockSetup func() *values.GenMock
		TestResults     []TestResult
	}{
		{
			Name: "ellipsis test",
			Path: "../../test/data/inputs/example_ellipsis",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("ArrayLen", mock.Anything).Return(1)
				genMock.On("String").Return(`"string"`)
				genMock.On("Int").Return(`42`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "EllipsisStringFunc",
					ResStmts: []string{
						`alfa := "string"`,
						"EllipsisStringFunc(alfa)",
					},
				},
				{
					Func: "EllipsisStructFunc",
					ResStmts: []string{
						`alfa := SomeStruct{X: 42}`,
						"EllipsisStructFunc(alfa)",
					},
				},
			},
		},
		{
			Name: "base types test",
			Path: "../../test/data/inputs/example_base_types",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("Int8").Return(`42`)
				genMock.On("Int16").Return(`42`)
				genMock.On("Int32").Return(`42`)
				genMock.On("Int64").Return(`42`)
				genMock.On("UInt").Return(`42`)
				genMock.On("UInt8").Return(`42`)
				genMock.On("UInt16").Return(`42`)
				genMock.On("UInt32").Return(`42`)
				genMock.On("UInt64").Return(`42`)
				genMock.On("UIntPtr").Return(`42`)
				genMock.On("Complex64").Return(`42`)
				genMock.On("Complex128").Return(`42`)
				genMock.On("Rune").Return(`42`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "UIntFunc",
					ResStmts: []string{`alfa := uint(42)`, "UIntFunc(alfa)"},
				},
				{
					Func:     "UInt8Func",
					ResStmts: []string{`alfa := uint8(42)`, "UInt8Func(alfa)"},
				},
				{
					Func:     "UInt16Func",
					ResStmts: []string{`alfa := uint16(42)`, "UInt16Func(alfa)"},
				},
				{
					Func:     "UInt32Func",
					ResStmts: []string{`alfa := uint32(42)`, "UInt32Func(alfa)"},
				},
				{
					Func:     "UInt64Func",
					ResStmts: []string{`alfa := uint64(42)`, "UInt64Func(alfa)"},
				},
				{
					Func:     "Int8Func",
					ResStmts: []string{`alfa := int8(42)`, "Int8Func(alfa)"},
				},
				{
					Func:     "Int16Func",
					ResStmts: []string{`alfa := int16(42)`, "Int16Func(alfa)"},
				},
				{
					Func:     "Int32Func",
					ResStmts: []string{`alfa := int32(42)`, "Int32Func(alfa)"},
				},
				{
					Func:     "Int64Func",
					ResStmts: []string{`alfa := int64(42)`, "Int64Func(alfa)"},
				},
				{
					Func:     "RuneFunc",
					ResStmts: []string{`alfa := rune(42)`, "RuneFunc(alfa)"},
				},
				{
					Func:     "Complex64Func",
					ResStmts: []string{`alfa := complex64(42)`, "Complex64Func(alfa)"},
				},
				{
					Func:     "Complex128Func",
					ResStmts: []string{`alfa := complex128(42)`, "Complex128Func(alfa)"},
				},
			},
		},
		{
			Name: "int test",
			Path: "../../test/data/inputs/example_int",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("ArrayLen", mock.Anything).Return(1)
				genMock.On("Int").Return("42")
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "IntFunc",
					ResStmts: []string{"alfa := 42", "IntFunc(alfa)"},
				},
				{
					Func:     "IntPointerFunc",
					ResStmts: []string{"alfaalfa := 42", "alfa := &alfaalfa", "IntPointerFunc(alfa)"},
				},
				{
					Func:     "IntArrayFunc",
					ResStmts: []string{"alfa := []int{42}", "IntArrayFunc(alfa)"},
				},
				{
					Func:     "IntPointerArrayFunc",
					ResStmts: []string{"alfaalfa := 42", "alfa := []*int{&alfaalfa}", "IntPointerArrayFunc(alfa)"},
				},
				{
					Func:     "IntArrayPointerFunc",
					ResStmts: []string{"alfaalfa := []int{42}", "alfa := &alfaalfa", "IntArrayPointerFunc(alfa)"},
				},
				{
					Func:     "IntPointerArrayPointerFunc",
					ResStmts: []string{"alfaalfaalfa := 42", "alfaalfa := []*int{&alfaalfaalfa}", "alfa := &alfaalfa", "IntPointerArrayPointerFunc(alfa)"},
				},
				{
					Func:     "IntArrayLenFunc",
					ResStmts: []string{"alfa := [2]int{42}", "IntArrayLenFunc(alfa)"},
				},
				{
					Func:     "IntDoubleArrayFunc",
					ResStmts: []string{"alfa := [][]int{[]int{42}}", "IntDoubleArrayFunc(alfa)"},
				},
				{
					Func:     "IntDoubleArrayLenFunc",
					ResStmts: []string{"alfa := [][2]int{[2]int{42}}", "IntDoubleArrayLenFunc(alfa)"},
				},
			},
		},
		{
			Name: "unnamed test",
			Path: "../../test/data/inputs/example_unnamed",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return("42")
				genMock.On("String").Return(`"string"`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "SimpleUnnamed",
					ResStmts: []string{
						"alfa := Simple{X: struct{ Y string }{Y: \"string\"}, Z: Other{Y: \"string\"}}",
						"SimpleUnnamed(alfa)",
					},
				},
				{
					Func: "SimpleInterfaceUnnamed",
					ResStmts: []string{
						"alfa := SimpleInterface{X: &Alfa{}}",
						"SimpleInterfaceUnnamed(alfa)",
					},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) Hi(x int) {\n\treturn\n}",
					},
				},
			},
		},
		{
			Name: "nset struct test",
			Path: "../../test/data/inputs/example_directly_nested_struct",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return("42")
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "NormalNestFunc",
					ResStmts: []string{"alfa := Normal{X: X(42)}", "NormalNestFunc(alfa)"},
				},
				{
					Func:     "NormalPointerNestFunc",
					ResStmts: []string{"xalfa := X(42)", "alfa := PointerNormal{X: &xalfa}", "NormalPointerNestFunc(alfa)"},
				},
				{
					Func:     "NestedImportFunc",
					ResStmts: []string{"alfa := ImportNested{Hello: nestedimport.Hello{X: 42}}", "NestedImportFunc(alfa)"},
				},
				{
					Func:     "NestedImportPointerFunc",
					ResStmts: []string{"helloalfa := nestedimport.Hello{X: 42}", "alfa := ImportPointerNested{Hello: &helloalfa}", "NestedImportPointerFunc(alfa)"},
				},
			},
		},
		{
			Name: "bool test",
			Path: "../../test/data/inputs/example_bool",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("ArrayLen", mock.Anything).Return(1)
				genMock.On("Bool").Return("true")
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "BoolFunc",
					ResStmts: []string{"alfa := true", "BoolFunc(alfa)"},
				},
				{
					Func:     "BoolPointerFunc",
					ResStmts: []string{"alfaalfa := true", "alfa := &alfaalfa", "BoolPointerFunc(alfa)"},
				},
			},
		},
		{
			Name: "string test",
			Path: "../../test/data/inputs/example_string",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("ArrayLen", mock.Anything).Return(1)
				genMock.On("String").Return(`"string"`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "StringFunc",
					ResStmts: []string{
						`alfa := "string"`,
						"StringFunc(alfa)",
					},
				},
			},
		},
		{
			Name: "multi param test",
			Path: "../../test/data/inputs/example_multi_param_same_name",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("String").Return(`"string"`)
				genMock.On("Int").Return(`42`)
				genMock.On("ArrayLen", mock.Anything).Return(1)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "MultiParamFunc",
					ResStmts: []string{`alfa := "string"`, `alfa := "string"`, "MultiParamFunc(alfa, alfa)"},
				},
				{
					Func:     "MultiParamPointerFunc",
					ResStmts: []string{`alfaalfa := "string"`, "alfa := &alfaalfa", `alfaalfa := "string"`, "alfa := &alfaalfa", "MultiParamPointerFunc(alfa, alfa)"},
				},
				{
					Func:     "MultiParamDifTypeFunc",
					ResStmts: []string{"alfa := 42", `alfa := "string"`, "MultiParamDifTypeFunc(alfa, alfa)"},
				},
			},
		},
		{
			Name: "float test",
			Path: "../../test/data/inputs/example_float",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Float64").Return(`64.0`)
				genMock.On("Float32").Return(`32.0`)
				genMock.On("ArrayLen", mock.Anything).Return(1)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "Float32Func",
					ResStmts: []string{`alfa := float32(32.0)`, "Float32Func(alfa)"},
				},
				{
					Func:     "Float64Func",
					ResStmts: []string{`alfa := 64.0`, "Float64Func(alfa)"},
				},
			},
		},
		{
			Name: "byte test",
			Path: "../../test/data/inputs/example_byte",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Byte").Return(`42`)
				genMock.On("ArrayLen", mock.Anything).Return(1)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "ByteFunc",
					ResStmts: []string{`alfa := byte(42)`, "ByteFunc(alfa)"},
				},
				{
					Func:     "BytePointerFunc",
					ResStmts: []string{`alfaalfa := byte(42)`, `alfa := &alfaalfa`, "BytePointerFunc(alfa)"},
				},
				{
					Func:     "ByteArrayFunc",
					ResStmts: []string{`alfa := []byte{byte(42)}`, "ByteArrayFunc(alfa)"},
				},
				{
					Func:     "ByteArrayPointerFunc",
					ResStmts: []string{"alfaalfa := byte(42)", `alfa := []*byte{&alfaalfa}`, "ByteArrayPointerFunc(alfa)"},
				},
			},
		},
		{
			Name: "struct test",
			Path: "../../test/data/inputs/example_struct",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("String").Return(`"string"`)
				genMock.On("ArrayLen", mock.Anything).Return(1)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "StructFuncSimple",
					ResStmts: []string{`alfa := Simple{Var: 42}`, "StructFuncSimple(alfa)"},
				},
				{
					Func:     "StructFuncComposite",
					ResStmts: []string{`varStringalfa := "string"`, `alfa := Composite{VarInt: 42, VarString: &varStringalfa}`, "StructFuncComposite(alfa)"},
				},
				{
					Func:     "StructFuncCompositeOther",
					ResStmts: []string{`alfa := CompositeOther{x: 42, y: 42}`, "StructFuncCompositeOther(alfa)"},
				},
				{
					Func:     "StructFuncNested",
					ResStmts: []string{`alfa := Nested{X: Simple{Var: 42}}`, "StructFuncNested(alfa)"},
				},
				{
					Func:     "StructFuncNestedNested",
					ResStmts: []string{"alfa := NestedNested{Nested: Nested{X: Simple{Var: 42}}, x: 42}", "StructFuncNestedNested(alfa)"},
				},
				{
					Func:     "StructFuncNestedInt",
					ResStmts: []string{"alfa := NestedInt{int: 42, x: \"string\"}", "StructFuncNestedInt(alfa)"},
				},
				{
					Func:     "StructFuncNestedCustomType",
					ResStmts: []string{"alfa := NestedCustomType{CustomType: CustomType(func(x int) string {\n\talfa := \"string\"\n\treturn alfa\n}), x: \"string\"}", "StructFuncNestedCustomType(alfa)"},
				},
			},
		},
		{
			Name: "custom type test",
			Path: "../../test/data/inputs/example_custom_type",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("Byte").Return(`42`)
				genMock.On("ArrayLen", mock.Anything).Return(1)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "CustomTypeSimple",
					ResStmts: []string{`alfa := CustomIntArray([]int{42})`, "CustomTypeSimple(alfa)"},
				},
				{
					Func:     "CustomTypeStructType",
					ResStmts: []string{`alfa := StructType(Test{X: 42})`, "CustomTypeStructType(alfa)"},
				},
				{
					Func:     "CustomByteArray",
					ResStmts: []string{`alfa := UUID([8]byte{byte(42)})`, "CustomByteArray(alfa)"},
				},
			},
		},
		{
			Name: "array test",
			Path: "../../test/data/inputs/example_array",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("ArrayLen", 0).Return(0)
				genMock.On("ArrayLen", 5).Return(5)
				genMock.On("ArrayLen", -1).Return(3)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "ArrayLenZero",
					ResStmts: []string{`alfa := [0]int{}`, "ArrayLenZero(alfa)"},
				},
				{
					Func:     "ArrayLenFive",
					ResStmts: []string{`alfa := [5]int{42, 42, 42, 42, 42}`, "ArrayLenFive(alfa)"},
				},
				{
					Func:     "ArrayNoLen",
					ResStmts: []string{`alfa := []int{42, 42, 42}`, "ArrayNoLen(alfa)"},
				},
				{
					Func:     "ArrayStructNoLen",
					ResStmts: []string{`alfa := []Simple{Simple{X: 42}, Simple{X: 42}, Simple{X: 42}}`, "ArrayStructNoLen(alfa)"},
				},
				{
					Func: "ArrayPointerStructNoLen",
					ResStmts: []string{
						"xalfa := 42",
						"xalfa := 42",
						"xalfa := 42",
						"alfaalfa := []Pointer{Pointer{x: &xalfa}, Pointer{x: &xalfa}, Pointer{x: &xalfa}}",
						`alfa := &alfaalfa`,
						"ArrayPointerStructNoLen(alfa)",
					},
				},
			},
		},
		{
			Name: "func on struct",
			Path: "../../test/data/inputs/example_func_on_struct",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "SimpleSimpleFuncOnStruct",
					ResStmts: []string{
						"alfa := Simple{X: 42}",
						`alfa := 42`,
						`alfa.SimpleFuncOnStruct(alfa)`,
					},
				},
				{
					Func: "CustomTypeFuncOnCustomType",
					ResStmts: []string{
						"alfa := CustomType(42)",
						`alfa := 42`,
						`alfa.FuncOnCustomType(alfa)`,
					},
				},
			},
		},
		{
			Name: "map test",
			Path: "../../test/data/inputs/example_map",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("String").Return(`"string"`)
				genMock.On("MapLen").Return(2)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "MapFunc",
					ResStmts: []string{
						"alfa := map[int]string{42: \"string\"}",
						`MapFunc(alfa)`,
					},
				},
				{
					Func: "MapFuncNested",
					ResStmts: []string{
						"alfa := map[int]map[int]string{42: map[int]string{42: \"string\"}}",
						`MapFuncNested(alfa)`,
					},
				},
			},
		},
		{
			Name: "interface import test",
			Path: "../../test/data/inputs/example_interface_import",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "SomeFuncImportingInterface",
					ResStmts: []string{
						"alfa := &Alfa{}",
						`SomeFuncImportingInterface(alfa)`,
					},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) SomeFunc(someimport.Y) someimport.Y {\n\talfa := someimport.Y{X: 42}\n\treturn alfa\n}",
					},
				},
			},
		},
		{
			Name: "interface private return test",
			Path: "../../test/data/inputs/example_priv_ret_interface",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("UInt").Return(`42`)
				genMock.On("UInt32").Return(`42`)
				genMock.On("ArrayLen", mock.Anything).Return(42)
				genMock.On("String").Return(`"string"`)
				genMock.On("Bool").Return(`true`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "ReflectType",
					ResStmts: []string{"alfa := func() reflect.Type {\n\treturn nil\n}()", "ReflectType(alfa)"},
					ResDecls: []string{},
				},
				{
					Func:     "AstSelection",
					ResStmts: []string{"alfa := func() ast.Selection {\n\treturn nil\n}()", "AstSelection(alfa)"},
					ResDecls: []string{},
				},
			},
		},
		{
			Name: "error test",
			Path: "../../test/data/inputs/example_error",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Error").Return(true)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "ErrorFunc",
					ResStmts: []string{"alfa := func() error {\n\treturn fmt.Errorf(\"very error\")\n}()", "ErrorFunc(alfa)"},
				},
				{
					Func:     "ErrorPointerFunc",
					ResStmts: []string{"alfaalfa := func() error {\n\treturn fmt.Errorf(\"very error\")\n}()", `alfa := &alfaalfa`, "ErrorPointerFunc(alfa)"},
				},
			},
		},
		{
			Name: "interface test",
			Path: "../../test/data/inputs/example_interface",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("String").Return(`"string"`)
				genMock.On("Error").Return(true)
				genMock.On("Int").Return(`42`)
				genMock.On("Type").Return(`int`)
				genMock.On("Byte").Return(`42`)
				genMock.On("ArrayLen", mock.Anything).Return(2)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "InterfaceFuncEmpty",
					ResStmts: []string{`alfa := 42`, "InterfaceFuncEmpty(alfa)"},
				},
				{
					Func:     "InterfaceFuncSimple",
					ResStmts: []string{`alfa := &Alfa{}`, "InterfaceFuncSimple(alfa)"},
					ResDecls: []string{"type Alfa struct {\n}", "func (s *Alfa) Hello(x int) {\n\treturn\n}"},
				},
				{
					Func:     "InterfaceFuncSimpleWithReturn",
					ResStmts: []string{`alfa := &Alfa{}`, "InterfaceFuncSimpleWithReturn(alfa)"},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) Hello(x int) int {\n\talfa := 42\n\treturn alfa\n}",
					},
				},
				{
					Func:     "InterfaceComplex",
					ResStmts: []string{`alfa := &Alfa{}`, "InterfaceComplex(alfa)"},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) Hello(x *int) [2]byte {\n\talfa := [2]byte{byte(42), byte(42)}\n\treturn alfa\n}",
						"func (s *Alfa) World(x Simple) (*Some, error) {\n\talfaalfa := Some{x: \"string\"}\n\talfa := &alfaalfa\n\talfa := func() error {\n\t\treturn fmt.Errorf(\"very error\")\n\t}()\n\treturn alfa, alfa\n}",
					},
				},
				{
					Func:     "InterfaceNested",
					ResStmts: []string{`alfa := &Alfa{}`, "InterfaceNested(alfa)"},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) Hello(x int) {\n\treturn\n}",
						"func (s *Alfa) World(x int) int {\n\talfa := 42\n\treturn alfa\n}",
					},
				},
			},
		},
		{
			Name: "interface 2 test",
			Path: "../../test/data/inputs/example_interface_2",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func:     "NestedDouble",
					ResStmts: []string{`alfa := &Alfa{}`, "NestedDouble(alfa)"},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) Hello(int) {\n\treturn\n}",
						"func (s *Alfa) World(int) {\n\treturn\n}",
					},
				},
			},
		},
		{
			Name: "cant gen func test",
			Path: "../../test/data/inputs/example_func_return_priv",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "Test",
					ResStmts: []string{
						"alfa := exampleimport.SomeFunc(nil)",
						`Test(alfa)`,
					},
				},
				{
					Func: "Test2",
					ResStmts: []string{
						"alfa := exampleimport.Other{}",
						`Test2(alfa)`,
					},
				},
			},
		},
		{
			Name: "chan test",
			Path: "../../test/data/inputs/example_chan",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "FuncChan",
					ResStmts: []string{
						"alfa := make(chan int)",
						"alfa := alfa",
						`FuncChan(alfa)`,
					},
				},
			},
		},
		{
			Name: "cycle test",
			Path: "../../test/data/inputs/example_cycle",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("Error").Return(false)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "FuncCycle",
					ResStmts: []string{
						"xalfa := A{}",
						"xalfa := B{X: &xalfa, Y: 42}",
						"xalfa := A{X: &xalfa, Y: 42}",
						"xalfa := B{X: &xalfa, Y: 42}",
						"xalfa := A{X: &xalfa, Y: 42}",
						"xalfa := B{X: &xalfa, Y: 42}",
						"alfa := A{X: &xalfa, Y: 42}",
						`FuncCycle(alfa)`,
					},
				},
				{
					Func: "CycleInterface",
					ResStmts: []string{
						"alfa := &Alfa{}",
						`CycleInterface(alfa)`,
					},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"type Alfa struct {\n}",
						"type Alfa struct {\n}",
						"func (s *Alfa) X() X {\n\talfa := &Alfa{}\n\treturn alfa\n}",
						"func (s *Alfa) X() X {\n\talfa := &Alfa{}\n\treturn alfa\n}",
						"func (s *Alfa) X() X {\n\talfa := &Alfa{}\n\treturn alfa\n}",
					},
				},
				{
					Func: "CycleComplicated",
					ResStmts: []string{
						"zalfa := somepkg.A{}",
						"balfa := somepkg.B{X: func() (io.ReadCloser, error) {\n\talfa := &Alfa{}\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa, alfa\n}, Y: &Alfa{}, Z: &zalfa}",
						"zalfa := somepkg.A{B: &balfa}",
						"balfa := somepkg.B{X: func() (io.ReadCloser, error) {\n\talfa := &Alfa{}\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa, alfa\n}, Y: &Alfa{}, Z: &zalfa}",
						"zalfa := somepkg.A{B: &balfa}",
						"balfa := somepkg.B{X: func() (io.ReadCloser, error) {\n\talfa := &Alfa{}\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa, alfa\n}, Y: &Alfa{}, Z: &zalfa}",
						"alfa := somepkg.A{B: &balfa}",
						`CycleComplicated(alfa)`,
					},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) Read(p []byte) (n int, err error) {\n\talfa := 42\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa, alfa\n}",
						"func (s *Alfa) Close() error {\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa\n}",
						"type Alfa struct {\n}",
						"func (s *Alfa) Read(p []byte) (n int, err error) {\n\talfa := 42\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa, alfa\n}",
						"func (s *Alfa) Close() error {\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa\n}",
						"type Alfa struct {\n}",
						"func (s *Alfa) Read(p []byte) (n int, err error) {\n\talfa := 42\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa, alfa\n}",
						"func (s *Alfa) Close() error {\n\talfa := func() error {\n\t\treturn nil\n\t}()\n\treturn alfa\n}",
					},
				},
			},
		},
		{
			Name: "func func test",
			Path: "../../test/data/inputs/example_func",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "FuncFunc",
					ResStmts: []string{
						"alfa := func(x int) {\n\treturn\n}",
						`FuncFunc(alfa)`,
					},
				},
				{
					Func: "FuncTypeFunc",
					ResStmts: []string{
						"alfa := FuncType(func(x int) {\n\treturn\n})",
						`FuncTypeFunc(alfa)`,
					},
				},
				{
					Func: "FuncFuncWithReturn",
					ResStmts: []string{
						"alfa := func(x int) int {\n\talfa := 42\n\treturn alfa\n}",
						`FuncFuncWithReturn(alfa)`,
					},
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			iGen := test.ValuesMockSetup()
			opts := &Options{
				MaxRecursion:     3,
				OrganismAmount:   1,
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
					s.Require().Equal(1, len(funcTestCases), fmt.Sprintf("Func: %s", testResult.Func))
					funcTestCase := funcTestCases[0]
					s.Require().True(ok)
					s.Require().Equal(len(testResult.ResStmts)-1, len(funcTestCase.Stmts), fmt.Sprintf("Func: %s", testResult.Func))
					for i, stmt := range funcTestCase.Stmts {
						s.Equal(testResult.ResStmts[i], stmt, fmt.Sprintf("at number: %d", i))
					}
					s.Equal(testResult.ResStmts[len(testResult.ResStmts)-1], funcTestCase.FuncStmt, fmt.Sprintf("Func: %s", testResult.Func))

					s.Require().Equal(len(testResult.ResDecls), len(funcTestCase.Decls), fmt.Sprintf("Func: %s", testResult.Func))
					for i, decl := range funcTestCase.Decls {
						s.Equal(testResult.ResDecls[i], decl, fmt.Sprintf("at number: %d", i))
					}
				})
			}
		})
	}
}

func (s *PrintStmtTestSuite) TestExamplesMultiFile() {
	tests := []struct {
		Name            string
		Path            string
		ValuesMockSetup func() *values.GenMock
		TestResults     []TestResult
	}{
		{
			Name: "func multi file",
			Path: "../../test/data/inputs/example_multi_file",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("String").Return(`"string"`)
				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "MultFileStruct",
					ResStmts: []string{
						"alfa := StructFromOtherFile{X: 42, Y: \"string\"}",
						`MultFileStruct(alfa)`,
					},
				},
				{
					Func: "MultFileInterface",
					ResStmts: []string{
						"alfa := &Alfa{}",
						`MultFileInterface(alfa)`,
					},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) Method(x int) {\n\treturn\n}",
					},
				},
				{
					Func: "MultFileCustomType",
					ResStmts: []string{
						"alfa := CustomType(func(x int) {\n\treturn\n})",
						`MultFileCustomType(alfa)`,
					},
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
			s.Require().Equal(2, len(files))

			for _, testResult := range test.TestResults {
				s.Run(testResult.Func, func() {
					funcTestCases := s.GetTestCase(files, testResult.Func)
					s.Require().Equal(1, len(funcTestCases))
					funcTestCase := funcTestCases[0]
					s.Require().Equal(len(testResult.ResStmts)-1, len(funcTestCase.Stmts))
					for i, stmt := range funcTestCase.Stmts {
						s.Equal(testResult.ResStmts[i], stmt)
					}
					s.Equal(testResult.ResStmts[len(testResult.ResStmts)-1], funcTestCase.FuncStmt)

					s.Require().Equal(len(testResult.ResDecls), len(funcTestCase.Decls))
					for i, decl := range funcTestCase.Decls {
						s.Equal(testResult.ResDecls[i], decl)
					}
				})
			}
		})
	}
}

func (s *PrintStmtTestSuite) TestExampleImports() {
	tests := []struct {
		Name            string
		Path            string
		ValuesMockSetup func() *values.GenMock
		TestResults     []TestResult
	}{
		{
			Name: "import funcs",
			Path: "../../test/data/inputs/example_imports",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				genMock.On("Byte").Return(`42`)
				genMock.On("UInt64").Return(`42`)
				genMock.On("UInt8").Return(`42`)
				genMock.On("Int64").Return(`42`)
				genMock.On("String").Return(`"string"`)
				genMock.On("Bool").Return(`true`)
				genMock.On("MapLen").Return(1)
				genMock.On("ArrayLen", mock.Anything).Return(8)

				return genMock
			},
			TestResults: []TestResult{
				{
					Func: "SimpleImport",
					ResStmts: []string{
						"alfa := somepkg.SomeStruct{X: 42}",
						`SimpleImport(alfa)`,
					},
				},
				{
					Func: "NestedImport",
					ResStmts: []string{
						"alfa := nestedimportpkg.NestedStruct{X: somepkg.SomeStruct{X: 42}}",
						`NestedImport(alfa)`,
					},
				},
				{
					Func: "ImportInterface",
					ResStmts: []string{
						"alfa := &Alfa{}",
						`ImportInterface(alfa)`,
					},
					ResDecls: []string{"type Alfa struct {\n}", "func (s *Alfa) Method(x int) {\n\treturn\n}"},
				},
				{
					Func: "ImportCustomType",
					ResStmts: []string{
						"alfa := somepkg.CustomType(func(x int) {\n\treturn\n})",
						`ImportCustomType(alfa)`,
					},
				},
				{
					Func: "ImportCustomTypeUUID",
					ResStmts: []string{
						"alfa := somepkg.UUID([8]byte{byte(42), byte(42), byte(42), byte(42), byte(42), byte(42), byte(42), byte(42)})",
						`ImportCustomTypeUUID(alfa)`,
					},
				},
				{
					Func: "NestedUnitInImport",
					ResStmts: []string{
						"alfa := somepkg.NestedStructInImport{X: somepkg.Nested{X: 42}}",
						`NestedUnitInImport(alfa)`,
					},
				},
				{
					Func: "ImportTime",
					ResStmts: []string{
						"alfa := time.Time{}",
						`ImportTime(alfa)`,
					},
				},
				{
					Func: "DirectlyNestedInterface",
					ResStmts: []string{
						"alfa := &Alfa{}",
						`DirectlyNestedInterface(alfa)`,
					},
					ResDecls: []string{
						"type Alfa struct {\n}",
						"func (s *Alfa) Hello(x int) {\n\treturn\n}",
						"func (s *Alfa) World(x int) {\n\treturn\n}",
					},
				},
				{
					Func: "ImportCustomTypeInstruct",
					ResStmts: []string{
						"alfa := somepkg.AStructWithCustomType{CustomTypeInt: somepkg.CustomTypeInt(42)}",
						`ImportCustomTypeInstruct(alfa)`,
					},
				},
				{
					Func: "ImportedMapVal",
					ResStmts: []string{
						"alfa := map[string]somepkg.SomeStruct{\"string\": somepkg.SomeStruct{X: 42}}",
						`ImportedMapVal(alfa)`,
					},
				},
				{
					Func: "ImportCustomTypeInMap",
					ResStmts: []string{
						"alfa := map[string]somepkg.CustomType{\"string\": somepkg.CustomType(func(x int) {\n\treturn\n})}",
						`ImportCustomTypeInMap(alfa)`,
					},
				},
				{
					Func: "ImportForm",
					ResStmts: []string{
						"alfa := somepkg.Form{File: map[somepkg.SomeStruct][]somepkg.SomeStruct{somepkg.SomeStruct{X: 42}: []somepkg.SomeStruct{somepkg.SomeStruct{X: 42}, somepkg.SomeStruct{X: 42}, somepkg.SomeStruct{X: 42}, somepkg.SomeStruct{X: 42}, somepkg.SomeStruct{X: 42}, somepkg.SomeStruct{X: 42}, somepkg.SomeStruct{X: 42}, somepkg.SomeStruct{X: 42}}}}",
						`ImportForm(alfa)`,
					},
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			iGen := test.ValuesMockSetup()
			opts := &Options{
				OrganismAmount:   1,
				TestCasesPerFunc: 1,
				MaxRecursion:     10,
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

			for _, testResult := range test.TestResults {
				s.Run(testResult.Func, func() {
					funcTestCases := s.GetTestCase(files, testResult.Func)
					s.Require().Equal(1, len(funcTestCases))
					funcTestCase := funcTestCases[0]
					s.Require().Equal(len(testResult.ResStmts)-1, len(funcTestCase.Stmts))
					for i, stmt := range funcTestCase.Stmts {
						s.Equal(testResult.ResStmts[i], stmt)
					}
					s.Equal(testResult.ResStmts[len(testResult.ResStmts)-1], funcTestCase.FuncStmt)

					s.Require().Equal(len(testResult.ResDecls), len(funcTestCase.Decls))
					for i, decl := range funcTestCase.Decls {
						s.Equal(testResult.ResDecls[i], decl)
					}
				})
			}
		})
	}
}

type TestChanResult struct {
	Func          string
	ResStmts      []string
	ResDecls      []string
	ResChanIdents []string
}

func (s *PrintStmtTestSuite) TestChan() {
	tests := []struct {
		Name            string
		Path            string
		ValuesMockSetup func() *values.GenMock
		TestResults     []TestChanResult
	}{
		{
			Name: "chan test",
			Path: "../../test/data/inputs/example_chan",
			ValuesMockSetup: func() *values.GenMock {
				genMock := &values.GenMock{}
				genMock.On("Int").Return(`42`)
				return genMock
			},
			TestResults: []TestChanResult{
				{
					Func: "FuncChan",
					ResStmts: []string{
						"alfa := make(chan int)",
						"alfa := alfa",
						`FuncChan(alfa)`,
					},
					ResChanIdents: []string{"alfa"},
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			iGen := test.ValuesMockSetup()
			opts := &Options{
				MaxRecursion:     3,
				OrganismAmount:   1,
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

					s.Require().Equal(len(testResult.ResChanIdents), len(funcTestCase.ChanIdents))
					for i, stmt := range funcTestCase.ChanIdents {
						s.Equal(testResult.ResChanIdents[i], stmt)
					}
					s.Require().Equal(len(testResult.ResStmts)-1, len(funcTestCase.Stmts))
					fmt.Println(funcTestCase.ChanIdents)
					for i, stmt := range funcTestCase.Stmts {
						s.Equal(testResult.ResStmts[i], stmt)
					}
					s.Equal(testResult.ResStmts[len(testResult.ResStmts)-1], funcTestCase.FuncStmt)

					s.Require().Equal(len(testResult.ResDecls), len(funcTestCase.Decls))
					for i, decl := range funcTestCase.Decls {
						s.Equal(testResult.ResDecls[i], decl)
					}
				})
			}
		})
	}
}

func TestPrintStmtTestSuite(t *testing.T) {
	suite.Run(t, new(PrintStmtTestSuite))
}

func (s *PrintStmtTestSuite) GetTestCase(files []*File, funcName string) []*testcase.TestCase {
	for _, file := range files {
		if x, ok := file.TestCases[funcName]; ok {
			return x
		}
	}
	s.Fail("test not found")
	return nil
}
