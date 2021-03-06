package directlynestedstruct

import (
	"github.com/wimspaargaren/final-unit/test/data/inputs/example_directly_nested_struct/nestedimport"
)

type X int

type Normal struct {
	X
}

type PointerNormal struct {
	*X
}

func NormalNestFunc(x Normal) {
}

func NormalPointerNestFunc(x PointerNormal) {
}

type ImportNested struct {
	nestedimport.Hello
}

func NestedImportFunc(x ImportNested) {
}

type ImportPointerNested struct {
	*nestedimport.Hello
}

func NestedImportPointerFunc(x ImportPointerNested) {
}
