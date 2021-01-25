// Package gen provides a generator for creating assignment statements for input parameters
// of functions for every function found in a given folder
package gen

import (
	"go/ast"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/decorator"
	"github.com/wimspaargaren/final-unit/internal/ident"
	"github.com/wimspaargaren/final-unit/internal/importer"
	"github.com/wimspaargaren/final-unit/internal/testcase"
	"github.com/wimspaargaren/final-unit/pkg/values"
	"github.com/wimspaargaren/final-unit/pkg/variables"
)

// Default constants
const (
	DefaultPopulationSize   = 30
	DefaultTestCasesPerFunc = 18
	// Currently the best way to detect cycles is to count
	// the amount of times some struct is created
	DefaultAmountRecursion = 3
)

// Organism organism is a set of testcases for functions of files in a given directory
type Organism struct {
	Fitness float64
	// Create test cases for all files in a pkg
	Files []*File
}

// UpdateAssertStmts sets an os assert statements based on printed runtime result
func (o *Organism) UpdateAssertStmts(printed string, firstRun bool) {
	for _, f := range o.Files {
		for funcName, testCases := range f.TestCases {
			for i, testCase := range testCases {
				testCase.RunTimeInfo.AssertStmtsForTestCase(printed, firstRun, funcName, i)
			}
		}
	}
}

// File file contains test cases for functions of a given file
type File struct {
	Version     string
	PackageName string
	Dir         string
	FileName    string
	// funcName, list of test cases
	TestCases map[string][]*testcase.TestCase
}

// SuiteName returns the name of the test suite for this file
func (f *File) SuiteName() string {
	ext := filepath.Ext(f.FileName)
	fileWithoutExt := strings.TrimSuffix(f.FileName, ext)
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	fileParts := reg.Split(fileWithoutExt, -1)
	res := ""
	for _, part := range fileParts {
		res += strings.Title(part)
	}
	return res
}

// Options the options for the generator
type Options struct {
	MaxRecursion     int
	OrganismAmount   int
	TestCasesPerFunc int
	ValGenerator     values.IGen
	// FIXME: Remove var generator as soon as IdentGen can be used everywhere
	VarGenerator variables.IGen
	IdentGen     ident.IGen
}

// DefaultOpts creates default generator options
func DefaultOpts() *Options {
	return &Options{
		MaxRecursion:     DefaultAmountRecursion,
		OrganismAmount:   DefaultPopulationSize,
		TestCasesPerFunc: DefaultTestCasesPerFunc,
		ValGenerator:     values.NewGenerator(),
		VarGenerator:     variables.NewGenerator(),
		IdentGen:         ident.New(),
	}
}

// Generator the generator
type Generator struct {
	Version     string
	Dir         string
	PackageInfo *importer.PackageInfo
	Opts        *Options
	Deco        *decorator.Deco
}

// New creates a new generator for generating assignment statements for function parameters
func New(dir, version string, opts *Options) (*Generator, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	packageInfo, err := importer.ParseRoot(dir)
	if err != nil {
		return nil, err
	}
	// Create decorator
	deco, err := decorator.GetDecorators(dir)
	if err != nil {
		return nil, err
	}
	return &Generator{
		Version:     version,
		Dir:         dir,
		PackageInfo: packageInfo,
		Opts:        opts,
		Deco:        deco,
	}, nil
}

// GetTestCases retrieve test cases, start of recursions
func (g *Generator) GetTestCases() []*Organism {
	res := []*Organism{}
	for i := 0; i < g.Opts.OrganismAmount; i++ {
		res = append(res, g.GetNewOrganism())
	}
	return res
}

// GetNewOrganism get a single organism
func (g *Generator) GetNewOrganism() *Organism {
	org := &Organism{}
	for k, v := range g.PackageInfo.GetRootPkg() {
		_, fileName := filepath.Split(k)
		if g.Deco.ShouldIgnoreFile(fileName) {
			continue
		}
		file := &File{
			Version:     g.Version,
			Dir:         g.Dir,
			FileName:    fileName,
			PackageName: v.Name.Name,
		}
		log.Debugf("GetNewOrganism for file: %s", fileName)
		testCasesPerFunc := g.GetTestCasesForFunctionsInFile(v, g.PackageInfo.RootPointerForFileName(fileName))
		file.TestCases = testCasesPerFunc

		org.Files = append(org.Files, file)
	}
	return org
}

// GetTestCasesForFunctionsInFile convert given input ast file
// to a set of test cases per function
// test case contains a set of decl and assignment statement to generate test cases
func (g *Generator) GetTestCasesForFunctionsInFile(f *ast.File, pointer *importer.PkgResolverPointer) map[string][]*testcase.TestCase {
	// List of test cases per func name
	res := make(map[string][]*testcase.TestCase)
	for _, decl := range f.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			log.Debugf("GetTestCasesForFunctionsInFile: %s", t.Name.Name)
			// g.printBodyStatements(t.Body)

			testCases := []*testcase.TestCase{}
			for i := 0; i < g.Opts.TestCasesPerFunc; i++ {
				if t.Name.Name == "main" {
					continue
				}
				_, fileName := filepath.Split(pointer.File)
				// Decorator can specify no test generation for given functions
				if g.Deco.ShouldIgnoreFunc(fileName, t.Name.Name) {
					continue
				}
				testCase := testcase.New(t, pointer, g.PackageInfo, testcase.Options{
					ValTestCase:  g.Opts.ValGenerator,
					VarTestCase:  g.Opts.VarGenerator,
					MaxRecursion: g.Opts.MaxRecursion,
					IdentGen:     g.Opts.IdentGen,
				}, g.Deco)
				testCase.Create()
				testCases = append(testCases, testCase)
			}

			res[g.TestCasePrefix(t)+t.Name.Name] = testCases
		default:
			// Only check function declarations
			continue
		}
	}
	return res
}

// TestCasePrefix in case of receiver create prefix
// this is need to ensure test results dont override eachother in case of:
// func X() func (r T) X()
func (g *Generator) TestCasePrefix(funcDecl *ast.FuncDecl) string {
	if funcDecl.Recv == nil {
		return ""
	}
	// Sanity check, a function can only have 1 receiver
	if len(funcDecl.Recv.List) == 1 {
		return g.TypeToPrefix(funcDecl.Recv.List[0].Type)
	}
	log.Warningf("expected func receiver to have only one field")
	return g.Opts.IdentGen.Create(&ast.Ident{Name: "prefix"}).Name
}

// TypeToPrefix converts a function receiver type to a prefix
func (g *Generator) TypeToPrefix(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return g.TypeToPrefix(t.X)
	case *ast.StarExpr:
		return g.TypeToPrefix(t.X)
	default:
		log.Warningf("unexpected field receiver type found: %T", e)
		return g.Opts.IdentGen.Create(&ast.Ident{Name: "prefix"}).Name
	}
}
