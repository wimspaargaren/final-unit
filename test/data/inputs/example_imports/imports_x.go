package importsexample

import (
	"time"

	"github.com/wimspaargaren/final-unit/test/data/inputs/example_imports/pkg/nestedimportpkg"
	"github.com/wimspaargaren/final-unit/test/data/inputs/example_imports/pkg/somepkg"
)

func SimpleImport(x somepkg.SomeStruct) {}

func NestedImport(x nestedimportpkg.NestedStruct) {}

func ImportInterface(x somepkg.SomeInterface) {}

func ImportCustomType(x somepkg.CustomType) {}

func ImportCustomTypeUUID(x somepkg.UUID) {}

func NestedUnitInImport(x somepkg.NestedStructInImport) {}

func ImportTime(x time.Time) {}

func DirectlyNestedInterface(x somepkg.Y) {}

func ImportCustomTypeInstruct(x somepkg.AStructWithCustomType) {}

func ImportedMapVal(x map[string]somepkg.SomeStruct) {}

func ImportCustomTypeInMap(x map[string]somepkg.CustomType) {}

func ImportForm(x somepkg.Form) {}

func main() {
	alfa := somepkg.AStructWithCustomType{
		CustomTypeInt: somepkg.CustomTypeInt(42),
	}
	ImportCustomTypeInstruct(alfa)
}
