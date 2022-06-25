// Package importer provides a function to parse a directory and returns the package information
package importer

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// error definitions
var (
	ErrAtleastOnePackageRequired = fmt.Errorf("expected atleast one package")
	ErrIncorrectPackageAmount    = fmt.Errorf("unexpected amount of packages")
	ErrUnableToFindIdentifier    = fmt.Errorf("unable to find identifier")
)

// PkgResolverPointer pointer of a the package resolver
// used to identify where to look for specific values
type PkgResolverPointer struct {
	Dir  string
	Pkg  string
	File string
}

// PackageInfo package info contains information about a package defined
// in a directory
type PackageInfo struct {
	RootDir string
	RootPkg string
	// dir is unique

	// map[dir][packagename]package: []Files
	PkgInfo map[string]map[string]*ast.Package
}

// ParseRoot parse a root directory
func ParseRoot(dir string) (*PackageInfo, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, FileFilter, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	res := make(map[string]map[string]*ast.Package)
	// FIXME checks for parsing root
	pkg := ""
	for k := range pkgs {
		pkg = k
	}
	res[dir] = pkgs
	ret := &PackageInfo{
		RootDir: dir,
		PkgInfo: res,
		RootPkg: pkg,
	}
	return ret, nil
}

// FileFilter filter files which should not be included in this case all files which end with _test.go
func FileFilter(fileInfo os.FileInfo) bool {
	return !strings.HasSuffix(fileInfo.Name(), "_test.go")
}

// GetImportSpecForIdentifierAndFile Find an import spec for given identifier in given file
func GetImportSpecForIdentifierAndFile(identifier string, file *ast.File) (*ast.ImportSpec, error) {
	for _, i := range file.Imports {
		if i.Name != nil {
			if i.Name.Name == identifier {
				return i, nil
			}
		}
	}
	// we assume idents can not equal pkg
	for _, i := range file.Imports {
		if importIdentifier(i) == identifier {
			return i, nil
		}
	}
	return TryToFindIdentifier(identifier, file)
}

// TryToFindIdentifier in case this is hit, we got unresolved imports
// an example could be "somepkg/v2"
func TryToFindIdentifier(identifier string, file *ast.File) (*ast.ImportSpec, error) {
	for _, i := range file.Imports {
		importPath := strings.ReplaceAll(i.Path.Value, `"`, "")
		ctx := build.Default
		if importPath == "C" {
			continue
		}
		resPkg, err := ctx.Import(importPath, ".", build.FindOnly)
		if err != nil {
			return nil, err
		}
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, resPkg.Dir, FileFilter, parser.AllErrors)
		if err != nil {
			return nil, err
		}
		if len(pkgs) != 1 {
			log.Warningf("need test case for multi pkgs in unresolvable import")
		}
		for n := range pkgs {
			if n == identifier {
				return i, nil
			}
		}
	}
	return nil, ErrUnableToFindIdentifier
}

// ImportPathToFilePath converts an import path to file path
func ImportPathToFilePath(i *ast.ImportSpec) (string, error) {
	importPath := strings.ReplaceAll(i.Path.Value, `"`, "")
	ctx := build.Default

	if importPath == "C" {
		return "", nil
	}
	pkg, err := ctx.Import(importPath, ".", build.FindOnly)
	if err != nil {
		return "", err
	}
	return pkg.Dir, nil
}

// importIdentifier converts a full import path to identifier value
func importIdentifier(i *ast.ImportSpec) string {
	importPathValue := strings.Trim(i.Path.Value, `"`)
	parts := strings.Split(importPathValue, "/")
	if len(parts) == 1 {
		return importPathValue
	}
	return parts[len(parts)-1]
}

// RootPointerForFileName get root pointer for given filename
func (p *PackageInfo) RootPointerForFileName(fileName string) *PkgResolverPointer {
	return &PkgResolverPointer{
		Dir:  p.RootDir,
		Pkg:  p.RootPkg,
		File: filepath.Join(p.RootDir, fileName),
	}
}

