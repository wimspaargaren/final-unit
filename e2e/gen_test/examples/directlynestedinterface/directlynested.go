package directlynested

import "strconv"

type X interface {
	Hello(x int)
}

type Y interface {
	X
	World(x int)
}

func ToTest(x Y) {
	x.Hello(3)
}

func AnotherFuncToTest(x, y int, z string) string {
	if x > 0 && y < 0 {
		return z
	}
	if len(z) == 4 {
		return strconv.Itoa(x)
	}
	return z + "foo"
}
