package testcase

import (
	"go/ast"

	log "github.com/sirupsen/logrus"
)

// IsBasicLit reports if an idenetifier is a basic literal
func (g *TestCase) IsBasicLit(identifier string) bool {
	switch identifier {
	case "int",
		"bool",
		"string",
		"float32",
		"float64",
		"byte",
		"rune",
		"uintptr",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"int8",
		"int16",
		"int32",
		"int64",
		"complex64",
		"complex128":
		return true
	default:
		return false
	}
}

// IsError checks if identifier is reserver error keyword
func (g *TestCase) IsError(identifier string) bool {
	return identifier == "error"
}

// IsBasicExpr checks if expr is basic expr
func (g *TestCase) IsBasicExpr(x ast.Expr) (string, bool) {
	if t, ok := x.(*ast.Ident); ok {
		if t.Obj == nil {
			if g.IsBasicLit(t.Name) {
				return t.Name, true
			}
		}
	}
	return "", false
}

// GetUnnamedStructIdent retrieves an identifier for an unnamed struct field
func (g *TestCase) GetUnnamedStructIdent(fieldType ast.Expr, input *RecursionInput) *ast.Ident {
	switch t := fieldType.(type) {
	case *ast.Ident:
		return t
	case *ast.StarExpr:
		return g.GetUnnamedStructIdent(t.X, input)
	case *ast.SelectorExpr:
		return t.Sel
	default:
		log.Warningf("unable to get unnamed struct field")
		return &ast.Ident{}
	}
}
