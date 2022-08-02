// Package gen provides a generator for creating assignment statements for input parameters
// of functions for every function found in a given folder
package gen

import (
	"go/ast"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/wimspaargaren/final-unit/internal/decorator"
	"github.com/wimspaargaren/final-unit/internal/ident"
	"github.com/wimspaargaren/final-unit/internal/importer"
	"github.com/wimspaargaren/final-unit/internal/testcase"
	"github.com/wimspaargaren/final-unit/pkg/values"
	"github.com/wimspaargaren/final-unit/pkg/variables"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Organism organism is a set of testcases for functions of files in a given directory
type Organism struct {
	Fitness float64
	// Create test cases for all files in a pkg
	Files []*File
}

// NewOrganism creates a new organism
func NewOrganism(files []*File) *Organism {
	sort.Slice(files, func(i, j int) bool {
		return files[i].FileName < files[j].FileName
	})
	return &Organism{Files: files}
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
	PackageName string
	FileName    string
	// funcName, list of test cases
	TestCases   map[string][]*testcase.TestCase
	PackageInfo *importer.PackageInfo
	IdentGen    ident.IGen
	Opts        *Options
	Deco        *decorator.Deco
}

// NewFile creates a new file object
func NewFile(pathName string, pkgInfo *importer.PackageInfo, opts *Options, deco *decorator.Deco) *File {
	astFile, ok := pkgInfo.GetRootPkg()[pathName]
	if !ok {
		return nil
	}

	file := &File{
		FileName:    pathName,
		PackageName: pkgInfo.RootPkg,
		PackageInfo: pkgInfo,
		Opts:        opts,
		Deco:        deco,
		IdentGen:    ident.New(),
	}
	file.TestCases = file.GetTestCasesForFunctionsInFile(pathName, astFile)
	return file
}

// SuiteName returns the name of the test suite for this file
func (f *File) SuiteName() string {
	_, fileName := filepath.Split(f.FileName)
	ext := filepath.Ext(fileName)
	fileWithoutExt := strings.TrimSuffix(fileName, ext)
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	fileParts := reg.Split(fileWithoutExt, -1)
	res := ""
	for _, part := range fileParts {
		res += cases.Title(language.English).String(part)
	}
	return res
}

// Options the options for the generator
type Options struct {
	MaxRecursion     int
	OrganismAmount   int
	TestCasesPerFunc int
}

// Generator the generator
type Generator struct {
	Dir         string
	PackageInfo *importer.PackageInfo
	Opts        *Options
	Deco        *decorator.Deco
}

// New creates a new generator for generating assignment statements for function parameters
func New(dir string, opts *Options) (*Generator, error) {
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
	var files []*File
	for fileName := range g.PackageInfo.GetRootPkg() {
		if g.Deco.ShouldIgnoreFile(fileName) {
			continue
		}
		file := NewFile(fileName, g.PackageInfo, g.Opts, g.Deco)
		log.Debugf("GetNewOrganism for file: %s", fileName)
		files = append(files, file)
	}
	return NewOrganism(files)
}

// GetTestCasesForFunctionsInFile convert given input ast file
// to a set of test cases per function
// test case contains a set of decl and assignment statement to generate test cases
func (f *File) GetTestCasesForFunctionsInFile(path string, astFile *ast.File) map[string][]*testcase.TestCase {
	// List of test cases per func name
	res := make(map[string][]*testcase.TestCase)
	for _, decl := range astFile.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			log.Debugf("GetTestCasesForFunctionsInFile: %s", t.Name.Name)

			testCases := []*testcase.TestCase{}
			for i := 0; i < f.Opts.TestCasesPerFunc; i++ {
				if t.Name.Name == "main" {
					continue
				}
				// Decorator can specify no test generation for given functions
				if f.Deco.ShouldIgnoreFunc(path, t.Name.Name) {
					continue
				}
				pointer := &importer.PkgResolverPointer{
					Dir:  f.PackageInfo.RootDir,
					Pkg:  f.PackageInfo.RootPkg,
					File: path,
				}
				testCase := testcase.New(t, pointer, f.PackageInfo, testcase.Options{
					ValTestCase:  values.NewGenerator(),
					VarTestCase:  variables.NewGenerator(),
					MaxRecursion: f.Opts.MaxRecursion,
					IdentGen:     f.IdentGen,
				}, f.Deco)
				testCase.Create()
				testCases = append(testCases, testCase)
			}

			res[f.TestCasePrefix(t)+t.Name.Name] = testCases
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
func (f *File) TestCasePrefix(funcDecl *ast.FuncDecl) string {
	if funcDecl.Recv == nil {
		return ""
	}
	// Sanity check, a function can only have 1 receiver
	if len(funcDecl.Recv.List) == 1 {
		return f.TypeToPrefix(funcDecl.Recv.List[0].Type)
	}
	log.Warningf("expected func receiver to have only one field")
	return f.IdentGen.Create(&ast.Ident{Name: "prefix"}).Name
}

// TypeToPrefix converts a function receiver type to a prefix
func (f *File) TypeToPrefix(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return f.TypeToPrefix(t.X)
	case *ast.StarExpr:
		return f.TypeToPrefix(t.X)
	default:
		log.Warningf("unexpected field receiver type found: %T", e)
		return f.IdentGen.Create(&ast.Ident{Name: "prefix"}).Name
	}
}
