package someimport

type SomeStruct struct {
	X int
	Y *[]int
}

type X map[int]*SomeStruct

type Y Z

type Z []int

type Nested struct {
	SomeStruct
	Foo string
}

type UnsupFields struct {
	privateField string
	X            int
	Y            interface{}
}

func UnSupGetFields() UnsupFields {
	return UnsupFields{
		privateField: "dont show me",
		X:            42,
		Y:            "dont show me",
	}
}
