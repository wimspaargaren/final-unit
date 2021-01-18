package exampleinterface

func InterfaceFuncEmpty(x interface{}) {
}

type Simple interface {
	Hello(x int)
}

func InterfaceFuncSimple(x Simple) {
}

type SimpleWithReturn interface {
	Hello(x int) int
}

func InterfaceFuncSimpleWithReturn(x SimpleWithReturn) {
}

type Some struct {
	x string
}

type Complex interface {
	Hello(x *int) [2]byte
	World(x Simple) (*Some, error)
}

func InterfaceComplex(x Complex) {
}

type Nested interface {
	Simple

	World(x int) int
}

func InterfaceNested(x Nested) {
}
