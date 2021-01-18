package otherimport

import "os"

func ReaderFromFile(f *os.File) error {
	return nil
}

type Y struct{}

type X struct {
	*Y
}

func Hi(x *X) *X {
	return x
}
