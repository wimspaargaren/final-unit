package testcase

import (
	"go/ast"

	log "github.com/sirupsen/logrus"
)

// DuplMapChecker checker which identifies duplicate map keys
type DuplMapChecker struct {
	StructTypeMap    map[ast.StructType]bool
	IdentMap         map[ast.Ident]bool
	BasicLitMap      map[ast.BasicLit]bool
	StarExprMap      map[ast.StarExpr]bool
	ArrayTypeMap     map[ast.ArrayType]bool
	MapTypeMap       map[ast.MapType]bool
	ChanTypeMap      map[ast.ChanType]bool
	FuncTypeMap      map[ast.FuncType]bool
	InterfaceTypeMap map[ast.InterfaceType]bool
	SelectorExprMap  map[ast.SelectorExpr]bool
	EllipsisMap      map[ast.Ellipsis]bool
	FuncLitMap       map[ast.FuncLit]bool
}

// NewDuplMapChecker creates a new dupl map checker
func NewDuplMapChecker() *DuplMapChecker {
	return &DuplMapChecker{
		StructTypeMap:    make(map[ast.StructType]bool),
		IdentMap:         make(map[ast.Ident]bool),
		BasicLitMap:      make(map[ast.BasicLit]bool),
		StarExprMap:      make(map[ast.StarExpr]bool),
		ArrayTypeMap:     make(map[ast.ArrayType]bool),
		MapTypeMap:       make(map[ast.MapType]bool),
		ChanTypeMap:      make(map[ast.ChanType]bool),
		FuncTypeMap:      make(map[ast.FuncType]bool),
		InterfaceTypeMap: make(map[ast.InterfaceType]bool),
		SelectorExprMap:  make(map[ast.SelectorExpr]bool),
		EllipsisMap:      make(map[ast.Ellipsis]bool),
		FuncLitMap:       make(map[ast.FuncLit]bool),
	}
}

// IsDuplExpr checks if we are dealing with a duplicate map key expression
func (d *DuplMapChecker) IsDuplExpr(e ast.Expr) bool { //nolint: funlen
	switch t := e.(type) {
	case *ast.StructType:
		dupl := d.StructTypeMap[*t]
		d.StructTypeMap[*t] = true
		return dupl
	case *ast.Ident:
		dupl := d.IdentMap[*t]
		d.IdentMap[*t] = true
		return dupl
	case *ast.BasicLit:
		dupl := d.BasicLitMap[*t]
		d.BasicLitMap[*t] = true
		return dupl
	case *ast.StarExpr:
		dupl := d.StarExprMap[*t]
		d.StarExprMap[*t] = true
		return dupl
	case *ast.ArrayType:
		dupl := d.ArrayTypeMap[*t]
		d.ArrayTypeMap[*t] = true
		return dupl
	case *ast.MapType:
		dupl := d.MapTypeMap[*t]
		d.MapTypeMap[*t] = true
		return dupl
	case *ast.ChanType:
		dupl := d.ChanTypeMap[*t]
		d.ChanTypeMap[*t] = true
		return dupl
	case *ast.FuncType:
		dupl := d.FuncTypeMap[*t]
		d.FuncTypeMap[*t] = true
		return dupl
	case *ast.InterfaceType:
		dupl := d.InterfaceTypeMap[*t]
		d.InterfaceTypeMap[*t] = true
		return dupl
	case *ast.SelectorExpr:
		dupl := d.SelectorExprMap[*t]
		d.SelectorExprMap[*t] = true
		return dupl
	case *ast.Ellipsis:
		dupl := d.EllipsisMap[*t]
		d.EllipsisMap[*t] = true
		return dupl
	case *ast.FuncLit:
		dupl := d.FuncLitMap[*t]
		d.FuncLitMap[*t] = true
		return dupl
	case *ast.CallExpr:
		return d.IsDuplExpr(t.Fun)
	case *ast.CompositeLit:
		return d.IsDuplExpr(t.Type)
	default:
		log.Warningf("type: %T unknown for dupl map key check", t)
		return true
	}
}
