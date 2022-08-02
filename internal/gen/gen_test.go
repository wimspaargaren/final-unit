package gen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wimspaargaren/final-unit/internal/testcase"
	"github.com/wimspaargaren/final-unit/pkg/seed"
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
		Name        string
		Path        string
		TestResults []TestResult
	}{
		{
			Name: "ellipsis test",
			Path: "../../test/data/inputs/example_ellipsis",
			TestResults: []TestResult{
				{
					Func: "EllipsisStringFunc",
					ResStmts: []string{
						"x := \"Bart Beatty\"",
						"EllipsisStringFunc(x)",
					},
				},
				{
					Func: "EllipsisStructFunc",
					ResStmts: []string{
						`x := SomeStruct{X: -73}`,
						"EllipsisStructFunc(x)",
					},
				},
			},
		},
		{
			Name: "base types test",
			Path: "../../test/data/inputs/example_base_types",
			TestResults: []TestResult{
				{
					Func:     "UIntFunc",
					ResStmts: []string{`x := uint(76)`, "UIntFunc(x)"},
				},
				{
					Func:     "UInt8Func",
					ResStmts: []string{`x := uint8(35)`, "UInt8Func(x)"},
				},
				{
					Func:     "UInt16Func",
					ResStmts: []string{`x := uint16(64)`, "UInt16Func(x)"},
				},
				{
					Func:     "UInt32Func",
					ResStmts: []string{`x := uint32(8)`, "UInt32Func(x)"},
				},
				{
					Func:     "UInt64Func",
					ResStmts: []string{`x := uint64(35)`, "UInt64Func(x)"},
				},
				{
					Func:     "Int8Func",
					ResStmts: []string{`x := int8(-47)`, "Int8Func(x)"},
				},
				{
					Func:     "Int16Func",
					ResStmts: []string{`x := int16(28)`, "Int16Func(x)"},
				},
				{
					Func:     "Int32Func",
					ResStmts: []string{`x := int32(31)`, "Int32Func(x)"},
				},
				{
					Func:     "Int64Func",
					ResStmts: []string{`x := int64(41)`, "Int64Func(x)"},
				},
				{
					Func:     "RuneFunc",
					ResStmts: []string{`x := rune(-61)`, "RuneFunc(x)"},
				},
				{
					Func:     "Complex64Func",
					ResStmts: []string{`x := complex64(42)`, "Complex64Func(x)"},
				},
				{
					Func:     "Complex128Func",
					ResStmts: []string{`x := complex128(90)`, "Complex128Func(x)"},
				},
			},
		},
		{
			Name: "int test",
			Path: "../../test/data/inputs/example_int",
			TestResults: []TestResult{
				{
					Func:     "IntFunc",
					ResStmts: []string{"x := -80", "IntFunc(x)"},
				},
				{
					Func:     "IntPointerFunc",
					ResStmts: []string{"pointerX := -45", "x := &pointerX", "IntPointerFunc(x)"},
				},
				{
					Func:     "IntArrayFunc",
					ResStmts: []string{"x := []int{-73, -92, 70, -41, 89, -47, 28}", "IntArrayFunc(x)"},
				},
				{
					Func:     "IntPointerArrayFunc",
					ResStmts: []string{"pointerX := 31", "pointerX2 := 41", "pointerX3 := -61", "pointerX4 := 90", "pointerX5 := -37", "pointerX6 := -77", "pointerX7 := 70", "pointerX8 := -95", "pointerX9 := -12", "x := []*int{&pointerX, &pointerX2, &pointerX3, &pointerX4, &pointerX5, &pointerX6, &pointerX7, &pointerX8, &pointerX9}", "IntPointerArrayFunc(x)"},
				},
				{
					Func:     "IntArrayPointerFunc",
					ResStmts: []string{"pointerX := []int{25, -85, -68, 91, -49, 60, 9, 2}", "x := &pointerX", "IntArrayPointerFunc(x)"},
				},
				{
					Func:     "IntPointerArrayPointerFunc",
					ResStmts: []string{"pointerX := []*int{}", "x := &pointerX", "IntPointerArrayPointerFunc(x)"},
				},
				{
					Func:     "IntArrayLenFunc",
					ResStmts: []string{"x := [2]int{38, 90}", "IntArrayLenFunc(x)"},
				},
				{
					Func:     "IntDoubleArrayFunc",
					ResStmts: []string{"x := [][]int{}", "IntDoubleArrayFunc(x)"},
				},
				{
					Func:     "IntDoubleArrayLenFunc",
					ResStmts: []string{"x := [][2]int{[2]int{81, 13}}", "IntDoubleArrayLenFunc(x)"},
				},
			},
		},
		{
			Name: "unnamed test",
			Path: "../../test/data/inputs/example_unnamed",
			TestResults: []TestResult{
				{
					Func: "SimpleUnnamed",
					ResStmts: []string{
						"x := Simple{X: struct{ Y string }{Y: \"Bart Beatty\"}, Z: Other{Y: \"Cordia Jacobi\"}}",
						"SimpleUnnamed(x)",
					},
				},
				{
					Func: "SimpleInterfaceUnnamed",
					ResStmts: []string{
						"x := SimpleInterface{X: &TestSimpleinterface{}}",
						"SimpleInterfaceUnnamed(x)",
					},
					ResDecls: []string{
						"type TestSimpleinterface struct {\n}",
						"func (s *TestSimpleinterface) Hi(x int) {\n\treturn\n}",
					},
				},
			},
		},
		{
			Name: "nset struct test",
			Path: "../../test/data/inputs/example_directly_nested_struct",
			TestResults: []TestResult{
				{
					Func:     "NormalNestFunc",
					ResStmts: []string{"x := Normal{X: X(-80)}", "NormalNestFunc(x)"},
				},
				{
					Func:     "NormalPointerNestFunc",
					ResStmts: []string{"pointerX := X(-45)", "x := PointerNormal{X: &pointerX}", "NormalPointerNestFunc(x)"},
				},
				{
					Func:     "NestedImportFunc",
					ResStmts: []string{"x := ImportNested{Hello: nestedimport.Hello{X: -73}}", "NestedImportFunc(x)"},
				},
				{
					Func:     "NestedImportPointerFunc",
					ResStmts: []string{"pointerX := nestedimport.Hello{X: -92}", "x := ImportPointerNested{Hello: &pointerX}", "NestedImportPointerFunc(x)"},
				},
			},
		},
		{
			Name: "bool test",
			Path: "../../test/data/inputs/example_bool",
			TestResults: []TestResult{
				{
					Func:     "BoolFunc",
					ResStmts: []string{"x := false", "BoolFunc(x)"},
				},
				{
					Func:     "BoolPointerFunc",
					ResStmts: []string{"pointerX := true", "x := &pointerX", "BoolPointerFunc(x)"},
				},
			},
		},
		{
			Name: "string test",
			Path: "../../test/data/inputs/example_string",
			TestResults: []TestResult{
				{
					Func: "StringFunc",
					ResStmts: []string{
						"x := \"Bart Beatty\"",
						"StringFunc(x)",
					},
				},
			},
		},
		{
			Name: "multi param test",
			Path: "../../test/data/inputs/example_multi_param_same_name",
			TestResults: []TestResult{
				{
					Func:     "MultiParamFunc",
					ResStmts: []string{"x := \"Bart Beatty\"", "y := \"Cordia Jacobi\"", "MultiParamFunc(x, y)"},
				},
				{
					Func:     "MultiParamPointerFunc",
					ResStmts: []string{"pointerX := \"Nickolas Emard\"", "x := &pointerX", "pointerY := \"Hollis Dickens\"", "y := &pointerY", "MultiParamPointerFunc(x, y)"},
				},
				{
					Func:     "MultiParamDifTypeFunc",
					ResStmts: []string{"x := 28", "y := \"Marc Murphy\"", "MultiParamDifTypeFunc(x, y)"},
				},
			},
		},
		{
			Name: "float test",
			Path: "../../test/data/inputs/example_float",
			TestResults: []TestResult{
				{
					Func:     "Float32Func",
					ResStmts: []string{`x := float32(20.932058)`, "Float32Func(x)"},
				},
				{
					Func:     "Float64Func",
					ResStmts: []string{`x := 88.101818`, "Float64Func(x)"},
				},
			},
		},
		{
			Name: "byte test",
			Path: "../../test/data/inputs/example_byte",
			TestResults: []TestResult{
				{
					Func:     "ByteFunc",
					ResStmts: []string{`x := byte(76)`, "ByteFunc(x)"},
				},
				{
					Func:     "BytePointerFunc",
					ResStmts: []string{`pointerX := byte(35)`, `x := &pointerX`, "BytePointerFunc(x)"},
				},
				{
					Func:     "ByteArrayFunc",
					ResStmts: []string{`x := []byte{byte(64), byte(8), byte(35), byte(92), byte(37), byte(16), byte(0)}`, "ByteArrayFunc(x)"},
				},
				{
					Func:     "ByteArrayPointerFunc",
					ResStmts: []string{"pointerX := byte(46)", "pointerX2 := byte(64)", "pointerX3 := byte(36)", "pointerX4 := byte(60)", "pointerX5 := byte(27)", "pointerX6 := byte(42)", "pointerX7 := byte(41)", "pointerX8 := byte(20)", "pointerX9 := byte(21)", "x := []*byte{&pointerX, &pointerX2, &pointerX3, &pointerX4, &pointerX5, &pointerX6, &pointerX7, &pointerX8, &pointerX9}", "ByteArrayPointerFunc(x)"},
				},
			},
		},
		{
			Name: "struct test",
			Path: "../../test/data/inputs/example_struct",
			TestResults: []TestResult{
				{
					Func:     "StructFuncSimple",
					ResStmts: []string{`x := Simple{Var: -80}`, "StructFuncSimple(x)"},
				},
				{
					Func:     "StructFuncComposite",
					ResStmts: []string{"pointerX := \"Cordia Jacobi\"", "x := Composite{VarInt: -45, VarString: &pointerX}", "StructFuncComposite(x)"},
				},
				{
					Func:     "StructFuncCompositeOther",
					ResStmts: []string{`x := CompositeOther{x: 70, y: -41}`, "StructFuncCompositeOther(x)"},
				},
				{
					Func:     "StructFuncNested",
					ResStmts: []string{"x := Nested{X: Simple{Var: 89}}", "StructFuncNested(x)"},
				},
				{
					Func:     "StructFuncNestedNested",
					ResStmts: []string{"x := NestedNested{Nested: Nested{X: Simple{Var: -47}}, x: 28}", "StructFuncNestedNested(x)"},
				},
				{
					Func:     "StructFuncNestedInt",
					ResStmts: []string{"x := NestedInt{int: 31, x: \"Aleen Legros\"}", "StructFuncNestedInt(x)"},
				},
				{
					Func:     "StructFuncNestedCustomType",
					ResStmts: []string{"x := NestedCustomType{CustomType: CustomType(func(x int) string {\n\to := \"Adelia Metz\"\n\treturn o\n}), x: \"Sunny Gerlach\"}", "StructFuncNestedCustomType(x)"},
				},
			},
		},
		{
			Name: "custom type test",
			Path: "../../test/data/inputs/example_custom_type",
			TestResults: []TestResult{
				{
					Func:     "CustomTypeSimple",
					ResStmts: []string{`x := CustomIntArray([]int{-80, -45, -73, -92, 70, -41, 89})`, "CustomTypeSimple(x)"},
				},
				{
					Func:     "CustomTypeStructType",
					ResStmts: []string{"x := StructType(Test{X: -47})", "CustomTypeStructType(x)"},
				},
				{
					Func:     "CustomByteArray",
					ResStmts: []string{"x := UUID([8]byte{byte(0), byte(46), byte(64), byte(36), byte(60), byte(27), byte(42), byte(41)})", "CustomByteArray(x)"},
				},
			},
		},
		{
			Name: "array test",
			Path: "../../test/data/inputs/example_array",
			TestResults: []TestResult{
				{
					Func:     "ArrayLenZero",
					ResStmts: []string{`x := [0]int{}`, "ArrayLenZero(x)"},
				},
				{
					Func:     "ArrayLenFive",
					ResStmts: []string{`x := [5]int{-80, -45, -73, -92, 70}`, "ArrayLenFive(x)"},
				},
				{
					Func:     "ArrayNoLen",
					ResStmts: []string{`x := []int{-41, 89, -47, 28, 31, 41, -61}`, "ArrayNoLen(x)"},
				},
				{
					Func:     "ArrayStructNoLen",
					ResStmts: []string{`x := []Simple{Simple{X: 90}}`, "ArrayStructNoLen(x)"},
				},
				{
					Func: "ArrayPointerStructNoLen",
					ResStmts: []string{
						"pointerX2 := -37",
						"pointerX3 := -77",
						"pointerX4 := 70",
						"pointerX := []Pointer{Pointer{x: &pointerX2}, Pointer{x: &pointerX3}, Pointer{x: &pointerX4}, Pointer{}, Pointer{}}",
						`x := &pointerX`,
						"ArrayPointerStructNoLen(x)",
					},
				},
			},
		},
		{
			Name: "func on struct",
			Path: "../../test/data/inputs/example_func_on_struct",
			TestResults: []TestResult{
				{
					Func: "SimpleSimpleFuncOnStruct",
					ResStmts: []string{
						"s2 := Simple{X: -80}",
						`x := -45`,
						`s2.SimpleFuncOnStruct(x)`,
					},
				},
				{
					Func: "CustomTypeFuncOnCustomType",
					ResStmts: []string{
						"c := CustomType(-73)",
						`x := -92`,
						`c.FuncOnCustomType(x)`,
					},
				},
			},
		},
		{
			Name: "map test",
			Path: "../../test/data/inputs/example_map",
			TestResults: []TestResult{
				{
					Func: "MapFunc",
					ResStmts: []string{
						"x := map[int]string{-80: \"Gerson Beahan\"}",
						`MapFunc(x)`,
					},
				},
				{
					Func: "MapFuncNested",
					ResStmts: []string{
						"x := map[int]map[int]string{-92: map[int]string{70: \"Lawson Kreiger\", -47: \"Stacy Dietrich\", 41: \"Eunice Kunde\", -37: \"Sunny Gerlach\", -95: \"Anika Durgan\", -85: \"Delaney Howell\", -49: \"Christian Bartoletti\"}, 2: map[int]string{38: \"Gerda Rosenbaum\", 13: \"Elias Roob\", -82: \"Alexandra Halvorson\", 12: \"Guido Witting\", 5: \"Sim Erdman\", -87: \"Vincenza Jacobi\", 56: \"Skye Lemke\", -33: \"Dorian Hartmann\", -57: \"Brody Walker\"}, 96: map[int]string{-16: \"Oleta Haley\"}, 78: map[int]string{-52: \"Eugenia Skiles\", 62: \"Mariah Bergstrom\", 95: \"Brittany Hermann\", 52: \"Makayla Kuhn\", -78: \"Alexandria Kihn\", -17: \"Ericka Schmitt\", -31: \"Marlene Wisozk\", 42: \"Amos Funk\"}, 58: map[int]string{-83: \"Olaf Flatley\", -33: \"Jewell Cartwright\", -40: \"Larry Kemmer\", -64: \"Minnie Adams\"}, 19: map[int]string{}, 18: map[int]string{-79: \"Osbaldo Ruecker\", -7: \"Nicholaus Gerhold\", -58: \"Filiberto Pollich\", 48: \"Adah McGlynn\", 7: \"Lempi Legros\", -68: \"Giovani Gorczany\"}}",
						`MapFuncNested(x)`,
					},
				},
			},
		},
		{
			Name: "interface import test",
			Path: "../../test/data/inputs/example_interface_import",
			TestResults: []TestResult{
				{
					Func: "SomeFuncImportingInterface",
					ResStmts: []string{
						"x := &testX{}",
						`SomeFuncImportingInterface(x)`,
					},
					ResDecls: []string{
						"type testX struct {\n}",
						"func (s *testX) SomeFunc(someimport.Y) someimport.Y {\n\to := someimport.Y{X: -80}\n\treturn o\n}",
					},
				},
			},
		},
		{
			Name: "interface private return test",
			Path: "../../test/data/inputs/example_priv_ret_interface",
			TestResults: []TestResult{
				{
					Func:     "ReflectType",
					ResStmts: []string{"x := func() reflect.Type {\n\treturn nil\n}()", "ReflectType(x)"},
					ResDecls: []string{},
				},
				{
					Func:     "AstSelection",
					ResStmts: []string{"x := func() ast.Selection {\n\treturn nil\n}()", "AstSelection(x)"},
					ResDecls: []string{},
				},
			},
		},
		{
			Name: "error test",
			Path: "../../test/data/inputs/example_error",
			TestResults: []TestResult{
				{
					Func:     "ErrorFunc",
					ResStmts: []string{"x := func() error {\n\treturn nil\n}()", "ErrorFunc(x)"},
				},
				{
					Func:     "ErrorPointerFunc",
					ResStmts: []string{"pointerX := func() error {\n\treturn fmt.Errorf(\"very error\")\n}()", `x := &pointerX`, "ErrorPointerFunc(x)"},
				},
			},
		},
		{
			Name: "interface test",
			Path: "../../test/data/inputs/example_interface",
			TestResults: []TestResult{
				{
					Func:     "InterfaceFuncEmpty",
					ResStmts: []string{"x := uint64(35)", "InterfaceFuncEmpty(x)"},
				},
				{
					Func:     "InterfaceFuncSimple",
					ResStmts: []string{"x := &TestSimple{}", "InterfaceFuncSimple(x)"},
					ResDecls: []string{"type TestSimple struct {\n}", "func (s *TestSimple) Hello(x int) {\n\treturn\n}"},
				},
				{
					Func:     "InterfaceFuncSimpleWithReturn",
					ResStmts: []string{"x := &TestSimplewithreturn{}", "InterfaceFuncSimpleWithReturn(x)"},
					ResDecls: []string{
						"type TestSimplewithreturn struct {\n}",
						"func (s *TestSimplewithreturn) Hello(x int) int {\n\to := -73\n\treturn o\n}",
					},
				},
				{
					Func:     "InterfaceComplex",
					ResStmts: []string{"x := &TestComplex{}", "InterfaceComplex(x)"},
					ResDecls: []string{
						"type TestComplex struct {\n}",
						"func (s *TestComplex) Hello(x *int) [2]byte {\n\to := [2]byte{byte(8), byte(35)}\n\treturn o\n}",
						"func (s *TestComplex) World(x Simple) (*Some, error) {\n\tpointerX := Some{x: \"Lawson Kreiger\"}\n\to2 := &pointerX\n\to3 := func() error {\n\t\treturn nil\n\t}()\n\treturn o2, o3\n}",
					},
				},
				{
					Func:     "InterfaceNested",
					ResStmts: []string{`x := &TestNested{}`, "InterfaceNested(x)"},
					ResDecls: []string{
						"type TestNested struct {\n}",
						"func (s *TestNested) Hello(x int) {\n\treturn\n}",
						"func (s *TestNested) World(x int) int {\n\to := 28\n\treturn o\n}",
					},
				},
			},
		},
		{
			Name: "interface 2 test",
			Path: "../../test/data/inputs/example_interface_2",
			TestResults: []TestResult{
				{
					Func:     "NestedDouble",
					ResStmts: []string{`x := &TestX{}`, "NestedDouble(x)"},
					ResDecls: []string{
						"type TestX struct {\n}",
						"func (s *TestX) Hello(int) {\n\treturn\n}",
						"func (s *TestX) World(int) {\n\treturn\n}",
					},
				},
			},
		},
		{
			Name: "cant gen func test",
			Path: "../../test/data/inputs/example_func_return_priv",
			TestResults: []TestResult{
				{
					Func: "Test",
					ResStmts: []string{
						"x := exampleimport.SomeFunc(nil)",
						`Test(x)`,
					},
				},
				{
					Func: "Test2",
					ResStmts: []string{
						"x := exampleimport.Other{}",
						`Test2(x)`,
					},
				},
			},
		},
		{
			Name: "chan test",
			Path: "../../test/data/inputs/example_chan",
			TestResults: []TestResult{
				{
					Func: "FuncChan",
					ResStmts: []string{
						"x2 := make(chan int)",
						"x := x2",
						`FuncChan(x)`,
					},
				},
			},
		},
		{
			Name: "cycle test",
			Path: "../../test/data/inputs/example_cycle",
			TestResults: []TestResult{
				{
					Func: "FuncCycle",
					ResStmts: []string{
						"pointerA3 := A{}",
						"pointerB2 := B{X: &pointerA3, Y: -80}",
						"pointerA2 := A{X: &pointerB2, Y: -45}",
						"pointerB := B{X: &pointerA2, Y: -73}",
						"pointerA := A{X: &pointerB, Y: -92}",
						"pointerX := B{X: &pointerA, Y: 70}",
						"x := A{X: &pointerX, Y: -41}",
						`FuncCycle(x)`,
					},
				},
				{
					Func: "CycleInterface",
					ResStmts: []string{
						"x := &TestX{}",
						`CycleInterface(x)`,
					},
					ResDecls: []string{
						"type TestX struct {\n}",
						"type TestX2 struct {\n}",
						"type TestX3 struct {\n}",
						"func (s *TestX3) X() X {\n\to3 := &TestX{}\n\treturn o3\n}",
						"func (s *TestX2) X() X {\n\to2 := &TestX3{}\n\treturn o2\n}",
						"func (s *TestX) X() X {\n\to := &TestX2{}\n\treturn o\n}",
					},
				},
				{
					Func: "CycleComplicated",
					ResStmts: []string{
						"pointerA3 := somepkg.A{}",
						"pointerB2 := somepkg.B{X: func() (io.ReadCloser, error) {\n\to14 := &TestB{}\n\to15 := func() error {\n\t\treturn fmt.Errorf(\"very error\")\n\t}()\n\treturn o14, o15\n}, Y: &TestB{}, Z: &pointerA3}",
						"pointerA2 := somepkg.A{B: &pointerB2}",
						"pointerB := somepkg.B{X: func() (io.ReadCloser, error) {\n\to9 := &TestB3{}\n\to13 := func() error {\n\t\treturn nil\n\t}()\n\treturn o9, o13\n}, Y: &TestB{}, Z: &pointerA2}",
						"pointerA := somepkg.A{B: &pointerB}",
						"pointerX := somepkg.B{X: func() (io.ReadCloser, error) {\n\to := &TestB{}\n\to5 := func() error {\n\t\treturn fmt.Errorf(\"very error\")\n\t}()\n\treturn o, o5\n}, Y: &TestB2{}, Z: &pointerA}",
						"x := somepkg.A{B: &pointerX}",
						`CycleComplicated(x)`,
					},
					ResDecls: []string{
						"type TestB struct {\n}",
						"func (s *TestB) Read(p []byte) (n int, err error) {\n\to2 := 89\n\to3 := func() error {\n\t\treturn nil\n\t}()\n\treturn o2, o3\n}",
						"func (s *TestB) Close() error {\n\to4 := func() error {\n\t\treturn nil\n\t}()\n\treturn o4\n}",
						"type TestB2 struct {\n}",
						"func (s *TestB2) Read(p []byte) (n int, err error) {\n\to6 := 41\n\to7 := func() error {\n\t\treturn fmt.Errorf(\"very error\")\n\t}()\n\treturn o6, o7\n}",
						"func (s *TestB2) Close() error {\n\to8 := func() error {\n\t\treturn nil\n\t}()\n\treturn o8\n}",
						"type TestB3 struct {\n}",
						"func (s *TestB3) Read(p []byte) (n int, err error) {\n\to10 := -37\n\to11 := func() error {\n\t\treturn fmt.Errorf(\"very error\")\n\t}()\n\treturn o10, o11\n}",
						"func (s *TestB3) Close() error {\n\to12 := func() error {\n\t\treturn fmt.Errorf(\"very error\")\n\t}()\n\treturn o12\n}",
					},
				},
			},
		},
		{
			Name: "func func test",
			Path: "../../test/data/inputs/example_func",
			TestResults: []TestResult{
				{
					Func: "FuncFunc",
					ResStmts: []string{
						"x := func(x int) {\n\treturn\n}",
						`FuncFunc(x)`,
					},
				},
				{
					Func: "FuncTypeFunc",
					ResStmts: []string{
						"x := FuncType(func(x int) {\n\treturn\n})",
						`FuncTypeFunc(x)`,
					},
				},
				{
					Func: "FuncFuncWithReturn",
					ResStmts: []string{
						"x := func(x int) int {\n\to := -80\n\treturn o\n}",
						`FuncFuncWithReturn(x)`,
					},
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			opts := &Options{
				MaxRecursion:     3,
				OrganismAmount:   1,
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
		Name        string
		Path        string
		TestResults []TestResult
	}{
		{
			Name: "func multi file",
			Path: "../../test/data/inputs/example_multi_file",
			TestResults: []TestResult{
				{
					Func: "MultFileStruct",
					ResStmts: []string{
						"x := StructFromOtherFile{X: -80, Y: \"Gerson Beahan\"}",
						`MultFileStruct(x)`,
					},
				},
				{
					Func: "MultFileInterface",
					ResStmts: []string{
						"x := &TestInterfacefromotherfile{}",
						`MultFileInterface(x)`,
					},
					ResDecls: []string{
						"type TestInterfacefromotherfile struct {\n}",
						"func (s *TestInterfacefromotherfile) Method(x int) {\n\treturn\n}",
					},
				},
				{
					Func: "MultFileCustomType",
					ResStmts: []string{
						"x := CustomType(func(x int) {\n\treturn\n})",
						`MultFileCustomType(x)`,
					},
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
		Name        string
		Path        string
		TestResults []TestResult
	}{
		{
			Name: "import funcs",
			Path: "../../test/data/inputs/example_imports",
			TestResults: []TestResult{
				{
					Func: "SimpleImport",
					ResStmts: []string{
						"x := somepkg.SomeStruct{X: -80}",
						`SimpleImport(x)`,
					},
				},
				{
					Func: "NestedImport",
					ResStmts: []string{
						"x := nestedimportpkg.NestedStruct{X: somepkg.SomeStruct{X: -45}}",
						`NestedImport(x)`,
					},
				},
				{
					Func: "ImportInterface",
					ResStmts: []string{
						"x := &testX{}",
						`ImportInterface(x)`,
					},
					ResDecls: []string{"type testX struct {\n}", "func (s *testX) Method(x int) {\n\treturn\n}"},
				},
				{
					Func: "ImportCustomType",
					ResStmts: []string{
						"x := somepkg.CustomType(func(x int) {\n\treturn\n})",
						`ImportCustomType(x)`,
					},
				},
				{
					Func: "ImportCustomTypeUUID",
					ResStmts: []string{
						"x := somepkg.UUID([8]byte{byte(64), byte(8), byte(35), byte(92), byte(37), byte(16), byte(0), byte(46)})",
						`ImportCustomTypeUUID(x)`,
					},
				},
				{
					Func: "NestedUnitInImport",
					ResStmts: []string{
						"x := somepkg.NestedStructInImport{X: somepkg.Nested{X: 41}}",
						`NestedUnitInImport(x)`,
					},
				},
				{
					Func: "ImportTime",
					ResStmts: []string{
						"x := time.Time{}",
						`ImportTime(x)`,
					},
				},
				{
					Func: "DirectlyNestedInterface",
					ResStmts: []string{
						"x := &testX2{}",
						`DirectlyNestedInterface(x)`,
					},
					ResDecls: []string{
						"type testX2 struct {\n}",
						"func (s *testX2) Hello(x int) {\n\treturn\n}",
						"func (s *testX2) World(x int) {\n\treturn\n}",
					},
				},
				{
					Func: "ImportCustomTypeInstruct",
					ResStmts: []string{
						"x := somepkg.AStructWithCustomType{CustomTypeInt: somepkg.CustomTypeInt(-61)}",
						`ImportCustomTypeInstruct(x)`,
					},
				},
				{
					Func: "ImportedMapVal",
					ResStmts: []string{
						"x := map[string]somepkg.SomeStruct{\"Adelia Metz\": somepkg.SomeStruct{X: -77}, \"Ariane Rice\": somepkg.SomeStruct{X: -12}, \"Briana Bauch\": somepkg.SomeStruct{X: -68}, \"Jarod Wolff\": somepkg.SomeStruct{X: 60}, \"Talia Hudson\": somepkg.SomeStruct{X: 38}, \"Gerda Rosenbaum\": somepkg.SomeStruct{X: 13}, \"Elias Roob\": somepkg.SomeStruct{X: -82}}",
						`ImportedMapVal(x)`,
					},
				},
				{
					Func: "ImportCustomTypeInMap",
					ResStmts: []string{
						"x := map[string]somepkg.CustomType{\"Alexandra Halvorson\": somepkg.CustomType(func(x int) {\n\treturn\n}), \"Miller Cormier\": somepkg.CustomType(func(x int) {\n\treturn\n}), \"Matilde Doyle\": somepkg.CustomType(func(x int) {\n\treturn\n}), \"Sim Erdman\": somepkg.CustomType(func(x int) {\n\treturn\n}), \"Marquise Erdman\": somepkg.CustomType(func(x int) {\n\treturn\n}), \"Helmer Crooks\": somepkg.CustomType(func(x int) {\n\treturn\n}), \"Skye Lemke\": somepkg.CustomType(func(x int) {\n\treturn\n})}",
						`ImportCustomTypeInMap(x)`,
					},
				},
				{
					Func: "ImportForm",
					ResStmts: []string{
						"x := somepkg.Form{File: map[somepkg.SomeStruct][]somepkg.SomeStruct{somepkg.SomeStruct{X: -33}: []somepkg.SomeStruct{somepkg.SomeStruct{X: -57}, somepkg.SomeStruct{X: -22}, somepkg.SomeStruct{X: -57}, somepkg.SomeStruct{X: 89}, somepkg.SomeStruct{X: -84}, somepkg.SomeStruct{X: 96}, somepkg.SomeStruct{X: -16}, somepkg.SomeStruct{X: 56}}, somepkg.SomeStruct{X: -68}: []somepkg.SomeStruct{}}}",
						`ImportForm(x)`,
					},
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			opts := &Options{
				OrganismAmount:   1,
				TestCasesPerFunc: 1,
				MaxRecursion:     10,
			}
			seed.SetRandomSeed(1)
			generator, err := New(test.Path, opts)
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
		Name        string
		Path        string
		TestResults []TestChanResult
	}{
		{
			Name: "chan test",
			Path: "../../test/data/inputs/example_chan",
			TestResults: []TestChanResult{
				{
					Func: "FuncChan",
					ResStmts: []string{
						"x2 := make(chan int)",
						"x := x2",
						`FuncChan(x)`,
					},
					ResChanIdents: []string{"x2"},
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			opts := &Options{
				MaxRecursion:     3,
				OrganismAmount:   1,
				TestCasesPerFunc: 1,
			}
			seed.SetRandomSeed(0)
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

					s.Require().Equal(len(testResult.ResChanIdents), len(funcTestCase.ChanIdents))
					for i, stmt := range funcTestCase.ChanIdents {
						s.Equal(testResult.ResChanIdents[i], stmt)
					}
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
