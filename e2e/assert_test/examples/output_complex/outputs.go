package outputs

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type HelloWorld struct {
	X int
	Y []Other
}

type Other struct {
	Z  []bool
	Hi X
}

type X byte

func SomeFunc(x int) ([]int, HelloWorld, chan (int), *float64, X, complex128, string, map[string]int, uuid.UUID, error) {
	id := uuid.Must(uuid.FromString("1662da63-eb21-4787-b0b7-37719bb70df0"))
	y := make(chan int)
	asdf := 34.23423
	lkj := make(map[string]int)
	lkj["foo"] = x

	if x < 10 {
		lkj["joe"] = 34
		return []int{x}, HelloWorld{X: 2}, y, &asdf, 9, 234.333333333333333, "h\"}ello", lkj, id, fmt.Errorf("hi")
	}
	return []int{x, 20}, HelloWorld{X: 43, Y: []Other{{Hi: 3, Z: []bool{false, true}}, {Z: []bool{false, true}}}}, y, nil, 0x01, 234.8, `worl()&
	^#$@d`, lkj, id, nil
}

func SomeFuncPanicFunc(x int) int {
	panic("AAAH")
	return 0
}
