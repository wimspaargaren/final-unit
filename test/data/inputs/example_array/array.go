package array

type Simple struct {
	X int
}

type Pointer struct {
	x *int
}

func ArrayLenZero(x [0]int)                {}
func ArrayLenFive(x [5]int)                {}
func ArrayNoLen(x []int)                   {}
func ArrayStructNoLen(x []Simple)          {}
func ArrayPointerStructNoLen(x *[]Pointer) {}
