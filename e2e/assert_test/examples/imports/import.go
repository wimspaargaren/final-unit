package withimport

import "github.com/wimspaargaren/final-unit/e2e/assert_test/examples/imports/someimport"

func StructImport() someimport.SomeStruct {
	return someimport.SomeStruct{
		X: 32,
		Y: &[]int{3, 5},
	}
}

func MapTypeImport() someimport.X {
	mapRes := make(map[int]*someimport.SomeStruct)

	mapRes[3] = &someimport.SomeStruct{
		X: 32,
	}
	mapRes[4] = &someimport.SomeStruct{
		X: 42,
		Y: &[]int{},
	}
	return mapRes
}

func NestedCustomTypeImport() someimport.Y {
	return someimport.Y(someimport.Z([]int{3, 3, 4}))
}

func DirectlyNestedStructImport() someimport.Nested {
	x := someimport.Nested{
		Foo: "foo",
	}

	x.X = 32

	x.Y = &[]int{234, 3, -34}

	return x
}

func UnSupFields() someimport.UnsupFields {
	return someimport.UnSupGetFields()
}
