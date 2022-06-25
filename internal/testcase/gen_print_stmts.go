package testcase

import (
	"fmt"
	"go/ast"
	"go/token"
	"unicode"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/importer"
	"github.com/wimspaargaren/final-unit/internal/utils"
)

// PrintRecursionInput input object for traversing the AST
type PrintRecursionInput struct {
	e          ast.Expr
	varName    string
	pkgPointer *importer.PkgResolverPointer
	counter    CycleInfo
	prefix     []ast.Stmt
	suffix     []ast.Stmt
}

// PrintResult result of recursion
type PrintResult struct {
	Stmts []ast.Stmt
}

// ResultsToPrintStmts converts a results list to print statements
func (g *TestCase) ResultsToPrintStmts(results *ast.FieldList, funcName string, pointer *importer.PkgResolverPointer) ([]ast.Expr, *PrintResult, []ast.Stmt) {
	if results == nil {
		return []ast.Expr{}, &PrintResult{}, []ast.Stmt{}
	}
	identsRes := []ast.Expr{}
	res := &PrintResult{}
	for _, p := range results.List {
		idents, temp := g.FieldToPrintStmt(p, funcName, pointer)
		res.Stmts = append(res.Stmts, temp.Stmts...)
		identsRes = append(identsRes, idents...)
	}
	resultUsage := []ast.Stmt{}
	for _, ident := range identsRes {
		// In case of empty arrays of maps we need to make sure the identifier variable
		// is atleast used
		if t, ok := ident.(*ast.Ident); ok {
			if t.Name != "_" {
				resultUsage = append(resultUsage, &ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "_",
						},
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						t,
					},
				})
			}
		} else {
			log.Warningf("unexpected ident expression fouund")
		}
	}
	return identsRes, res, resultUsage
}

// FieldToPrintStmt convert field to print statement
func (g *TestCase) FieldToPrintStmt(field *ast.Field, funcName string, pointer *importer.PkgResolverPointer) ([]ast.Expr, *PrintResult) {
	// Named returns
	if len(field.Names) != 0 {
		expressions := []ast.Expr{}
		printResult := &PrintResult{}
		for _, n := range field.Names {
			newIdent := g.Opts.IdentGen.Create(n)

			res := g.TypeExpressionToPrintStmt(NewPrintRecursionInput(field.Type, newIdent.Name, pointer))
			printResult.Stmts = append(printResult.Stmts, res.Stmts...)
			if len(res.Stmts) == 0 {
				expressions = append(expressions, &ast.Ident{
					Name: "_",
				})
			} else {
				expressions = append(expressions, newIdent)
			}
		}
		return expressions, printResult
	}
	// Normal returns
	newIdent := g.Opts.IdentGen.Create(&ast.Ident{Name: "out"})

	res := g.TypeExpressionToPrintStmt(NewPrintRecursionInput(field.Type, newIdent.Name, pointer))
	if len(res.Stmts) == 0 {
		return []ast.Expr{&ast.Ident{
			Name: "_",
		}}, res
	}
	return []ast.Expr{newIdent}, res
}

