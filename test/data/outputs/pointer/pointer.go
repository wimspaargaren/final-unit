package pointer

func PointerFunc() *int {
	x := 3
	y := &x
	return y
}
