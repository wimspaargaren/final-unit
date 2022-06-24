// Package testcase contains functionality for generating all necessary information in order to create test cases
package testcase

import (
	"fmt"
	"go/ast"
	"go/token"
	"path/filepath"
	"unicode"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/decorator"
	"github.com/wimspaargaren/final-unit/internal/ident"
	"github.com/wimspaargaren/final-unit/internal/identlist"
	"github.com/wimspaargaren/final-unit/internal/importer"
	"github.com/wimspaargaren/final-unit/internal/runtime"
	"github.com/wimspaargaren/final-unit/pkg/values"
	"github.com/wimspaargaren/final-unit/pkg/variables"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Options test case generation options
type Options struct {
	ValTestCase  values.IGen
	VarTestCase  variables.IGen
	IdentGen     ident.IGen
	MaxRecursion int
}

// TestCase contains all information for generating a test case
type TestCase struct {
	// Properties used for generating mutations when doing crossover in evolution
	FuncDecl    *ast.FuncDecl
	Pointer     *importer.PkgResolverPointer
	PackageInfo *importer.PackageInfo
	Opts        Options
	Deco        *decorator.Deco
	Dynamic     Dynamic
	RunTimeInfo *runtime.Info

	// Properties used to create value stmts in test cases
	Decls      []string
	Stmts      []string
	FuncStmt   string
	ChanIdents []string
	// Properties used for creating assert stmts in test cases
	ResultStmts      []string
	ResultUsageStmts []string
	FuncPrintStmt    string
}

// HasPrintStmts check if any print statements are generated for current test case
func (g *TestCase) HasPrintStmts() bool {
	return g.FuncPrintStmt != ""
}

// HasChan reports if test case has a channel receiver
func (g *TestCase) HasChan() bool {
	return len(g.ChanIdents) > 0
}

// New creates a new test case
func New(f *ast.FuncDecl,
	pointer *importer.PkgResolverPointer,
	pkgInfo *importer.PackageInfo,
	opts Options,
	decorator *decorator.Deco,
) *TestCase {
	return &TestCase{
		FuncDecl:    f,
		Pointer:     pointer,
		PackageInfo: pkgInfo,
		Opts:        opts,
		Deco:        decorator,
		RunTimeInfo: runtime.NewInfo(runtime.NewTestifySuitePrinter("s")),
		Dynamic: Dynamic{
			CanGenInterface: make(map[string]bool),
		},
	}
}

// Dynamic struct for performance improvements using dynamic programming
type Dynamic struct {
	CanGenInterface map[string]bool
}

// Create converts a function declaration to a list of assignment statements
// and declaration statements
func (g *TestCase) Create() {
	// Reset local scope counter whenever creating new testcase
	g.Opts.IdentGen.ResetLocal()

	// Get receiver statements and declarations
	receiverResult := g.GetFuncReceiverStmts(g.FuncDecl.Recv, g.FuncDecl.Name.Name, g.Pointer)
	// Get param statements and declarations
	fieldToAssignResult := g.FieldToAssignStmts(g.FuncDecl.Type.Params, g.FuncDecl.Name.Name, g.Pointer)
	// Create print statements for generating assert statements
	identsPrint, results, resultUsages := g.ResultsToPrintStmts(g.FuncDecl.Type.Results, g.FuncDecl.Name.Name, g.Pointer)

	// Create function statements for just calling(used for evolution execution)
	// as well as assigning the return values(used for creating assert stmts)
	funcStmt, funcPrintStmt := g.FuncDeclToExprStmt(g.FuncDecl, receiverResult.Idents, fieldToAssignResult.Idents, identsPrint)
	tempStmts := receiverResult.Statements
	tempStmts = append(tempStmts, fieldToAssignResult.Statements...)
	resStmts := []string{}
	for _, tempStmt := range tempStmts {
		resStmts = append(resStmts, MustPrettyPrintElement(tempStmt))
	}
	tempDecls := receiverResult.Declarations
	tempDecls = append(tempDecls, fieldToAssignResult.Declarations...)
	resDecls := []string{}
	for _, tempDecl := range tempDecls {
		resDecls = append(resDecls, MustPrettyPrintElement(tempDecl))
	}

	resultStmts := []string{}
	for _, resultStmt := range results.Stmts {
		resultStmts = append(resultStmts, MustPrettyPrintElement(resultStmt))
	}

	resultUsageStmts := []string{}
	for _, resultUsage := range resultUsages {
		resultUsageStmts = append(resultUsageStmts, MustPrettyPrintElement(resultUsage))
	}

	chanIdents := []string{}
	for _, chanIdent := range receiverResult.ChanIdents {
		chanIdents = append(chanIdents, MustPrettyPrintElement(chanIdent))
	}

	for _, chanIdent := range fieldToAssignResult.ChanIdents {
		chanIdents = append(chanIdents, MustPrettyPrintElement(chanIdent))
	}

	// Create resulting test case
	g.Stmts = resStmts
	g.Decls = resDecls
	g.FuncStmt = MustPrettyPrintElement(funcStmt)
	g.ResultStmts = resultStmts
	g.ResultUsageStmts = resultUsageStmts
	g.ChanIdents = chanIdents
	// In case all output values are not verifiable funcPrintStmt is nil
	if funcPrintStmt != nil {
		g.FuncPrintStmt = MustPrettyPrintElement(funcPrintStmt)
	}
}

// FuncDeclToExprStmt converts func declaration to expression statement
func (g *TestCase) FuncDeclToExprStmt(f *ast.FuncDecl, recvIdent, paramIdent []*ast.Ident, printIdents []ast.Expr) (ast.Stmt, ast.Stmt) {
	callExpr := &ast.CallExpr{
		Fun: f.Name,
	}
	if f.Recv != nil &&
		// sanity checks
		len(f.Recv.List) == 1 &&
		len(f.Recv.List[0].Names) == 1 {
		if len(recvIdent) != 1 {
			log.Warningf("receiver ident should always be 1, but is: %d", len(recvIdent))
		}
		callExpr = &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   recvIdent[0],
				Sel: f.Name,
			},
		}
	}

	// Put all parameter identifiers into the call expression
	for _, x := range paramIdent {
		callExpr.Args = append(callExpr.Args, x)
	}
	assignToken := token.DEFINE
	shouldUseAssign := true
	// If only one param is present and is "_" use "=" for assign, as we can't assign _ := someVar
	for _, e := range printIdents {
		if i, ok := e.(*ast.Ident); ok {
			if i.Name != "_" {
				shouldUseAssign = false
			}
		}
	}
	if shouldUseAssign {
		assignToken = token.ASSIGN
	}
	var resPrint ast.Stmt = &ast.AssignStmt{
		Lhs: printIdents,
		Tok: assignToken,
		Rhs: []ast.Expr{
			callExpr,
		},
	}
	if len(printIdents) == 0 {
		resPrint = nil
	}
	return &ast.ExprStmt{
		X: callExpr,
	}, resPrint
}

