package interfaceimport

type y struct{}

type x struct{}

type Weird interface {
	X(x) y
	x(x) y
}

type OnlyMethodNotExported interface {
	x(int)
}

type OnlyImportNotExported interface {
	X(x)
}

type OnlyOutputNotExported interface {
	X() x
}
