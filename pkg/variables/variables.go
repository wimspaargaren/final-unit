// Package variables provides a generator for creating random variable names
// nolint: gosec
package variables

import (
	"math/rand"
	"time"
)

// IGen variable name generator interface
type IGen interface {
	Generate() string
}

// not linted for speeding up generating
// nolint: gochecknoglobals
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const (
	varNameLength = 5
)

// Generator default generator implementation
type Generator struct{}

// NewGenerator creates a new generator
func NewGenerator() IGen {
	rand.Seed(time.Now().UnixNano())
	return &Generator{}
}

// Generate generates a random variable name of length 10
func (g *Generator) Generate() string {
	return randSeq(varNameLength)
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
