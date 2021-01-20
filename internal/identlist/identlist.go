// Package identlist keeps track of a list of identifiers while traversing to the AST
package identlist

import (
	"go/ast"
)

// Item item in the list
type Item struct {
	Prev *Item
	I    *ast.Ident
}

// IdentList list of identifiers
type IdentList struct {
	root    Item
	current Item
}

// New creates new idnetlist
func New(i *ast.Ident) IdentList {
	return IdentList{
		root:    Item{I: i},
		current: Item{I: i},
	}
}

// Add adds new identifier to the list
func (l *IdentList) Add(i *ast.Ident) {
	x := l.current
	l.current = Item{
		Prev: &x,
		I:    i,
	}
}

// Previous retrieve last identifier
func (l *IdentList) Previous() *ast.Ident {
	prev := l.current.Prev
	if prev == nil {
		return l.current.I
	}
	return prev.I
}

// Current retrieve current
func (l *IdentList) Current() *ast.Ident {
	return l.current.I
}
