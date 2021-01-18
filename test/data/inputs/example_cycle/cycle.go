package cycle

import "github.com/wimspaargaren/final-unit/test/data/inputs/example_cycle/somepkg"

type A struct {
	X *B
	Y int
}

type B struct {
	X *A
	Y int
}

func FuncCycle(x A) {}

func main() {
	xalfa := B{}
	alfa := A{X: &xalfa, Y: 42}
	FuncCycle(alfa)
}

type X interface {
	X() X
}

func CycleInterface(x X) {
}

func CycleComplicated(x somepkg.A) {
}
