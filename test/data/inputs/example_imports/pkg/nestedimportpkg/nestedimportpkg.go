package nestedimportpkg

import "github.com/wimspaargaren/final-unit/test/data/inputs/example_imports/pkg/somepkg"

type NestedStruct struct {
	X somepkg.SomeStruct
}