// PreIdentToPrintStmt converts an Ident type to print statements(diff from IdentToPrintStmt, using by TypeExpressionToPrintStmt)
func (g *TestCase) PreIdentToPrintStmt(t *ast.Ident, input *PrintRecursionInput) *PrintResult {
	if t.Obj == nil {
		if g.IsBasicLit(t.Name) {
			return g.BasicExprToPrintStmt(input, t.Name, input.varName)
		}
		if g.IsError(t.Name) {
			return g.ErrExprToPrintStmt(input)
		}
		// t.Name != basic val this is from another file in the same package
		found, expr, newPointer := g.PackageInfo.FindInCurrent(input.pkgPointer, t.Name)
		if !found {
			log.Warningf("identifier not present in this file not found in other file: %s", t.Name)
		} else {
			return g.TypeExpressionToPrintStmt(&PrintRecursionInput{
				e:          expr,
				varName:    input.varName,
				pkgPointer: newPointer,
				counter:    input.counter,
				prefix:     input.prefix,
				suffix:     input.suffix,
			})
		}
	}
	// handle objects
	switch objectDeclType := t.Obj.Decl.(type) {
	// Object type
	case *ast.TypeSpec:
		switch oType := objectDeclType.Type.(type) {
		case *ast.StructType:
			return g.StructExprToPrintStmt(&PrintRecursionInput{
				e:          oType,
				varName:    input.varName,
				pkgPointer: input.pkgPointer,
				counter:    input.counter,
				prefix:     input.prefix,
				suffix:     input.suffix,
			})
		case *ast.Ident:
			return g.IdentToPrintStmt(t, oType, input)
		case *ast.ArrayType:
			// In case of  array we dont need to adjust prefix and suffix
			return g.ArrayExprToPrintStmt(oType, input)
		case *ast.MapType:
			// In case of map we dont need to adjust prefix and suffix
			return g.MapExprToPrintStmt(oType, input)
			// ignore types
		case *ast.ChanType,
			*ast.InterfaceType,
			*ast.FuncType:
			return &PrintResult{}
		case *ast.SelectorExpr:
			return g.SelectorExprToPrintStmt(&PrintRecursionInput{
				e:          oType,
				counter:    input.counter,
				pkgPointer: input.pkgPointer,
				prefix:     input.prefix,
				suffix:     input.suffix,
				varName:    input.varName,
			})
		default:
			log.Warningf("Unsupported validate type: %T", oType)
			return &PrintResult{}
		}
	default:
		log.Warningf("unimplemented object declaration type")
		return &PrintResult{}
	}
}

// TypeExpressionToPrintStmt converts type expression of a result to print statement
func (g *TestCase) TypeExpressionToPrintStmt(input *PrintRecursionInput) *PrintResult {
	switch t := input.e.(type) {
	case *ast.Ident:
		return g.PreIdentToPrintStmt(t, input)
	case *ast.ArrayType:
		return g.ArrayExprToPrintStmt(t, input)
	case *ast.StarExpr:
		return g.StarExprToPrintStmt(t, input)
	case *ast.MapType:
		return g.MapExprToPrintStmt(t, input)
	case *ast.SelectorExpr:
		return g.SelectorExprToPrintStmt(input)
	// ignore types:
	case *ast.ChanType,
		*ast.InterfaceType,
		// FIXME: implement struct type
		*ast.StructType,
		*ast.FuncType:
		return &PrintResult{}
	default:
		log.Warningf("Unsupported print stmt: %T", t)
		return &PrintResult{}
	}
}

// IdentToPrintStmt converts ident to print statement
func (g *TestCase) IdentToPrintStmt(t, oType *ast.Ident, input *PrintRecursionInput) *PrintResult {
	prefix := CreatePrintfStmt([]ast.Expr{
		BasicLitString(`{ "type": "%s", "var_name": "%s", "child": `),
		BasicLitString("custom"),
		BasicLitString(t.Obj.Name),
	})
	suffix := CreatePrintfStmt([]ast.Expr{
		BasicLitString(`}`),
	})
	return g.TypeExpressionToPrintStmt(&PrintRecursionInput{
		e:          oType,
		varName:    input.varName,
		pkgPointer: input.pkgPointer,
		counter:    input.counter,
		prefix:     append(input.prefix, prefix),
		suffix:     append(input.suffix, suffix),
	})
}

// SelectorExprToPrintStmt converts a selector expression to a print statement
func (g *TestCase) SelectorExprToPrintStmt(input *PrintRecursionInput) *PrintResult {
	t, ok := input.e.(*ast.SelectorExpr)
	// Sanity check
	if !ok {
		log.Warningf("SelectorExprToPrintStmt is not  used correctly: %T", input.e)
		return &PrintResult{}
	}

	// Resolve imports
	if selectorIdent, ok := t.X.(*ast.Ident); ok {
		found, expr, newPointer := g.PackageInfo.FindImport(input.pkgPointer, selectorIdent.Name, t.Sel.Name)
		if newPointer == nil {
			log.Warning("new pointer nil")
		}
		if !found {
			log.Warningf("identifier not found in imports: %s, ident: %s", selectorIdent.Name, t.Sel.Name)
			return &PrintResult{}
		}
		res := g.TypeExpressionToPrintStmt(&PrintRecursionInput{
			e:          expr,
			varName:    input.varName,
			pkgPointer: newPointer,
			counter:    input.counter,
			prefix:     input.prefix,
			suffix:     input.suffix,
		})

		return res
	}
	log.Warningf("unimplemented selector X: %T", t.Sel)
	return &PrintResult{}
}

