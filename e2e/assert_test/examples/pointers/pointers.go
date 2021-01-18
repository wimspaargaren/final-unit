package pointers

func NoResultFunc() {
}

func PointerFunc() *int {
	x := 4
	return &x
}

func PointerMapFunc() map[int]*string {
	x := make(map[int]*string)
	y := 2
	z := "hi"
	x[y] = &z
	return x
}
