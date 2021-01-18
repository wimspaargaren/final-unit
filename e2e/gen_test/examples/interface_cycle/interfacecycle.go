package interfacecycle

type X interface {
	X() Y
}

type Y interface {
	Y() X
}

func CycleInterface(x X) {
}
