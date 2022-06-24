package testcase

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strconv"

	"github.com/asherascout/final-unit/internal/identlist"
	"github.com/asherascout/final-unit/internal/importer"
	log "github.com/sirupsen/logrus"
)

// Writer io writer implementation for gathering element formatted strings
type Writer struct {
	Res []byte
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.Res = append(w.Res, p...)
	return 0, nil
}

// GetString convert byte array to string
func (w *Writer) GetString() string {
	return string(w.Res)
}

func newWriter() *Writer {
	return &Writer{
		Res: []byte{},
	}
}

// PrettyPrintElement pretty prints an AST element
func PrettyPrintElement(e interface{}) (string, error) {
	writer := newWriter()
	fset := token.NewFileSet()
	err := printer.Fprint(writer, fset, e)
	if err != nil {
		return "", err
	}
	return writer.GetString(), nil
}

// MustPrettyPrintElement pretty prints an AST element or panics
func MustPrettyPrintElement(e interface{}) string {
	if e == nil {
		return ""
	}
	writer := newWriter()
	fset := token.NewFileSet()
	err := printer.Fprint(writer, fset, e)
	if err != nil {
		panic(err)
	}
	return writer.GetString()
}

func getArrayLen(expr ast.Expr) int {
	if expr == nil {
		return -1
	}
	switch t := expr.(type) {
	case *ast.BasicLit:
		len, err := strconv.Atoi(t.Value)
		if err != nil {
			log.WithError(err).Errorf("unable to convert basic lit val to integer")
		}
		return len
	case *ast.Ident:
		if vSpec, ok := t.Obj.Decl.(*ast.ValueSpec); ok {
			if len(vSpec.Values) != 1 {
				log.Warningf("unexpected amount of vspec values")
			}
			for _, e := range vSpec.Values {
				return getArrayLen(e)
			}
		} else {
			log.Warningf("WUT: %T", t)
		}
	default:
		log.Warningf("unknown array len expression: %T", t)
		return -1
	}
	return -1
}

// NewRecursionInput creates new recursion input
func NewRecursionInput(valType ast.Expr, identifierName string, pointer *importer.PkgResolverPointer, ident *ast.Ident) *RecursionInput {
	return &RecursionInput{
		e:          valType,
		varName:    identifierName,
		pkgPointer: pointer,
		counter:    FreshCycleInfo(),
		identList:  identlist.New(ident),
	}
}

// NewPrintRecursionInput creates new recursion input
func NewPrintRecursionInput(valType ast.Expr, identifierName string, pointer *importer.PkgResolverPointer) *PrintRecursionInput {
	return &PrintRecursionInput{
		e:          valType,
		varName:    identifierName,
		pkgPointer: pointer,
		counter:    FreshCycleInfo(),
	}
}