// GetFuncReceiverStmts retrieve statements for receiverr field
func (g *TestCase) GetFuncReceiverStmts(recv *ast.FieldList, funcName string, pointer *importer.PkgResolverPointer) *FieldToAssignRes {
	result := &FieldToAssignRes{}

	if recv == nil {
		return result
	}
	_, fileName := filepath.Split(pointer.File)

	hasVal := g.Deco.HasReceiverVal(fileName, funcName)
	if hasVal && g.Opts.ValTestCase.DecoratorVal() {
		newIdent := g.Opts.IdentGen.Create(&ast.Ident{Name: funcName})
		result.Idents = append(result.Idents, newIdent)
		values := g.Deco.GetReceiverVal(fileName, funcName)
		result.Statements = append(result.Statements, assignStmt(newIdent, values[g.Opts.ValTestCase.DecoratorIndex(len(values))].Call))
		return result
	}

	// Generate assignment for receiver field
	for _, field := range recv.List {
		result.AppendRes(g.FieldToAssignStmt(field, funcName, pointer))
	}
	return result
}

// FieldToAssignRes result for converting func fields to assign stmts
type FieldToAssignRes struct {
	Idents       []*ast.Ident
	Statements   []ast.Stmt
	Declarations []ast.Decl
	// FIXME:
	ChanIdents []*ast.Ident
	// ChanStmts []ast.Stmt
}

// AppendRes appends other into this
func (res *FieldToAssignRes) AppendRes(otherRes *FieldToAssignRes) {
	res.Append(otherRes.Idents, otherRes.ChanIdents, otherRes.Statements, otherRes.Declarations)
}

// Append appends idents, statements and declarations to a fieldToAssignRes
func (res *FieldToAssignRes) Append(idents, chanIdents []*ast.Ident, stmts []ast.Stmt, decls []ast.Decl) {
	res.Idents = append(res.Idents, idents...)
	res.ChanIdents = append(res.ChanIdents, chanIdents...)
	res.Statements = append(res.Statements, stmts...)
	res.Declarations = append(res.Declarations, decls...)
}

// FieldToAssignStmts converts a parameter to assignment statements
func (g *TestCase) FieldToAssignStmts(params *ast.FieldList, funcName string, pointer *importer.PkgResolverPointer) *FieldToAssignRes {
	result := &FieldToAssignRes{}
	for _, p := range params.List {
		result.AppendRes(g.FieldToAssignStmt(p, funcName, pointer))
	}
	return result
}

// CycleInfo stores information about possible cycles in parsed code
type CycleInfo struct {
	Structs        map[string]int
	StructByStruct map[*ast.StructType]int
	Interfaces     map[*ast.InterfaceType]int
	StructMem      map[*ast.StructType]ast.Expr
	InterfaceMem   map[*ast.InterfaceType]ast.Decl
}

// RecursionInput input object for traversing the AST
type RecursionInput struct {
	e          ast.Expr
	varName    string
	pkgPointer *importer.PkgResolverPointer
	counter    CycleInfo
	identList  identlist.IdentList
}

// FreshCycleInfo creates a fresh cycle info struct
func FreshCycleInfo() CycleInfo {
	return CycleInfo{
		Structs:        make(map[string]int),
		StructByStruct: make(map[*ast.StructType]int),
		Interfaces:     make(map[*ast.InterfaceType]int),
		InterfaceMem:   make(map[*ast.InterfaceType]ast.Decl),
		StructMem:      make(map[*ast.StructType]ast.Expr),
	}
}

