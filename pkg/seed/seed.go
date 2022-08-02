// Package seed provides random seed for evolving test cases
package seed

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

// SetRandomSeed set random seed for global usage
func SetRandomSeed(seed int64) {
	rand.Seed(seed)
	gofakeit.SetGlobalFaker(gofakeit.New(seed))
}
