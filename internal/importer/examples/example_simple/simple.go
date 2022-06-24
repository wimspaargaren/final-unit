package simple

import (
	"fmt"

	"github.com/asherascout/final-unit/internal/importer/examples/example_simple/pkg/somepkg"
	foo "github.com/gofrs/uuid"
)

func HelloWorld(x somepkg.SomeStruct, id foo.UUID) {
	fmt.Printf("hi: %v", x)
}
