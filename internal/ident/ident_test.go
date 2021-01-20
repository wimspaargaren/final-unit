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

func (s *IdentTestSuite) TestDuplWithUpperCaseStart() {
	gen := New()
	ident := &ast.Ident{
		Name: "XX",
	}
	result1 := gen.Create(ident)
	s.Equal("xX", result1.Name)
	result2 := gen.Create(ident)
	s.Equal("xX2", result2.Name)
	gen.ResetLocal()
	s.Equal("xX", gen.Create(ident).Name)
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
	s.Equal("testX", result1.Name)
	result2 := gen.Create(ident)
	s.Equal("testX2", result2.Name)
}

func (s *IdentTestSuite) TestGlobalDupl() {
	gen := New()
	ident := &ast.Ident{
		Name: "x",
	}
	result1 := gen.CreateGlobal(ident)
	s.Equal("testX", result1.Name)
	result2 := gen.CreateGlobal(ident)
	s.Equal("testX2", result2.Name)
}

func TestIdentTestSuite(t *testing.T) {
	suite.Run(t, new(IdentTestSuite))
}