// Copy copies recursion input with a fresh cycle info
func (r *RecursionInput) Copy() *RecursionInput {
	return &RecursionInput{
		e:          r.e,
		counter:    FreshCycleInfo(),
		pkgPointer: r.pkgPointer,
		varName:    r.varName,
		identList:  r.identList,
	}
}

// FieldToAssignStmt converts a function parameter to an assignment field
func (g *TestCase) FieldToAssignStmt(p *ast.Field, funcName string, pointer *importer.PkgResolverPointer) *FieldToAssignRes {
	res := []ast.Stmt{}
	decls := []ast.Decl{}
	idents := []*ast.Ident{}
	chanIdents := []*ast.Ident{}
	for _, param := range p.Names {
		newIdent := g.Opts.IdentGen.Create(param)
		_, fileName := filepath.Split(pointer.File)

		// If decorators have been specified use to generate value statements
		hasVal := g.Deco.HasVal(fileName, funcName, param.Name)
		if hasVal && g.Opts.ValTestCase.DecoratorVal() {
			idents = append(idents, newIdent)
			values := g.Deco.GetVal(fileName, funcName, param.Name)
			res = append(res, assignStmt(newIdent, values[g.Opts.ValTestCase.DecoratorIndex(len(values))].Call))
			continue
		}
		i := NewRecursionInput(p.Type, newIdent.Name, pointer, newIdent)

		recursionResult := g.TypeExprToValExpr(i)

		res = append(res, recursionResult.Statements...)
		decls = append(decls, recursionResult.Declarations...)
		res = append(res, assignStmt(newIdent, recursionResult.Expr))
		idents = append(idents, newIdent)
		chanIdents = append(chanIdents, recursionResult.ChanIdents...)
	}

	return &FieldToAssignRes{
		Idents:       idents,
		Declarations: decls,
		Statements:   res,
		ChanIdents:   chanIdents,
	}
}

// TypeExprToValExprRes result for converting input to value expressions
type TypeExprToValExprRes struct {
	Expr         ast.Expr
	Statements   []ast.Stmt
	Declarations []ast.Decl
	// FIXME:
	ChanIdents []*ast.Ident
	ChanStmts  []ast.Stmt
	ChanDecls  []ast.Stmt
}

// Merge merges two type exr to val expr without using the expression
func (t *TypeExprToValExprRes) Merge(other *TypeExprToValExprRes) {
	t.Statements = append(t.Statements, other.Statements...)
	t.Declarations = append(t.Declarations, other.Declarations...)

	t.ChanIdents = append(t.ChanIdents, other.ChanIdents...)
	t.ChanStmts = append(t.ChanStmts, other.ChanStmts...)
	t.ChanDecls = append(t.ChanDecls, other.ChanDecls...)
}

// EmptyResult creates empty result
func EmptyResult() *TypeExprToValExprRes {
	return &TypeExprToValExprRes{
		Expr:         &ast.BasicLit{},
		Statements:   []ast.Stmt{},
		Declarations: []ast.Decl{},
	}
}

// TypeExprToValExpr converts a type expression, the type definition in a function parameter, to an expression used in an assignment statement
func (g *TestCase) TypeExprToValExpr(input *RecursionInput) *TypeExprToValExprRes {
	switch t := input.e.(type) {
	// Handle unnamed structs
	case *ast.StructType:
		return g.UnnamedStructToValExpr(t, input)
	// Handle identifiers
	case *ast.Ident:
		input.identList.Add(t)
		if t.Obj == nil {
			return g.IdentWithNilObjectToValExpr(t, input)
		}
		// handle objects
		switch objectDeclType := t.Obj.Decl.(type) {
		// Object type
		case *ast.TypeSpec:
			return g.TypeSpecToValExpr(t, objectDeclType, input)
		default:
			log.Warningf("unimplemented object declaration type")
			return EmptyResult()
		}
	// Handle pointer typess
	case *ast.StarExpr:
		return g.StarExprToValExpr(input)
	// Handle array types
	case *ast.ArrayType:
		return g.ArrayExprToValExpr(input)
	// Handle map typesc
	case *ast.MapType:
		return g.MapExprToValExpr(input)
	// Handle channel types
	case *ast.ChanType:
		return g.ChanTypeToValExpr(t, input)
	// Handle function types
	case *ast.FuncType:
		return g.FuncTypeToValExpr(input)
	// Handle interface type
	case *ast.InterfaceType:
		return g.InterfaceTypeToValExpr(input)
	// Handle selectors e.g. pkg.Something
	case *ast.SelectorExpr:
		return g.SelectorExprToValExpr(input)
	// Handle ellipsis type e.g. ...X
	case *ast.Ellipsis:
		return g.TypeExprToValExpr(&RecursionInput{
			e:          t.Elt,
			counter:    input.counter,
			pkgPointer: input.pkgPointer,
			varName:    input.varName,
			identList:  input.identList,
		})
	// Default should not be hit all types are handled accordingly
	default:
		log.Warningf("typeExprToValExpr not implemented yet: %T", t)
		return EmptyResult()
	}
}

