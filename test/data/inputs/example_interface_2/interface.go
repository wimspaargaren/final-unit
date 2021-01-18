package exampleinterface

type X interface {
	A
	B
}

type A interface {
	Hello(int)
}

type B interface {
	World(int)
}

func NestedDouble(x X) {}
