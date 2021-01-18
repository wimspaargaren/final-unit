package exampleimport

type i interface {
	Hello()
}

type Hi struct {
	X func() i
}

type SomeFunc func() i

type Other struct {
	X func() i
}