// TypeSpecToValExpr converts type spec to val expression
// Start of type recursion
func (g *TestCase) TypeSpecToValExpr(t *ast.Ident, objectDeclType *ast.TypeSpec, input *RecursionInput) *TypeExprToValExprRes {
	switch oType := objectDeclType.Type.(type) {
	case *ast.StructType:
		return g.StructExprToValExpr(&RecursionInput{
			e:          oType,
			varName:    objectDeclType.Name.Name,
			pkgPointer: input.pkgPointer,
			counter:    input.counter,
			identList:  input.identList,
		})
	default:
		// Detect if we are dealing with ungeneratable interfaces
		shouldReturn := g.ShouldReturnForInterface(objectDeclType.Type, &RecursionInput{
			e:          objectDeclType.Type,
			counter:    FreshCycleInfo(),
			pkgPointer: input.pkgPointer,
			varName:    input.varName,
			identList:  input.identList,
		})
		if shouldReturn {
			return g.InterfaceNilFunc(t, input)
		}

		// Detect if we are dealing with ungeneratable functions
		shouldReturn = g.ShouldReturnForFunc(objectDeclType.Type, &RecursionInput{
			e:          objectDeclType.Type,
			counter:    FreshCycleInfo(),
			pkgPointer: input.pkgPointer,
			varName:    input.varName,
			identList:  input.identList,
		})
		if shouldReturn {
			return g.FuncNilFunc(t, input)
		}

		recursionResult := g.TypeExprToValExpr(&RecursionInput{
			e:          objectDeclType.Type,
			varName:    input.varName,
			pkgPointer: input.pkgPointer,
			counter:    input.counter,
			identList:  input.identList,
		})

		// we don't need to create call expression for interface type
		if _, ok := objectDeclType.Type.(*ast.InterfaceType); ok {
			return recursionResult
		}
		result := &TypeExprToValExprRes{}
		result.Merge(recursionResult)

		res := &ast.CallExpr{
			Fun:  g.CorrectTypeExpr(objectDeclType.Name, input),
			Args: []ast.Expr{recursionResult.Expr},
		}
		result.Expr = res
		return result
	}
}

// IdentWithNilObjectToValExpr converts identifier with nil object to val expression
func (g *TestCase) IdentWithNilObjectToValExpr(t *ast.Ident, input *RecursionInput) *TypeExprToValExprRes {
	if g.IsBasicLit(t.Name) {
		return &TypeExprToValExprRes{
			Expr:         g.BasicExprToValExpr(t.Name),
			Statements:   []ast.Stmt{},
			Declarations: []ast.Decl{},
		}
	}
	if g.IsError(t.Name) {
		return g.ErrExprToValExpr()
	}
	// t.Name != basic val this is from another file in the same package
	found, expr, newPointer := g.PackageInfo.FindInCurrent(input.pkgPointer, t.Name)
	if !found {
		log.Warningf("identifier not present in this file not found in other file: %s", t.Name)
	}
	return g.TypeExprToValExpr(&RecursionInput{
		e:          expr,
		varName:    t.Name,
		pkgPointer: newPointer,
		counter:    input.counter,
		identList:  input.identList,
	})
}

// UnnamedStructToValExpr converts unnamed struct to val expression
func (g *TestCase) UnnamedStructToValExpr(t *ast.StructType, input *RecursionInput) *TypeExprToValExprRes {
	/* e.g.
	type x struct{
		y struct{x int}
	}
	*/
	res := &ast.CompositeLit{
		Type: t,
	}
	return g.StructFieldsToKeyValExpr(res, input)
}

// SelectorExprToValExpr converts a selector expression to a value expression
func (g *TestCase) SelectorExprToValExpr(input *RecursionInput) *TypeExprToValExprRes {
	t, ok := input.e.(*ast.SelectorExpr)
	// Sanity check
	if !ok {
		log.Warningf("SelectorExprToValExpr is not  used correctly: %T", input.e)
		return EmptyResult()
	}

	if selectorIdent, ok := t.X.(*ast.Ident); ok {
		// Resolve imports
		found, expr, newPointer := g.PackageInfo.FindImport(input.pkgPointer, selectorIdent.Name, t.Sel.Name)
		if newPointer == nil {
			log.Warning("new pointer nil")
		}
		if !found {
			log.Warningf("identifier not found in imports: %s, expr: %v", selectorIdent.Name, t.X)
			return EmptyResult()
		}
		shouldReturn := g.ShouldReturnForInterface(expr, &RecursionInput{
			e:          expr,
			counter:    FreshCycleInfo(),
			pkgPointer: newPointer,
			varName:    input.varName,
			identList:  input.identList,
		})
		if shouldReturn {
			return g.InterfaceNilFunc(t, input)
		}
		recursionResult := g.TypeExprToValExpr(&RecursionInput{
			e:          expr,
			varName:    input.varName,
			pkgPointer: newPointer,
			counter:    input.counter,
			identList:  input.identList,
		})
		result := &TypeExprToValExprRes{}
		result.Merge(recursionResult)
		switch recursionType := recursionResult.Expr.(type) {
		case *ast.CompositeLit:
			recursionType.Type = t
			result.Expr = recursionType
			return result
		case *ast.CallExpr:
			recursionType.Fun = t
			result.Expr = recursionType
			return result
		default:
			result.Expr = recursionResult.Expr
			return result
		}
	} else {
		log.Warningf("unimplemented selector X: %T", t.Sel)
	}
	return EmptyResult()
}

