package ident

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IdentTestSuite struct {
	suite.Suite
}

func (s *IdentTestSuite) TestSuiteScope() {
	gen := New()
	ident := &ast.Ident{
		Name: "s",
	}
	result1 := gen.Create(ident)
	s.Equal("s2", result1.Name)
}

func (s *IdentTestSuite) TestNormalDupl() {
	gen := New()
	ident := &ast.Ident{
		Name: "x",
	}
	result1 := gen.Create(ident)
	s.Equal("x", result1.Name)
	result2 := gen.Create(ident)
	s.Equal("x2", result2.Name)
	gen.ResetLocal()
	s.Equal("x", gen.Create(ident).Name)
}

func (s *IdentTestSuite) TestGlobalDuplForNormal() {
	gen := NewGenWithGlobal(map[string]int{"x": 2})
	ident := &ast.Ident{
		Name: "x",
	}
	result1 := gen.Create(ident)
	s.Equal("x3", result1.Name)
	result2 := gen.Create(ident)
	s.Equal("x4", result2.Name)
}

func (s *IdentTestSuite) TestGlobalDupl() {
	gen := New()
	ident := &ast.Ident{
		Name: "x",
	}
	result1 := gen.CreateGlobal(ident)
	s.Equal("x", result1.Name)
	result2 := gen.CreateGlobal(ident)
	s.Equal("x2", result2.Name)
}

func TestIdentTestSuite(t *testing.T) {
	suite.Run(t, new(IdentTestSuite))
}