// ErrExprToPrintStmt converts err expression to print statement
func (g *TestCase) ErrExprToPrintStmt(input *PrintRecursionInput) *PrintResult {
	ifStmt := &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  &ast.Ident{Name: input.varName},
			Op: token.EQL,
			Y:  &ast.Ident{Name: "nil"},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				CreatePrintfStmt([]ast.Expr{
					BasicLitString(`{ "type": "%s", "var_name": "%s", "val": "nil" } `),
					BasicLitString("error"),
					BasicLitString(input.varName),
				}),
				Println(),
			},
		},
		Else: &ast.BlockStmt{
			List: []ast.Stmt{
				CreatePrintfStmt([]ast.Expr{
					BasicLitString(`{ "type": "%s", "var_name": "%s", "val": "notnil" } `),
					BasicLitString("error"),
					BasicLitString(input.varName),
				}),
				Println(),
			},
		},
	}
	return &PrintResult{
		Stmts: []ast.Stmt{ifStmt},
	}
}

// StarExprToPrintStmt converts a star expression to a print statement
func (g *TestCase) StarExprToPrintStmt(t *ast.StarExpr, input *PrintRecursionInput) *PrintResult {
	ifNilRes := []ast.Stmt{}
	ifNilRes = append(ifNilRes, input.prefix...)
	ifNilRes = append(ifNilRes, CreatePrintfStmt([]ast.Expr{
		BasicLitString(`{ "type": "%s", "var_name": "%s", "val": "nil" } `),
		BasicLitString("pointer"),
		BasicLitString(input.varName),
	}))
	ifNilRes = append(ifNilRes, input.suffix...)
	ifNilRes = append(ifNilRes, Println())
	ifStmt := &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  &ast.Ident{Name: input.varName},
			Op: token.EQL,
			Y:  &ast.Ident{Name: "nil"},
		},
		Body: &ast.BlockStmt{
			List: ifNilRes,
		},
	}
	// Create new identifier for pointer value
	pointerValIdent := g.Opts.IdentGen.Create(&ast.Ident{Name: "pointerOut"})
	newValStmt := &ast.AssignStmt{
		Lhs: []ast.Expr{
			pointerValIdent,
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.UnaryExpr{
				Op: token.MUL,
				X: &ast.Ident{
					Name: input.varName,
				},
			},
		},
	}
	elseStmts := []ast.Stmt{}
	elseStmts = append(elseStmts, newValStmt)
	prefix := CreatePrintfStmt([]ast.Expr{
		BasicLitString(`{ "type": "%s", "var_name": "%s", "child": `),
		BasicLitString("pointer"),
		BasicLitString(input.varName),
	})
	suffix := CreatePrintfStmt([]ast.Expr{
		BasicLitString(`}`),
	})
	res := g.TypeExpressionToPrintStmt(&PrintRecursionInput{
		e:          t.X,
		varName:    pointerValIdent.Name,
		pkgPointer: input.pkgPointer,
		counter:    input.counter,
		prefix:     append(input.prefix, prefix),
		suffix:     append(input.suffix, suffix),
	})
	ifStmt.Else = &ast.BlockStmt{
		List: append(elseStmts, res.Stmts...),
	}

	return &PrintResult{
		Stmts: []ast.Stmt{ifStmt},
	}
}

