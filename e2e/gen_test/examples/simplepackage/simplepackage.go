// Package simplepackage is used for e2e testing a simple package
// nolint: gomnd, gocritic
package simplepackage

// NestStuff example directly nested struct
type NestStuff struct {
	X func(x int) int
}

// SomeStruct some struct example
type SomeStruct struct {
	NestStuff
	X SomeInterface

	SomeString string
}

// AnotherStruct another struct example
type AnotherStruct struct {
	A uint
	B uint8
	C uint16
	D uint32
	E uint64
	F uintptr
	G int
	H int8
	I int16
	J int32
	K int64
	L byte
	M rune
	N float32
	O float64
	P string
	Q bool
	R map[string]int
	S []string
	T func(x int) int
	U complex64
	V complex128
}

// SomeOtherInterface some other interface example
type SomeOtherInterface interface {
	World(x string) error
}

// SomeInterface some interface example
type SomeInterface interface {
	SomeOtherInterface

	Hello(x []int) int
}

// Hello hello function needs testing
func (s *SomeStruct) Hello(arr []int, x AnotherStruct, y func(s AnotherStruct) error) (string, []int) {
	for i, x := range arr {
		if x > 3 && x < 10 {
			arr[i] = s.X.Hello(arr)
		} else if x < 10 {
			arr[i] = arr[i] + 10
		} else {
			arr[i] = 0
		}
	}
	err := y(x)
	if err != nil {
		return "empty", []int{}
	}
	return s.SomeString, arr
}
