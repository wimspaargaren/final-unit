// Package ident provides a identifier creator which enforces unique variable names within the same scope
package ident

import (
	"fmt"
	"go/ast"
	"unicode"

	"github.com/asherascout/final-unit/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// IGen variable generator interface
type IGen interface {
	Create(i *ast.Ident) *ast.Ident
	CreateGlobal(i *ast.Ident) *ast.Ident
	ResetLocal()
}

// Gen IGen implementation, responsible for generating unique variable names
type Gen struct {
	LocalScopeMem   map[string]int
	GlobaleScopeMem map[string]int
}

// New creates a new variable generator
func New() IGen {
	return NewGenWithGlobal(make(map[string]int))
}

// NewGenWithGlobal creates a new variable generator from a global map
func NewGenWithGlobal(global map[string]int) IGen {
	return &Gen{
		// Suite test functions always have a local scope with s
		LocalScopeMem:   newLocalScope(),
		GlobaleScopeMem: global,
	}
}

// Create creates a new local scope identifier
func (g *Gen) Create(i *ast.Ident) *ast.Ident {
	i = identToLower(i)
	_, ok := g.GlobaleScopeMem[i.Name]
	if ok {
		return g.CreateGlobal(i)
	}

	defer g.increaseLocalMem(i.Name)
	x, ok := g.LocalScopeMem[i.Name]
	if !ok {
		return i
	}

	return &ast.Ident{
		Name: fmt.Sprintf("%s%d", i.Name, x+1),
	}
}

// CreateGlobal creates new global scope identifier
func (g *Gen) CreateGlobal(i *ast.Ident) *ast.Ident {
	i = globalIdentPrefix(i)
	defer g.increaseGlobalMem(i.Name)
	x, ok := g.GlobaleScopeMem[i.Name]
	if !ok {
		return i
	}

	return &ast.Ident{
		Name: fmt.Sprintf("%s%d", i.Name, x+1),
	}
}

// ResetLocal resets the local scope memory
func (g *Gen) ResetLocal() {
	g.LocalScopeMem = newLocalScope()
}

func (g *Gen) increaseLocalMem(name string) {
	g.LocalScopeMem[name]++
}

func (g *Gen) increaseGlobalMem(name string) {
	g.GlobaleScopeMem[name]++
}

func newLocalScope() map[string]int {
	return map[string]int{"s": 1}
}

func globalIdentPrefix(i *ast.Ident) *ast.Ident {
	if unicode.IsLower(rune(i.Name[0])) {
		return &ast.Ident{
			Name: "test" + cases.Title(language.English).String(i.Name),
		}
	}
	return &ast.Ident{
		Name: "Test" + cases.Title(language.English).String(i.Name),
	}
}

func identToLower(i *ast.Ident) *ast.Ident {
	return &ast.Ident{
		Name: utils.LowerCaseFirstLetter(i.Name),
	}
}
