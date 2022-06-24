package packagewithimports

import (
	"net/http"
	"time"

	"github.com/wimspaargaren/final-unit/e2e/gen_test/examples/packagewithimports/pkg/some"
)

func SomeFunc(s *some.SomeStruct, t time.Time, r http.Request) bool {
	if s.X > 50 {
		return true
	}
	return false
}
