package customtype

type CustomIntArray []int

func CustomTypeSimple(x CustomIntArray) {
}

type Test struct {
	X int
}
type StructType Test

func CustomTypeStructType(x StructType) {
}

type UUID [8]byte

func CustomByteArray(x UUID) {
}