// StructExprToPrintStmt converts a struct expression to print statement
func (g *TestCase) StructExprToPrintStmt(input *PrintRecursionInput) *PrintResult { // nolint: funlen
	t, ok := input.e.(*ast.StructType)
	// Sanity check
	if !ok {
		log.Warningf("StructExprToPrintStmt is not  used correctly: %T", input.e)
		return &PrintResult{}
	}
	// Store current struct in memory
	input.counter.StructByStruct[t]++

	// Detect struct cycles!
	// If cycle exceeds max recursion val return, we get into an infinite loop otherwise
	if input.counter.StructByStruct[t] > g.Opts.MaxRecursion && ok {
		// Return memory
		return &PrintResult{
			Stmts: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "_",
						},
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.Ident{
							Name: input.varName,
						},
					},
				},
			},
		}
	}

	totalRes := &PrintResult{}

	totalRes.Stmts = append(totalRes.Stmts, &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				Name: "_",
			},
		},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.Ident{Name: input.varName},
		},
	})
	for _, field := range t.Fields.List {
		// Directly nested struct is indicated by field without names
		if len(field.Names) == 0 {
			n := g.GetUnnamedStructIdent(field.Type, &RecursionInput{
				e:          input.e,
				counter:    input.counter,
				pkgPointer: input.pkgPointer,
				varName:    input.varName,
			})
			if !g.PackageInfo.IsRoot(input.pkgPointer) {
				isLower := unicode.IsLower(rune(n.Name[0]))
				if isLower {
					continue
				}
			}
			prefix := CreatePrintfStmt([]ast.Expr{
				BasicLitString(`{ "type": "%s", "var_name": "%s", "child": `),
				BasicLitString("struct"),
				BasicLitString(input.varName),
			})
			suffix := CreatePrintfStmt([]ast.Expr{
				BasicLitString(`}`),
			})
			res := g.TypeExpressionToPrintStmt(&PrintRecursionInput{
				e:          field.Type,
				varName:    input.varName + "." + n.Name,
				pkgPointer: input.pkgPointer,
				counter:    input.counter,
				prefix:     append(input.prefix, prefix),
				suffix:     append(input.suffix, suffix),
			})
			totalRes.Stmts = append(totalRes.Stmts, res.Stmts...)
		}
		for _, n := range field.Names {
			if !g.PackageInfo.IsRoot(input.pkgPointer) {
				isLower := unicode.IsLower(rune(n.Name[0]))
				if isLower {
					continue
				}
			}
			prefix := CreatePrintfStmt([]ast.Expr{
				BasicLitString(`{ "type": "%s", "var_name": "%s", "child": `),
				BasicLitString("struct"),
				BasicLitString(input.varName),
			})
			suffix := CreatePrintfStmt([]ast.Expr{
				BasicLitString(`}`),
			})
			res := g.TypeExpressionToPrintStmt(&PrintRecursionInput{
				e:          field.Type,
				varName:    input.varName + "." + n.Name,
				pkgPointer: input.pkgPointer,
				counter:    input.counter,
				prefix:     append(input.prefix, prefix),
				suffix:     append(input.suffix, suffix),
			})
			totalRes.Stmts = append(totalRes.Stmts, res.Stmts...)
		}
	}

	return totalRes
}

// MapExprToPrintStmt converts a map type to print statements
func (g *TestCase) MapExprToPrintStmt(t *ast.MapType, input *PrintRecursionInput) *PrintResult {
	res := []ast.Stmt{}

	keyType, isBasicLit := g.IsBasicExpr(t.Key)
	if !isBasicLit {
		return &PrintResult{}
	}
	keyIdent := &ast.Ident{
		Name: utils.LowerCaseFirstLetter(g.Opts.VarTestCase.Generate()),
	}

	rangeStmt := &ast.RangeStmt{
		Key: keyIdent,
		Tok: token.DEFINE,
		X:   &ast.Ident{Name: input.varName},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{},
		},
	}
	prefix := CreatePrintfStmt([]ast.Expr{
		BasicLitString(`{ "type": "%s", "arr_ident": "%s", "map_key_type": "%s", "var_name": "%s", "val": "%+v", "child": `),
		BasicLitString("map"),
		BasicLitString(keyIdent.Name),
		BasicLitString(keyType),
		BasicLitString(input.varName),
		keyIdent,
	})
	suffix := CreatePrintfStmt([]ast.Expr{
		BasicLitString(`}`),
	})

	r := &PrintRecursionInput{
		e:          t.Value,
		counter:    input.counter,
		pkgPointer: input.pkgPointer,
		varName:    fmt.Sprintf("%s[%s]", input.varName, keyIdent.Name),
		prefix:     append(input.prefix, prefix),
		suffix:     append(input.suffix, suffix),
	}

	recursionRes := g.TypeExpressionToPrintStmt(r)
	if len(recursionRes.Stmts) == 0 {
		return &PrintResult{}
	}

	rangeStmt.Body.List = append(rangeStmt.Body.List, recursionRes.Stmts...)
	res = append(res, rangeStmt)
	return &PrintResult{
		Stmts: res,
	}
}

