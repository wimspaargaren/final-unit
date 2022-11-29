// Package variables provides a generator for creating random variable names
package variables

import (
	"math/rand"
)

// IGen variable name generator interface
type IGen interface {
	Generate() string
}

const (
	varNameLength = 5
)

// Generator default generator implementation
type Generator struct{}

// NewGenerator creates a new generator
func NewGenerator() IGen {
	return &Generator{}
}

// Generate generates a random variable name of length 10
func (g *Generator) Generate() string {
	return randSeq(varNameLength)
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