// InterfaceNilFunc creates a function returning nil value for expression
func (g *TestCase) InterfaceNilFunc(t ast.Expr, input *RecursionInput) *TypeExprToValExprRes {
	return &TypeExprToValExprRes{
		Expr: &ast.CallExpr{
			Fun: &ast.FuncLit{
				Type: &ast.FuncType{
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: g.CorrectTypeExpr(t, input),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.Ident{Name: "nil"},
							},
						},
					},
				},
			},
		},
	}
}

// FuncNilFunc creates a nil func
func (g *TestCase) FuncNilFunc(t ast.Expr, input *RecursionInput) *TypeExprToValExprRes {
	return &TypeExprToValExprRes{
		Expr: &ast.CallExpr{
			Fun: g.CorrectTypeExpr(t, input),
			Args: []ast.Expr{
				&ast.Ident{Name: "nil"},
			},
		},
	}
}

// InterfaceTypeToValExpr converts an interface type to val expression and declarations
func (g *TestCase) InterfaceTypeToValExpr(input *RecursionInput) *TypeExprToValExprRes {
	t, ok := input.e.(*ast.InterfaceType)
	// Sanity check
	if !ok {
		log.Warningf("InterfaceTypeToValExpr is not  used correctly: %T", input.e)
		return EmptyResult()
	}

	if t.Incomplete {
		log.Warningf("Incomplete interface detected")
	}
	// Empty interface
	if t.Methods.List == nil {
		return &TypeExprToValExprRes{
			Expr:         g.BasicExprToValExpr(g.Opts.ValTestCase.Type()),
			Statements:   []ast.Stmt{},
			Declarations: []ast.Decl{},
		}
	}
	// Normal interface

	// Detect interface cycles!
	input.counter.Interfaces[t]++
	// Create name for interface implementation
	interfaceImplIdent := g.Opts.IdentGen.CreateGlobal(input.identList.Current())

	result := &TypeExprToValExprRes{}

	interfaceGenDecl := g.InterfaceGenDecl(input, interfaceImplIdent)
	mem, ok := input.counter.InterfaceMem[t]
	if !ok {
		input.counter.InterfaceMem[t] = interfaceGenDecl
	}

	// If cycle exceeds max recursion val return, we get into an infinite loop otherwise
	if input.counter.Interfaces[t] > g.Opts.MaxRecursion && ok {
		if x, ok := mem.(*ast.GenDecl); ok {
			if y, ok := x.Specs[0].(*ast.TypeSpec); ok {
				// Empty instance of given type spec
				return &TypeExprToValExprRes{
					Expr: &ast.UnaryExpr{
						Op: token.AND,
						X: &ast.CompositeLit{
							Type: y.Name,
						},
					},
					Statements:   []ast.Stmt{},
					Declarations: []ast.Decl{},
				}
			}
		}
	}

	// Create struct declaration for interface implementation
	result.Declarations = append(result.Declarations, g.InterfaceGenDecl(input, interfaceImplIdent))

	// Create function declaration; implementation for all methods declared in the interface
	implementationResult := g.InterfaceTypeToFuncImpl(input, interfaceImplIdent)
	result.Merge(implementationResult)

	elts := []ast.Expr{}
	result.Expr = &ast.UnaryExpr{
		Op: token.AND,
		X: &ast.CompositeLit{
			Type: interfaceImplIdent,
			Elts: elts,
		},
	}

	// return the interface implementation as pointer
	return result
}

func (g *TestCase) MethodFuncTypeToFuncImpl(funcType *ast.FuncType, method *ast.Field, input *RecursionInput, interfaceImplIdent *ast.Ident, result *TypeExprToValExprRes) *TypeExprToValExprRes {
	recursionResult := g.FuncReturnListToBodyStatements(&RecursionInput{
		e:          funcType,
		varName:    input.varName,
		pkgPointer: input.pkgPointer,
		counter:    input.counter,
		identList:  input.identList,
	})
	result.Merge(recursionResult)
	// Statements are used for body
	result.Statements = []ast.Stmt{}

	if len(method.Names) != 1 {
		log.Warningf("expected 1 method name got: %d", len(method.Names))
	}

	funcExpr := g.CorrectTypeExpr(funcType, input)
	if x, ok := funcExpr.(*ast.FuncType); ok {
		funcType = x
	}
	// Create the function declaration
	result.Declarations = append(result.Declarations, &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					// use one receiver name for consistency
					Names: []*ast.Ident{{Name: "s"}},
					Type: &ast.StarExpr{
						X: interfaceImplIdent,
					},
				},
			},
		},
		Name: method.Names[0],
		Type: funcType,
		Body: &ast.BlockStmt{
			List: recursionResult.Statements,
		},
	})
	return result
}

