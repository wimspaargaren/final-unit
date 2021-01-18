package gen

import (
	"encoding/json"
	"go/ast"
	"unicode"

	log "github.com/sirupsen/logrus"
)

// NewRecursionInputWithExpr helper function for creating new recursion inputs
func NewRecursionInputWithExpr(e ast.Expr, input *RecursionInput) *RecursionInput {
	return &RecursionInput{
		e:          e,
		counter:    input.counter,
		pkgPointer: input.pkgPointer,
		varName:    input.varName,
	}
}

// ShouldReturnForInterface checks if should return because unimplementable interface
func (g *Generator) ShouldReturnForInterface(expr ast.Expr, input *RecursionInput) bool {
	if interfaceType, ok := expr.(*ast.InterfaceType); ok {
		canGen := g.CheckIfCanGenInterface(interfaceType, input)
		return !canGen
	}
	return false
}

// CheckIfCanGenInterface if any non exported values are present in an interface, we need to return it as nil
func (g *Generator) CheckIfCanGenInterface(t *ast.InterfaceType, input *RecursionInput) bool {
	// Detect interface cycles!
	input.counter.Interfaces[t]++

	// If cycle exceeds max recursion val return, we get into an infinite loop otherwise
	if input.counter.Interfaces[t] > g.Opts.MaxRecursion {
		// In case cyclo we cant determine if can gen so make sure to return false
		return false
	}
	if t.Methods == nil {
		return true
	}
	interfaceOK := true
	for _, f := range t.Methods.List {
		if _, ok := f.Type.(*ast.FuncType); ok {
			if !g.PackageInfo.IsRoot(input.pkgPointer) {
				isLower := unicode.IsLower(rune(f.Names[0].Name[0]))
				if isLower {
					return false
				}
			}
		}
		interfaceOK = interfaceOK && g.CheckIfCanGenExpr(NewRecursionInputWithExpr(f.Type, input))
	}
	// loop then recurse until identifiers
	return interfaceOK
}

// ShouldReturnForFunc checks if should return because unimplementable func type
func (g *Generator) ShouldReturnForFunc(expr ast.Expr, input *RecursionInput) bool {
	if funcType, ok := expr.(*ast.FuncType); ok {
		canGen := g.CheckIfCanGenFunc(funcType, input)
		return !canGen
	}
	return false
}

// CheckIfCanGenFunc check if we are able to gen func implementation
func (g *Generator) CheckIfCanGenFunc(f *ast.FuncType, input *RecursionInput) bool {
	funcOK := true
	if f.Results == nil {
		return true
	}
	for _, f := range f.Results.List {
		if _, ok := f.Type.(*ast.FuncType); ok {
			if !g.PackageInfo.IsRoot(input.pkgPointer) {
				isLower := unicode.IsLower(rune(f.Names[0].Name[0]))
				if isLower {
					return false
				}
			}
		}

		funcOK = funcOK && g.CheckIfCanGenExpr(NewRecursionInputWithExpr(f.Type, input))
	}
	return funcOK
}

