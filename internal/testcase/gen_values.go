package testcase

import (
	"go/ast"
	"go/token"

	log "github.com/sirupsen/logrus"
)

// InterfaceGenDecl creates interface gen decl
func (g *TestCase) InterfaceGenDecl(input *RecursionInput, interfaceImplIdent *ast.Ident) *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: interfaceImplIdent,
				Type: &ast.StructType{
					// If nested interfaces are declared they should be added as fields
					Fields: &ast.FieldList{
						List: []*ast.Field{},
					},
				},
			},
		},
	}
}

// ErrExprToValExpr converts error expression to val expression
func (g *TestCase) ErrExprToValExpr() *TypeExprToValExprRes {
	var errReturn ast.Expr
	// Init default return value to nil
	errReturn = &ast.Ident{
		Name: "nil",
	}
	// if val TestCase provides true return err
	if g.Opts.ValTestCase.Error() {
		errReturn = &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "fmt"},
				Sel: &ast.Ident{Name: "Errorf"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: `"very error"`,
				},
			},
		}
	}
	// Create function returning either nil or an error
	e := &ast.CallExpr{
		Fun: &ast.FuncLit{
			Type: &ast.FuncType{
				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Type: &ast.Ident{
								Name: "error",
							},
						},
					},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{
							errReturn,
						},
					},
				},
			},
		},
	}
	return &TypeExprToValExprRes{
		Expr:         e,
		Statements:   []ast.Stmt{},
		Declarations: []ast.Decl{},
	}
}

// BasicExprToValExpr converts a basic identifier to an expression
func (g *TestCase) BasicExprToValExpr(identifier string) ast.Expr { //nolint: gocyclo
	switch identifier {
	case "int":
		return &ast.BasicLit{
			Kind:  token.INT,
			Value: g.Opts.ValTestCase.Int(),
		}
	case "bool":
		return &ast.Ident{
			Name: g.Opts.ValTestCase.Bool(),
		}
	case "string":
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: g.Opts.ValTestCase.String(),
		}
	case "float32":
		return &ast.CallExpr{
			Fun: &ast.Ident{
				Name: identifier,
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.FLOAT,
					Value: g.Opts.ValTestCase.Float32(),
				},
			},
		}
	case "float64":
		return &ast.BasicLit{
			Kind:  token.FLOAT,
			Value: g.Opts.ValTestCase.Float64(),
		}
	case "byte":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.Byte())
	case "rune":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.Rune())
	case "uintptr":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.UInt())
	case "uint":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.UInt())
	case "uint8":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.UInt8())
	case "uint16":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.UInt16())
	case "uint32":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.UInt32())
	case "uint64":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.UInt64())
	case "int8":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.Int8())
	case "int16":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.Int16())
	case "int32":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.Int32())
	case "int64":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.Int64())
	case "complex64":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.Complex64())
	case "complex128":
		return g.numericBasicType(identifier, g.Opts.ValTestCase.Complex128())
	default:
		log.Warningf("basic lit not implemented yet: %s", identifier)
	}
	return &ast.BasicLit{}
}

func (g *TestCase) numericBasicType(identifier, value string) ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.Ident{
			Name: identifier,
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.INT,
				Value: value,
			},
		},
	}
}
