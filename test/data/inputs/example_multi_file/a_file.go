package multifile

type StructFromOtherFile struct {
	X int
	Y string
}

type InterfaceFromOtherFile interface {
	Method(x int)
}

type CustomType func(x int)
