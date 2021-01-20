package ident

import "go/ast"

// NewMock creates new mock
func NewMock() IGen {
	return &Mock{}
}

// Mock IGen mock
type Mock struct{}

// Create mocks local creation
func (m *Mock) Create(i *ast.Ident) *ast.Ident {
	return &ast.Ident{
		Name: "alfa",
	}
}

// CreateGlobal mocks global creation
func (m *Mock) CreateGlobal(i *ast.Ident) *ast.Ident {
	return &ast.Ident{
		Name: "Alfa",
	}
}

// ResetLocal reset local scope mock
func (m *Mock) ResetLocal() {}
