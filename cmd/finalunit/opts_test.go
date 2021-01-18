// Assert template
package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OptsSuite struct {
	suite.Suite
}

func TestOptsSuite(t *testing.T) {
	suite.Run(t, new(OptsSuite))
}