// InterfaceTypeToFuncImpl converts interface type to function implementation declarations
func (g *TestCase) InterfaceTypeToFuncImpl(input *RecursionInput, interfaceImplIdent *ast.Ident) *TypeExprToValExprRes {
	t, ok := input.e.(*ast.InterfaceType)
	// Sanity check
	if !ok {
		log.Warningf("InterfaceTypeToFuncImpl is not  used correctly: %T", input.e)
		return EmptyResult()
	}
	result := &TypeExprToValExprRes{}
	for _, method := range t.Methods.List {
		// Normal method definitions
		// nolint: nestif
		if funcType, ok := method.Type.(*ast.FuncType); ok {
			result = g.MethodFuncTypeToFuncImpl(funcType, method, input, interfaceImplIdent, result)
			// Nested interface
		} else if ident, ok := method.Type.(*ast.Ident); ok {
			// Resolve directly nested interfaces
			if typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec); ok {
				if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
					recursionResult := g.InterfaceTypeToFuncImpl(&RecursionInput{
						e:          interfaceType,
						counter:    input.counter,
						pkgPointer: input.pkgPointer,
						varName:    input.varName,
						identList:  input.identList,
					}, interfaceImplIdent)
					result.Merge(recursionResult)
				} else {
					log.Warningf("unexpected type spec type: %T", typeSpec.Type)
				}
			} else {
				log.Warningf("Unexpected object type: %T", ident.Obj.Decl)
			}
		} else if t, ok := method.Type.(*ast.SelectorExpr); ok {
			// In case directly nested interface as selector
			// Resolve import and recurse
			selectorIdent, ok := t.X.(*ast.Ident)
			if !ok {
				log.Warningf("identifier not found in imports: %s, ident: %s", selectorIdent.Name, selectorIdent.Name)
				return EmptyResult()
			}
			found, expr, newPointer := g.PackageInfo.FindImport(input.pkgPointer, selectorIdent.Name, t.Sel.Name)
			if newPointer == nil {
				log.Warning("new pointer nil")
			}
			if !found {
				log.Warningf("identifier not found in imports: %s, ident: %s", selectorIdent.Name, selectorIdent.Name)
				return EmptyResult()
			}
			recursionResult := g.InterfaceTypeToFuncImpl(&RecursionInput{
				e:          expr,
				counter:    input.counter,
				pkgPointer: input.pkgPointer,
				varName:    input.varName,
				identList:  input.identList,
			}, interfaceImplIdent)
			result.Merge(recursionResult)
		} else {
			log.Warningf("interface specified non functype type: %T", method.Type)
		}
	}
	return result
}

// ChanTypeToValExpr converts a chan type to a value expression
func (g *TestCase) ChanTypeToValExpr(t *ast.ChanType, input *RecursionInput) *TypeExprToValExprRes {
	newIdent := g.Opts.IdentGen.Create(input.identList.Current())

	res := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "make"},
		Args: []ast.Expr{t},
	}
	if t.Dir != ast.RECV {
		// FIXME use recevier chan values
		// recursionResult := g.TypeExprToValExpr(&RecursionInput{
		// 	e:          t.Value,
		// 	counter:    input.counter,
		// 	pkgPointer: input.pkgPointer,
		// 	varName:    input.varName,
		// })

		return &TypeExprToValExprRes{
			Expr:         newIdent,
			Statements:   []ast.Stmt{assignStmt(newIdent, res)},
			Declarations: []ast.Decl{},
			// ChanIdents:   append(recursionResult.ChanIdents, &ast.Ident{Name: input.varName}),
			ChanIdents: []*ast.Ident{{Name: newIdent.Name}},
		}
	}

	return &TypeExprToValExprRes{
		Expr:         newIdent,
		Statements:   []ast.Stmt{assignStmt(newIdent, res)},
		Declarations: []ast.Decl{},
	}
}

// FuncTypeToValExpr converts a func type to a value expression
func (g *TestCase) FuncTypeToValExpr(input *RecursionInput) *TypeExprToValExprRes {
	// Create correct type for resulting function
	resFunc := g.CorrectTypeExpr(input.e, input)

	t, ok := resFunc.(*ast.FuncType)
	// Sanity check
	if !ok {
		log.Warningf("FuncTypeToValExpr is not  used correctly: %T", input.e)
		return EmptyResult()
	}
	result := &TypeExprToValExprRes{}
	// Retrieve statements for body of function
	recursionResult := g.FuncReturnListToBodyStatements(input)
	result.Merge(recursionResult)
	// Since statements are  used in body, remove from recursion result
	result.Statements = []ast.Stmt{}
	res := &ast.FuncLit{
		Type: t,
		Body: &ast.BlockStmt{
			List: recursionResult.Statements,
		},
	}
	result.Expr = res
	return result
}

