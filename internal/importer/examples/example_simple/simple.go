package simple

import (
	"fmt"

	foo "github.com/gofrs/uuid"
	"github.com/wimspaargaren/final-unit/internal/importer/examples/example_simple/pkg/somepkg"
)

func HelloWorld(x somepkg.SomeStruct, id foo.UUID) {
	fmt.Printf("hi: %v", x)
}
