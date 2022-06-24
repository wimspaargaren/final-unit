package nestedimportpkg

import "github.com/asherascout/final-unit/test/data/inputs/example_imports/pkg/somepkg"

type NestedStruct struct {
	X somepkg.SomeStruct
}