// FuncReturnListToBodyStatements converts a return list of function to statements return values
func (g *TestCase) FuncReturnListToBodyStatements(input *RecursionInput) *TypeExprToValExprRes {
	t, ok := input.e.(*ast.FuncType)
	// Sanity check
	if !ok {
		log.Warningf("FuncReturnListToBodyStatements  is not  used correctly: %T", input.e)
		return EmptyResult()
	}
	result := &TypeExprToValExprRes{}

	retVars := []ast.Expr{}
	if t.Results != nil {
		// Loop through all return values
		for _, res := range t.Results.List {
			if len(res.Names) == 0 {
				recursionResult := g.ReturnFieldToStmt(res, input)
				result.Merge(recursionResult)
				retVars = append(retVars, recursionResult.Expr)
			}
			// Loop through all names
			for range res.Names {
				recursionResult := g.ReturnFieldToStmt(res, input)
				result.Merge(recursionResult)
				retVars = append(retVars, recursionResult.Expr)
			}
		}
	}
	// Finish body with return statement
	result.Statements = append(result.Statements, &ast.ReturnStmt{
		Results: retVars,
	})
	return result
}

// ReturnFieldToStmt converts return field to statement and declarations
func (g *TestCase) ReturnFieldToStmt(res *ast.Field, input *RecursionInput) *TypeExprToValExprRes {
	newIdent := g.Opts.IdentGen.Create(&ast.Ident{Name: "o"})
	result := &TypeExprToValExprRes{
		Expr: newIdent,
	}

	recursionResult := g.TypeExprToValExpr(&RecursionInput{
		e:          res.Type,
		varName:    newIdent.Name,
		pkgPointer: input.pkgPointer,
		counter:    input.counter,
		identList:  input.identList,
	})
	result.Merge(recursionResult)
	// Create assignment statement of return value
	result.Statements = append(result.Statements, assignStmt(result.Expr, recursionResult.Expr))
	return result
}

// ArrayExprToValExpr convert  a array expression to value expression
func (g *TestCase) ArrayExprToValExpr(input *RecursionInput) *TypeExprToValExprRes {
	t, ok := input.e.(*ast.ArrayType)
	// Sanity check
	if !ok {
		log.Warningf("ArrayExprToValExpr is not  used correctly: %T", input.e)
		return EmptyResult()
	}

	arrayLen := getArrayLen(t.Len)

	result := &TypeExprToValExprRes{}
	arrayLenToUse := g.Opts.ValTestCase.ArrayLen(arrayLen)
	exprRes := []ast.Expr{}
	for i := 0; i < arrayLenToUse; i++ {
		// Create values for array type
		recursionResult := g.TypeExprToValExpr(&RecursionInput{
			e:          t.Elt,
			varName:    input.varName,
			pkgPointer: input.pkgPointer,
			counter:    input.counter,
			identList:  input.identList,
		})
		result.Merge(recursionResult)
		exprRes = append(exprRes, recursionResult.Expr)
	}
	// Resulting array type
	res := &ast.CompositeLit{
		Type: &ast.ArrayType{
			Elt: g.CorrectTypeExpr(t.Elt, input),
			Len: g.CorrectTypeExpr(t.Len, input),
		},
		Elts: exprRes,
	}
	result.Expr = res
	return result
}

// MapExprToValExpr convert a map expression to value expression
func (g *TestCase) MapExprToValExpr(input *RecursionInput) *TypeExprToValExprRes {
	t, ok := input.e.(*ast.MapType)
	// Sanity check
	if !ok {
		log.Warningf("MapExprToValExpr is not  used correctly: %T", input.e)
		return EmptyResult()
	}

	mapLen := g.Opts.ValTestCase.MapLen()

	// Resulting map declaration
	res := &ast.CompositeLit{
		Type: &ast.MapType{
			Value: g.CorrectTypeExpr(t.Value, input),
			Key:   g.CorrectTypeExpr(t.Key, input),
		},
		Elts: []ast.Expr{},
	}
	duplCheck := NewDuplMapChecker()
	result := &TypeExprToValExprRes{}
	for i := 0; i < mapLen; i++ {
		// Create expressions for key value
		keyRecursionResult := g.TypeExprToValExpr(&RecursionInput{
			e:          t.Key,
			varName:    input.varName,
			pkgPointer: input.pkgPointer,
			counter:    input.counter,
			identList:  input.identList,
		})
		isDupl := duplCheck.IsDuplExpr(keyRecursionResult.Expr)
		if isDupl {
			continue
		}
		result.Merge(keyRecursionResult)
		// Create expressions for value value
		valRecursionInput := g.TypeExprToValExpr(&RecursionInput{
			e:          t.Value,
			varName:    input.varName,
			pkgPointer: input.pkgPointer,
			counter:    input.counter,
			identList:  input.identList,
		})
		result.Merge(valRecursionInput)

		res.Elts = append(res.Elts, &ast.KeyValueExpr{
			Key:   keyRecursionResult.Expr,
			Value: valRecursionInput.Expr,
		})
	}
	result.Expr = res
	return result
}

