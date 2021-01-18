package somepkg

import "io"

type A struct {
	B *B
}

type B struct {
	X func() (io.ReadCloser, error)
	Y io.ReadCloser
	Z *A
}