// FileForPointer retrieve ast file for pointer
func (p *PackageInfo) FileForPointer(pointer *PkgResolverPointer) *ast.File {
	// FIME built in sanity checks
	x, ok := p.PkgInfo[pointer.Dir]
	if !ok {
		log.Warningf("directory: %s not found", pointer.Dir)
		return nil
	}

	pkg, ok := x[pointer.Pkg]
	if !ok {
		log.Warningf("pkg: %s not found", pointer.Pkg)
		return nil
	}
	f, ok := pkg.Files[pointer.File]
	if !ok {
		log.Warningf("file: %s not found", pointer.File)
		return nil
	}
	return f
}

// PkgForPointer retrieve ast package for pointer
func (p *PackageInfo) PkgForPointer(pointer *PkgResolverPointer) *ast.Package {
	// FIME built in sanity checks
	return p.PkgInfo[pointer.Dir][pointer.Pkg]
}

// PkgsForDir resolve directory
func (p *PackageInfo) PkgsForDir(dir string) map[string]*ast.Package {
	pkgs, ok := p.PkgInfo[dir]
	if ok {
		return pkgs
	}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, FileFilter, parser.AllErrors)
	if err != nil {
		log.WithError(err).Errorf("unable to convert importSpec to file path: %s", dir)
		return nil
	}
	p.PkgInfo[dir] = pkgs
	return pkgs
}

// GetRootPkg retrieve the root package
func (p *PackageInfo) GetRootPkg() map[string]*ast.File {
	// FIXME add sanity checks
	pkgs := p.PkgInfo[p.RootDir]
	for _, v := range pkgs {
		return v.Files
	}
	return nil
}

// IsRoot check if pointer is in root
// some decisions need to be based on this
func (p *PackageInfo) IsRoot(pointer *PkgResolverPointer) bool {
	return p.RootDir == pointer.Dir
}

// FindImport finds a imported object on selector and identifier
func (p *PackageInfo) FindImport(pointer *PkgResolverPointer, selector, identifier string) (bool, ast.Expr, *PkgResolverPointer) {
	f := p.FileForPointer(pointer)
	if f == nil {
		log.Warningf("file not found for current pointer")
		return false, nil, pointer
	}
	importSpec, err := GetImportSpecForIdentifierAndFile(selector, f)
	if err != nil {
		log.WithError(err).Errorf("unable to get import spec for identifier and file: %s", selector)
		return false, nil, pointer
	}
	dir, err := ImportPathToFilePath(importSpec)
	if err != nil {
		log.WithError(err).Errorf("unable to convert importSpec to file path: %s", importSpec.Path.Value)
		return false, nil, pointer
	}
	pkgs := p.PkgsForDir(dir)

	for packageName, pkg := range pkgs {
		for fileName := range pkg.Files {
			found, expr, pointer := p.FindInCurrent(&PkgResolverPointer{
				Dir:  dir,
				Pkg:  packageName,
				File: fileName,
			}, identifier)
			if found {
				return found, expr, pointer
			}
		}
	}

	log.Warningf("was not able to find identifier in import. selector: %s, identifier: %s, pointer; %v", selector, identifier, pointer)
	return false, nil, pointer
}

// FindInCurrent tries to find an indentifier in current package
func (p *PackageInfo) FindInCurrent(pointer *PkgResolverPointer, identifier string) (bool, ast.Expr, *PkgResolverPointer) {
	pkg := p.PkgForPointer(pointer)
	// Iterate files in root
	for fileName, f := range pkg.Files {
		// Iterate declartions in file
		for _, d := range f.Decls {
			// Check if declarations is of type GenDecl
			if genDecl, ok := d.(*ast.GenDecl); ok {
				// Iterater Specs
				for _, s := range genDecl.Specs {
					// If spec is type spec & identifier matches
					// return found
					if typeSpec, ok := s.(*ast.TypeSpec); ok {
						if typeSpec.Name.Name == identifier {
							// types declared in other packages should always be retruned as
							// identifiers with a filled object, execept for interfaces
							// which can be just returned as typeSpec.Type
							switch typeSpec.Type.(type) {
							case *ast.InterfaceType:
								return true, typeSpec.Type, &PkgResolverPointer{Dir: pointer.Dir, Pkg: pointer.Pkg, File: fileName}
							default:
								return true, &ast.Ident{
									Name: typeSpec.Name.Name,
									Obj: &ast.Object{
										Decl: typeSpec,
									},
								}, &PkgResolverPointer{Dir: pointer.Dir, Pkg: pointer.Pkg, File: fileName}
							}
						}
					}
				}
			}
		}
	}
	return false, nil, pointer
}