// ArrayExprToPrintStmt converts an array type to print statements
func (g *TestCase) ArrayExprToPrintStmt(t *ast.ArrayType, input *PrintRecursionInput) *PrintResult {
	res := []ast.Stmt{}

	indexIdent := &ast.Ident{
		Name: utils.LowerCaseFirstLetter(g.Opts.VarTestCase.Generate()),
	}
	stmt := &ast.ForStmt{
		Init: &ast.AssignStmt{
			Lhs: []ast.Expr{indexIdent},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.INT,
					Value: "0",
				},
			},
		},
		Cond: &ast.BinaryExpr{
			X:  indexIdent,
			Op: token.LSS,
			Y: &ast.CallExpr{
				Fun: &ast.Ident{Name: "len"},
				Args: []ast.Expr{
					&ast.Ident{Name: input.varName},
				},
			},
		},
		Post: &ast.IncDecStmt{
			X:   indexIdent,
			Tok: token.INC,
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{},
		},
	}
	prefix := CreatePrintfStmt([]ast.Expr{
		BasicLitString(`{ "type": "%s", "arr_ident": "%s", "var_name": "%s", "val": "%+v", "child": `),
		BasicLitString("arr"),
		BasicLitString(indexIdent.Name),
		BasicLitString(input.varName),
		indexIdent,
	})
	suffix := CreatePrintfStmt([]ast.Expr{
		BasicLitString(`}`),
	})
	r := &PrintRecursionInput{
		e:          t.Elt,
		counter:    input.counter,
		pkgPointer: input.pkgPointer,
		varName:    fmt.Sprintf("%s[%s]", input.varName, indexIdent.Name),
		prefix:     append(input.prefix, prefix),
		suffix:     append(input.suffix, suffix),
	}

	recursionRes := g.TypeExpressionToPrintStmt(r)
	stmt.Body.List = append(stmt.Body.List, recursionRes.Stmts...)

	res = append(res, stmt)

	return &PrintResult{
		Stmts: res,
	}
}

// BasicExprToPrintStmt converts a basic identifier to a print stmt
func (g *TestCase) BasicExprToPrintStmt(input *PrintRecursionInput, typeIdentifier, varIdentifier string) *PrintResult {
	// nolint: goconst
	switch typeIdentifier {
	case "bool",
		"float32",
		"float64",
		"complex64",
		"complex128",
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
		"int":
		res := []ast.Stmt{}
		res = append(res, input.prefix...)
		res = append(res, CreatePrintfStmt([]ast.Expr{
			BasicLitString(`{ "type": "%s", "var_name": "%s", "val": "%#v"}`),
			BasicLitString(typeIdentifier),
			BasicLitString(varIdentifier),
			&ast.Ident{Name: varIdentifier},
		}))
		res = append(res, input.suffix...)
		res = append(res, Println())
		return &PrintResult{
			Stmts: res,
		}
	case "string":
		res := []ast.Stmt{}
		res = append(res, input.prefix...)
		res = append(res, CreatePrintfStmt([]ast.Expr{
			BasicLitString(`{ "type": "%s", "var_name": "%s", "val": %#v}`),
			BasicLitString(typeIdentifier),
			BasicLitString(varIdentifier),
			&ast.Ident{Name: varIdentifier},
		}))
		res = append(res, input.suffix...)
		res = append(res, Println())
		return &PrintResult{
			Stmts: res,
		}
	default:
		return &PrintResult{}
	}
}

// CreatePrintfStmt creates sprintf statement from a list of expressions
func CreatePrintfStmt(args []ast.Expr) ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.Ident{
					Name: "fmt",
				},
				Sel: &ast.Ident{
					Name: "Printf",
				},
			},
			Args: args,
		},
	}
}

// Println creates println statement
func Println() ast.Stmt {
	return PrintlnVal("")
}

// PrintlnVal create println statement for input value
func PrintlnVal(val string) ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.Ident{
					Name: "fmt",
				},
				Sel: &ast.Ident{
					Name: "Println",
				},
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf("\"%s\"", val),
				},
			},
		},
	}
}

// BasicLitString creates an ast string expression from value
func BasicLitString(val string) ast.Expr {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: fmt.Sprintf("`%s`", val),
	}
}
