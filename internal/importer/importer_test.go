package importer

import (
	"go/ast"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	dir  = "examples/example_simple"
	pkg  = "simple"
	file = "examples/example_simple/simple.go"
)

type ImporterTestSuite struct {
	suite.Suite
}

func (s *ImporterTestSuite) TestV2Import() {
	dir := "examples/example_v2"
	// Package info needed in recursion
	pkg := "v2"
	file := "examples/example_v2/v2.go"
	// On start gen
	pointer := PkgResolverPointer{
		Dir:  dir,
		Pkg:  pkg,
		File: file,
	}
	res, err := ParseRoot(pointer.Dir)
	s.Require().NoError(err)

	// Dir info needed in recursion
	x := res.PkgInfo[res.RootDir]
	s.Equal(1, len(x))

	rootPkg := x[pointer.Pkg]
	s.NotNil(rootPkg)

	// File info needed in recursion
	simpleGoFile := rootPkg.Files[pointer.File]
	s.NotNil(simpleGoFile)
	s.Require().NoError(err)
	s.resolveImport("resty", "/resty/v2", simpleGoFile)
}

func (s *ImporterTestSuite) TestResolveImportsForPointer() {
	// Package info needed in recursion
	// On start gen
	pointer := PkgResolverPointer{
		Dir:  dir,
		Pkg:  pkg,
		File: file,
	}
	res, err := ParseRoot(pointer.Dir)
	s.Require().NoError(err)

	// MUUCHO IMPORTANTE:
	// Need dir name
	// Need pkg name
	// Need file name

	// Dir info needed in recursion
	x := res.PkgInfo[res.RootDir]
	s.Equal(1, len(x))

	rootPkg := x[pointer.Pkg]
	s.NotNil(rootPkg)

	// File info needed in recursion
	simpleGoFile := rootPkg.Files[pointer.File]
	s.NotNil(simpleGoFile)
	s.resolveImport("fmt", "/fmt", simpleGoFile)
	s.resolveImport("foo", "github.com/gofrs/uuid", simpleGoFile)
	s.resolveImport("somepkg", "/somepkg", simpleGoFile)
}

func (s *ImporterTestSuite) TestFindInCurrent() {
	// On start gen
	pointer := &PkgResolverPointer{
		Dir:  dir,
		Pkg:  pkg,
		File: file,
	}
	res, err := ParseRoot(pointer.Dir)
	s.Require().NoError(err)
	s.True(res.IsRoot(pointer))
	found, expr, newPointer := res.FindInCurrent(pointer, "StructWeAreLookingFor")
	s.True(found)
	s.True(res.IsRoot(newPointer))
	x, ok := expr.(*ast.Ident)
	s.Require().True(ok)
	s.Equal("StructWeAreLookingFor", x.Name)
	s.Equal(pointer.Dir, newPointer.Dir)
	s.Equal(pointer.Pkg, newPointer.Pkg)
	s.Equal("examples/example_simple/simple_extra_file.go", newPointer.File)
}

func (s *ImporterTestSuite) TestFindInImport() {
	// On start gen
	pointer := &PkgResolverPointer{
		Dir:  dir,
		Pkg:  pkg,
		File: file,
	}
	res, err := ParseRoot(pointer.Dir)
	s.Require().NoError(err)
	s.True(res.IsRoot(pointer))

	identifier := "SomeStruct"
	selector := "somepkg"

	found, expr, newPointer := res.FindImport(pointer, selector, identifier)
	s.True(found)
	s.False(res.IsRoot(newPointer))
	x, ok := expr.(*ast.Ident)
	s.Require().True(ok)
	s.Equal("SomeStruct", x.Name)
	s.True(strings.HasSuffix(newPointer.Dir, "internal/importer/examples/example_simple/pkg/somepkg"))
	s.Equal(selector, newPointer.Pkg)
	s.True(strings.HasSuffix(newPointer.File, "internal/importer/examples/example_simple/pkg/somepkg/somepkg.go"))

	found, expr, newPointer = res.FindInCurrent(newPointer, "SomeOtherStruct")
	s.True(found)
	x, ok = expr.(*ast.Ident)
	s.Require().True(ok)
	s.Equal("SomeOtherStruct", x.Name)
	s.True(strings.HasSuffix(newPointer.Dir, "internal/importer/examples/example_simple/pkg/somepkg"))

	s.Equal(selector, newPointer.Pkg)
	s.True(strings.HasSuffix(newPointer.File, "internal/importer/examples/example_simple/pkg/somepkg/somepkg_addon.go"))
}

func (s *ImporterTestSuite) TestOtherExample() {
	dir := "examples/example_other"
	// Package info needed in recursion
	pkg := "other"
	file := "examples/example_other/other.go"
	// On start gen
	pointer := &PkgResolverPointer{
		Dir:  dir,
		Pkg:  pkg,
		File: file,
	}
	res, err := ParseRoot(pointer.Dir)
	s.Require().NoError(err)
	s.True(res.IsRoot(pointer))

	identifier := "Time"
	selector := "time"

	found, expr, newPointer := res.FindImport(pointer, selector, identifier)
	s.True(found)
	s.False(res.IsRoot(newPointer))
	x, ok := expr.(*ast.Ident)
	s.Require().True(ok)
	s.Equal(identifier, x.Name)
	s.Contains(newPointer.Dir, "/time")
	s.Equal(selector, newPointer.Pkg)
	s.Contains(newPointer.File, "/time/time.go")
}

func (s *ImporterTestSuite) resolveImport(identifier, expectedPath string, file *ast.File) {
	importSpec, err := GetImportSpecForIdentifierAndFile(identifier, file)
	s.Require().NoError(err)
	path, err := ImportPathToFilePath(importSpec)
	s.NoError(err)
	s.Contains(path, expectedPath)
}

func TestImporterTestSuite(t *testing.T) {
	suite.Run(t, new(ImporterTestSuite))
}
