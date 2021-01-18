package testcase

import (
	"go/ast"

	log "github.com/sirupsen/logrus"
)

// CorrectTypeExpr corrects type expressions for imports
func (g *TestCase) CorrectTypeExpr(e ast.Expr, input *RecursionInput) ast.Expr { // nolint: funlen, gocyclo
	switch t := e.(type) {
	case *ast.ArrayType:
		return &ast.ArrayType{
			Len: g.CorrectTypeExpr(t.Len, input),
			Elt: g.CorrectTypeExpr(t.Elt, input),
		}
	case *ast.Ident:
		if g.IsBasicLit(t.Name) || g.IsError(t.Name) {
			return t
		}
		if !g.PackageInfo.IsRoot(input.pkgPointer) {
			pkg := g.PackageInfo.PkgForPointer(input.pkgPointer)
			// FIXME: we might need to return imports + we might want to create custom identifiers in order
			// to make sure we dont have duplicated import tags
			return &ast.SelectorExpr{
				Sel: t,
				X: &ast.Ident{
					Name: pkg.Name,
				},
			}
		}
		return t
	case *ast.StarExpr:
		return &ast.StarExpr{
			X: g.CorrectTypeExpr(t.X, input),
		}
	case *ast.CallExpr:
		return e
	case *ast.MapType:
		return &ast.MapType{
			Value: g.CorrectTypeExpr(t.Value, input),
			Key:   g.CorrectTypeExpr(t.Key, input),
		}
	case *ast.SelectorExpr:
		return e
	case *ast.BasicLit:
		return e
	case *ast.CompositeLit:
		return &ast.CompositeLit{
			Type:       g.CorrectTypeExpr(t.Type, input),
			Elts:       t.Elts,
			Incomplete: t.Incomplete,
		}
	case nil:
		return e
	case *ast.ChanType:
		return e
	case *ast.FuncType:
		newParamsList := &ast.FieldList{
			List: []*ast.Field{},
		}
		if t.Params != nil {
			for _, f := range t.Params.List {
				newParamsList.List = append(newParamsList.List, &ast.Field{
					Type:  g.CorrectTypeExpr(f.Type, input),
					Names: f.Names,
					Tag:   f.Tag,
				})
			}
		}

		newResultList := &ast.FieldList{
			List: []*ast.Field{},
		}
		if t.Results != nil {
			for _, f := range t.Results.List {
				newResultList.List = append(newResultList.List, &ast.Field{
					Type:  g.CorrectTypeExpr(f.Type, input),
					Names: f.Names,
					Tag:   f.Tag,
				})
			}
		}
		return &ast.FuncType{
			Params:  newParamsList,
			Results: newResultList,
		}
	case *ast.InterfaceType:
		newFields := &ast.FieldList{
			List: []*ast.Field{},
		}
		if t.Methods != nil {
			for _, f := range t.Methods.List {
				newFields.List = append(newFields.List, &ast.Field{
					Type:  g.CorrectTypeExpr(f.Type, input),
					Names: f.Names,
					Tag:   f.Tag,
				})
			}
		}

		return &ast.InterfaceType{
			Interface:  t.Interface,
			Incomplete: t.Incomplete,
			Methods:    newFields,
		}
	case *ast.UnaryExpr:
		return &ast.UnaryExpr{
			Op:    t.Op,
			OpPos: t.OpPos,
			X:     g.CorrectTypeExpr(t.X, input),
		}
	case *ast.StructType:
		newFields := &ast.FieldList{
			List: []*ast.Field{},
		}
		if t.Fields != nil {
			for _, f := range t.Fields.List {
				newFields.List = append(newFields.List, &ast.Field{
					Type:  g.CorrectTypeExpr(f.Type, input),
					Names: f.Names,
					Tag:   f.Tag,
				})
			}
		}
		return &ast.StructType{
			Struct:     t.Struct,
			Fields:     newFields,
			Incomplete: t.Incomplete,
		}
	case *ast.Ellipsis:
		return &ast.Ellipsis{
			Ellipsis: t.Ellipsis,
			Elt:      g.CorrectTypeExpr(t.Elt, input),
		}
	default:
		log.Warningf("unable to correct type:  %T", t)
	}
	return e
}