// CheckIfCanGenExpr recurse interface expression and discover if we find any non exported expressions
func (g *Generator) CheckIfCanGenExpr(input *RecursionInput) bool { // nolint: funlen, gocyclo, gocognit
	if input.e == nil {
		return true
	}
	b, err := json.Marshal(input.e)
	canGen, ok := g.Dynamic.CanGenInterface[string(b)]
	if ok && err == nil {
		return canGen
	}
	switch t := input.e.(type) {
	case *ast.Ident:
		if t.Obj == nil {
			if g.IsBasicLit(t.Name) {
				return true
			}
			if g.IsError(t.Name) {
				return true
			}
			// t.Name != basic val this is from another file in the same package
			found, expr, newPointer := g.PackageInfo.FindInCurrent(input.pkgPointer, t.Name)
			if !found {
				log.Warningf("identifier not present in this file not found in other file: %s, dir: %s", t.Name, input.pkgPointer.Dir)
				return false
			}
			return g.CheckIfCanGenExpr(&RecursionInput{
				e:          expr,
				varName:    t.Name,
				pkgPointer: newPointer,
				counter:    input.counter,
			})
		}

		// handle objects
		switch objectDeclType := t.Obj.Decl.(type) {
		// Object type
		case *ast.TypeSpec:
			if !g.PackageInfo.IsRoot(input.pkgPointer) {
				isLower := unicode.IsLower(rune(objectDeclType.Name.Name[0]))
				if isLower {
					return false
				}
			}
			switch oType := objectDeclType.Type.(type) {
			case *ast.StructType:
				return g.CheckIfCanGenExpr(&RecursionInput{
					e:          oType,
					varName:    objectDeclType.Name.Name,
					pkgPointer: input.pkgPointer,
					counter:    input.counter,
				})
			default:
				isOK := g.CheckIfCanGenExpr(&RecursionInput{
					e:          objectDeclType.Type,
					varName:    input.varName,
					pkgPointer: input.pkgPointer,
					counter:    input.counter,
				})
				// we don't need to create call expression for interface type
				if _, ok := objectDeclType.Type.(*ast.InterfaceType); ok {
					return true
				}
				return isOK
			}
		default:
			log.Warningf("unimplemented object declaration type")
			return false
		}
	case *ast.FuncType:
		funcOK := true
		if t.Params.List != nil {
			for _, f := range t.Params.List {
				funcOK = funcOK && g.CheckIfCanGenExpr(NewRecursionInputWithExpr(f.Type, input))
			}
		}

		if t.Results != nil && t.Results.List != nil {
			for _, f := range t.Results.List {
				funcOK = funcOK && g.CheckIfCanGenExpr(NewRecursionInputWithExpr(f.Type, input))
			}
		}
		return funcOK
	case *ast.ArrayType:
		return g.CheckIfCanGenExpr(NewRecursionInputWithExpr(t.Elt, input)) &&
			g.CheckIfCanGenExpr(NewRecursionInputWithExpr(t.Len, input))
	case *ast.SelectorExpr:
		if selectorIdent, ok := t.X.(*ast.Ident); ok {
			found, expr, newPointer := g.PackageInfo.FindImport(input.pkgPointer, selectorIdent.Name, t.Sel.Name)
			if newPointer == nil {
				log.Warning("new pointer nil")
			}
			if !found {
				log.Warningf("identifier not found in imports: %s, sel: %s", selectorIdent.Name, t.Sel.Name)
				return false
			}
			return g.CheckIfCanGenExpr(&RecursionInput{
				e:          expr,
				counter:    input.counter,
				pkgPointer: newPointer,
				varName:    input.varName,
			})
		}
		log.Warningf("unexpected selector expr")
		return false
	case *ast.StructType:
		isStructOK := true
		if t.Fields.List != nil {
			for _, f := range t.Fields.List {
				isStructOK = isStructOK && g.CheckIfCanGenExpr(NewRecursionInputWithExpr(f.Type, input))
			}
		}
		return isStructOK
	case *ast.StarExpr:
		return g.CheckIfCanGenExpr(NewRecursionInputWithExpr(t.X, input))
	case *ast.InterfaceType:
		return g.CheckIfCanGenInterface(t, input)
	case *ast.ChanType:
		return g.CheckIfCanGenExpr(NewRecursionInputWithExpr(t.Value, input))
	case *ast.MapType:
		keyOK := g.CheckIfCanGenExpr(NewRecursionInputWithExpr(t.Key, input))
		valOK := g.CheckIfCanGenExpr(NewRecursionInputWithExpr(t.Value, input))
		return keyOK && valOK
	case *ast.BasicLit:
		return true
	case *ast.ParenExpr:
		return g.CheckIfCanGenExpr(NewRecursionInputWithExpr(t.X, input))
	default:
		log.Warningf("Implement interface recurse type: %T", t)
		return false
	}
}