// StarExprToValExpr converts star expression to value expression
func (g *TestCase) StarExprToValExpr(input *RecursionInput) *TypeExprToValExprRes {
	t, ok := input.e.(*ast.StarExpr)
	// Sanity check
	if !ok {
		log.Warningf("StarExprToValExpr is not  used correctly: %T", input.e)
		return EmptyResult()
	}
	identTemp := g.Opts.IdentGen.Create(&ast.Ident{
		Name: "pointer" + cases.Title(language.English).String(input.identList.Previous().Name),
	})

	result := &TypeExprToValExprRes{}

	// Get value of t.X
	recursionResult := g.TypeExprToValExpr(&RecursionInput{
		e:          t.X,
		varName:    identTemp.Name,
		pkgPointer: input.pkgPointer,
		counter:    input.counter,
		identList:  input.identList,
	})
	result.Merge(recursionResult)
	// Create assignment of initial val
	tempAssignStmt := assignStmt(identTemp, recursionResult.Expr)
	// Create pointer from recursive expression
	// Return temp var as pointer expression
	result.Expr = &ast.UnaryExpr{
		Op: token.AND,
		X:  identTemp,
	}
	result.Statements = append(result.Statements, tempAssignStmt)
	return result
}

// StructExprToValExpr create struct expression
func (g *TestCase) StructExprToValExpr(input *RecursionInput) *TypeExprToValExprRes {
	structExpr, ok := input.e.(*ast.StructType)
	// Sanity check
	if !ok {
		log.Warningf("StructExprToValExpr is not  used correctly: %T", input.e)
		return EmptyResult()
	}
	// Create identifier for input variable name
	ident := g.CorrectTypeExpr(&ast.Ident{
		Name: input.varName,
	}, input)
	// Create result as composite literal
	res := &ast.CompositeLit{
		Type: ident,
	}
	// Store current struct in memory
	mem, ok := input.counter.StructMem[structExpr]
	if !ok {
		input.counter.StructMem[structExpr] = &ast.CompositeLit{
			Type: ident,
		}
	}
	// Detect struct cycles!
	name := input.varName
	if !g.PackageInfo.IsRoot(input.pkgPointer) {
		name += g.PackageInfo.PkgForPointer(input.pkgPointer).Name
	}
	input.counter.Structs[name]++

	// If cycle exceeds max recursion val return, we get into an infinite loop otherwise
	if input.counter.Structs[name] > g.Opts.MaxRecursion && ok {
		// Return memory
		return &TypeExprToValExprRes{
			Expr:         mem,
			Statements:   []ast.Stmt{},
			Declarations: []ast.Decl{},
		}
	}

	if structExpr.Incomplete {
		log.Warningf("Incomplete struct detected")
	}

	return g.StructFieldsToKeyValExpr(res, input)
}

// StructFieldsToKeyValExpr converts struct expression to key value expressions for initialising
// the fields of a struct
func (g *TestCase) StructFieldsToKeyValExpr(res *ast.CompositeLit, input *RecursionInput) *TypeExprToValExprRes {
	structExpr, ok := input.e.(*ast.StructType)
	// Sanity check
	if !ok {
		log.Warningf("StructFieldsToKeyValExpr is not  used correctly: %T", input.e)
		return EmptyResult()
	}
	result := &TypeExprToValExprRes{}
	elts := []ast.Expr{}
	for _, field := range structExpr.Fields.List {
		// Directly nested struct is indicated by field without names
		if len(field.Names) == 0 {
			n := g.GetUnnamedStructIdent(field.Type, input)
			if !g.PackageInfo.IsRoot(input.pkgPointer) {
				isLower := unicode.IsLower(rune(n.Name[0]))
				if isLower {
					continue
				}
			}
			recursionResult := g.TypeExprToValExpr(&RecursionInput{
				e:          field.Type,
				varName:    n.Name,
				pkgPointer: input.pkgPointer,
				counter:    input.counter,
				identList:  input.identList,
			})
			result.Merge(recursionResult)
			elts = append(elts, &ast.KeyValueExpr{
				Key:   &ast.Ident{Name: n.Name},
				Value: recursionResult.Expr,
			})
		}
		for _, n := range field.Names {
			if !g.PackageInfo.IsRoot(input.pkgPointer) {
				isLower := unicode.IsLower(rune(n.Name[0]))
				if isLower {
					continue
				}
			}
			// Detect if we are dealing with ungeneratable functions
			cantGen := g.ShouldReturnForFunc(field.Type, &RecursionInput{
				e:          field.Type,
				counter:    FreshCycleInfo(),
				pkgPointer: input.pkgPointer,
				varName:    input.varName,
				identList:  input.identList,
			})
			if cantGen {
				continue
			}

			recursionResult := g.TypeExprToValExpr(&RecursionInput{
				e:          field.Type,
				varName:    n.Name,
				pkgPointer: input.pkgPointer,
				counter:    input.counter,
				identList:  input.identList,
			})
			result.Merge(recursionResult)
			elts = append(elts, &ast.KeyValueExpr{
				Key:   &ast.Ident{Name: n.Name},
				Value: recursionResult.Expr,
			})
		}
	}
	// Fix Elts of compisite lit(resulting expression)
	res.Elts = elts
	// Add expression to resulut
	result.Expr = res
	return result
}

// nolint: unused
func (g *TestCase) printBodyStatements(funcBody *ast.BlockStmt) {
	if funcBody == nil {
		return
	}
	for _, stmt := range funcBody.List {
		fmt.Println("-----STATEMENT------")
		fmt.Printf("%+v\n", stmt)
		fmt.Println("------END OF STATEMENT-----")
	}
}

func assignStmt(variable, value ast.Expr) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{variable},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{value},
	}
}
