package identlist

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IdentTestSuite struct {
	suite.Suite
}

func (s *IdentTestSuite) TestTwoItems() {
	i := &ast.Ident{Name: "first"}
	list := New(i)
	list.Add(&ast.Ident{Name: "new"})
	s.Equal("first", list.Previous().Name)
}

func (s *IdentTestSuite) TestOneItem() {
	i := &ast.Ident{Name: "first"}
	list := New(i)
	s.Equal("first", list.Previous().Name)
}

func TestIdentTestSuite(t *testing.T) {
	suite.Run(t, new(IdentTestSuite))
}
