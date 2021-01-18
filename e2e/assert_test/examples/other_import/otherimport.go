package otherimport

import "os"

type Y struct {
	x int
}

type X struct {
	*Y
}

func Hi(x *X) *X {
	return x
}

type B struct{}

type A struct {
	*B
}

func Second(x *A) *A {
	return x
}

// ReaderFromFile creates csv reader from file at given filepath
func ReaderFromFile() (*os.File, error) {
	return &os.File{}, nil
}
