package somepkg

type SomeStruct struct {
	X int
}

type SomeInterface interface {
	Method(x int)
}

type CustomType func(x int)

type UUID [8]byte

type Nested struct {
	X int
}

type NestedStructInImport struct {
	X Nested
}

type X interface {
	Hello(x int)
}

type Y interface {
	X
	World(x int)
}

type CustomTypeInt int

type AStructWithCustomType struct {
	CustomTypeInt CustomTypeInt
}

type Form struct {
	File map[SomeStruct][]SomeStruct
}
