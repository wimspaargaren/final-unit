// Package decorator provides functionality for parsing a decorater yaml
// and returns the result
package decorator

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// error definitions
var (
	ErrIncorrectDecoratorType    = fmt.Errorf("incorrect dececorator type")
	ErrParamNotFoundInFunc       = fmt.Errorf("param name not found")
	ErrReceiverNotFoundInFunc    = fmt.Errorf("receiver not found")
	ErrNoInputParamsExpected     = fmt.Errorf("expected no input parameters")
	ErrExpectedOneRetrunVal      = fmt.Errorf("expected only one return value")
	ErrDecoratorFuncNameNotFound = fmt.Errorf("decorator func name not found")
	ErrMissingFileName           = fmt.Errorf("missing file name")
)

// Deco result of a decorator file
type Deco struct {
	Files map[string]*File
}

// HasReceiverVal checks if a receiver val is specified
func (d *Deco) HasReceiverVal(fileName, funcName string) bool {
	f, ok := d.Files[fileName]
	if !ok {
		return false
	}
	function, ok := f.Funcs[funcName]
	if !ok {
		return false
	}
	return len(function.ReceiverValues) > 0
}

// GetReceiverVal retrieves receiver values for given file and func
func (d *Deco) GetReceiverVal(fileName, funcName string) []*CustomVal {
	if d.HasReceiverVal(fileName, funcName) {
		return d.Files[fileName].Funcs[funcName].ReceiverValues
	}
	return []*CustomVal{}
}

// GetVal retrieves a value
func (d *Deco) GetVal(fileName, funcName, paramName string) []*CustomVal {
	if d.HasVal(fileName, funcName, paramName) {
		return d.Files[fileName].Funcs[funcName].Params[paramName].Values
	}

	return []*CustomVal{}
}

// ShouldIgnoreFile checks if a file should be ignored
func (d *Deco) ShouldIgnoreFile(fileName string) bool {
	f, ok := d.Files[fileName]
	if !ok {
		return false
	}
	return f.Ignore
}

// ShouldIgnoreFunc checks if a function should be ignored
func (d *Deco) ShouldIgnoreFunc(fileName, funcName string) bool {
	f, ok := d.Files[fileName]
	if !ok {
		return false
	}
	function, ok := f.Funcs[funcName]
	if !ok {
		return false
	}
	return function.Ignore
}

// HasVal check if decorator has value for given file, func and param name
func (d *Deco) HasVal(fileName, funcName, paramName string) bool {
	f, ok := d.Files[fileName]
	if !ok {
		return false
	}
	function, ok := f.Funcs[funcName]
	if !ok {
		return false
	}
	param, ok := function.Params[paramName]
	if !ok {
		return false
	}
	return len(param.Values) != 0
}

// File file decorator
type File struct {
	Ignore bool
	Funcs  map[string]*Func
}

// Func func decorator
type Func struct {
	Ignore         bool
	ReceiverValues []*CustomVal
	Params         map[string]*Param
}

// Param param decorator
type Param struct {
	Values []*CustomVal
}

// CustomVal value for a given parameter or receiver
// consists of a type and call expression
type CustomVal struct {
	Type ast.Expr
	Call *ast.CallExpr
}

// Spec spec of decorator file
type Spec struct {
	CustomVals string     `yaml:"custom_vals"`
	Files      []FileSpec `yaml:"files"`
}

// FileSpec file spec of decorator file
type FileSpec struct {
	Name   string     `yaml:"name"`
	Ignore bool       `yaml:"ignore"`
	Funcs  []FuncSpec `yaml:"funcs"`
}

// FuncSpec function spec of decorator file
type FuncSpec struct {
	Name           string      `yaml:"name"`
	Ignore         bool        `yaml:"ignore"`
	ReceiverValues []string    `yaml:"receiver_values"`
	Params         []ParamSpec `yaml:"params"`
}

// ParamSpec param spec of decorator file
type ParamSpec struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values"`
}

// GetDecorators retrieves decorators if specified in given file
func GetDecorators(dir string) (*Deco, error) {
	return YamlToSpec(dir)
}

// YamlToSpec converts yaml file to result
func YamlToSpec(dir string) (*Deco, error) {
	spec, err := ParseYaml(dir)
	if err != nil {
		var pathError *os.PathError
		ok := errors.As(err, &pathError)
		if ok {
			return &Deco{
				Files: make(map[string]*File),
			}, nil
		}
		return nil, err
	}

	// Get ast custom file
	n, err := GetDecoratorFile(dir, spec.CustomVals)
	if err != nil {
		return nil, err
	}
	res, err := ConvertSpec(n, spec)
	if err != nil {
		return nil, err
	}
	err = ValidateRes(res, dir)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ValidateRes validate the resulting decorator for given dir
func ValidateRes(res *Deco, dir string) error {
	for fileName, file := range res.Files {
		n, err := ParseFile(filepath.Join(dir, fileName))
		if err != nil {
			return err
		}
		for funcName, function := range file.Funcs {
			for paramName, param := range function.Params {
				for _, v := range param.Values {
					err := ValidateParamVals(n, funcName, paramName, v.Type)
					if err != nil {
						return err
					}
				}
			}
			for _, r := range function.ReceiverValues {
				err := ValidateReceiverVals(n, funcName, r.Type)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// ValidateReceiverVals validates the type of a custom receiver value
func ValidateReceiverVals(f *ast.File, funcName string, decoType ast.Expr) error {
	info := types.Info{}
	for _, decl := range f.Decls {
		//nolint: nestif
		if t, ok := decl.(*ast.FuncDecl); ok {
			if t.Name.Name == funcName {
				if t.Recv == nil {
					return fmt.Errorf("function %s has no receiver type", funcName)
				}
				if len(t.Recv.List) != 1 {
					return fmt.Errorf("expected function %s to have only 1 receiver, but found: %d", funcName, len(t.Recv.List))
				}
				if types.Identical(info.TypeOf(t.Recv.List[0].Type), info.TypeOf(decoType)) {
					return nil
				}
			}
		}
	}
	return fmt.Errorf("%w in func %s", ErrReceiverNotFoundInFunc, funcName)
}

// ValidateParamVals validate file with given decorator type
func ValidateParamVals(f *ast.File, funcName, paramName string, decoType ast.Expr) error {
	info := types.Info{}
	for _, decl := range f.Decls {
		//nolint: nestif
		if t, ok := decl.(*ast.FuncDecl); ok {
			if t.Name.Name == funcName {
				for _, p := range t.Type.Params.List {
					for _, n := range p.Names {
						if n.Name == paramName {
							if types.Identical(info.TypeOf(p.Type), info.TypeOf(decoType)) {
								return nil
							}
							return fmt.Errorf("%w param %s got type %T, but func %s defined type %T", ErrIncorrectDecoratorType, paramName, decoType, funcName, p.Type)
						}
					}
				}
			}
		}
	}
	return fmt.Errorf("%w param %s not found in func %s", ErrParamNotFoundInFunc, paramName, funcName)
}

// ParseYaml parses a yaml for given file
func ParseYaml(dir string) (*Spec, error) {
	spec := Spec{}
	//nolint: gosec
	yamlData, err := os.ReadFile(filepath.Join(dir, "evo.yaml"))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlData, &spec)
	if err != nil {
		return nil, err
	}
	return &spec, err
}

// ConvertSpec convert spec to decorator result
func ConvertSpec(n *ast.File, spec *Spec) (*Deco, error) { //nolint: gocognit
	res := &Deco{
		Files: make(map[string]*File),
	}
	for i := 0; i < len(spec.Files); i++ {
		fileSpec := spec.Files[i]
		if fileSpec.Name == "" {
			x := []string{}
			for i := 0; i < len(fileSpec.Funcs); i++ {
				x = append(x, fileSpec.Funcs[i].Name)
			}
			return nil, fmt.Errorf("%w: %s", ErrMissingFileName, strings.Join(x, ","))
		}
		file := &File{
			Ignore: fileSpec.Ignore,
			Funcs:  make(map[string]*Func),
		}
		if fileSpec.Ignore {
			res.Files[fileSpec.Name] = file
			continue
		}
		for j := 0; j < len(fileSpec.Funcs); j++ {
			funcSpec := fileSpec.Funcs[j]
			funcDecl := &Func{
				Ignore:         funcSpec.Ignore,
				ReceiverValues: []*CustomVal{},
				Params:         make(map[string]*Param),
			}
			if fileSpec.Ignore {
				file.Funcs[funcSpec.Name] = funcDecl
				continue
			}
			for _, rVal := range funcSpec.ReceiverValues {
				x, err := FindCustomVal(n, rVal)
				if err != nil {
					return nil, err
				}
				funcDecl.ReceiverValues = append(funcDecl.ReceiverValues, x)
			}

			for k := 0; k < len(funcSpec.Params); k++ {
				paramSpec := funcSpec.Params[k]
				p := &Param{
					Values: []*CustomVal{},
				}
				for _, pVal := range paramSpec.Values {
					x, err := FindCustomVal(n, pVal)
					if err != nil {
						return nil, err
					}
					p.Values = append(p.Values, x)
				}
				funcDecl.Params[paramSpec.Name] = p
			}
			file.Funcs[funcSpec.Name] = funcDecl
		}
		res.Files[fileSpec.Name] = file
	}
	return res, nil
}

// FindCustomVal find parameter value for given spec
func FindCustomVal(f *ast.File, funcName string) (*CustomVal, error) {
	for _, decl := range f.Decls {
		//nolint: nestif
		if t, ok := decl.(*ast.FuncDecl); ok {
			if t.Name.Name == funcName {
				if len(t.Type.Params.List) != 0 {
					return nil, fmt.Errorf("%w for decorator with function name %s", ErrNoInputParamsExpected, funcName)
				}
				if len(t.Type.Results.List) != 1 {
					return nil, fmt.Errorf("%w for decorator with function name %s", ErrExpectedOneRetrunVal, funcName)
				}
				return &CustomVal{
					Type: t.Type.Results.List[0].Type,
					Call: &ast.CallExpr{
						Fun: t.Name,
					},
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("%w with name %s", ErrDecoratorFuncNameNotFound, funcName)
}

// GetDecoratorFile retrieve file with decorator spec
func GetDecoratorFile(dir, file string) (*ast.File, error) {
	if file == "" {
		return &ast.File{}, nil
	}
	return ParseFile(filepath.Join(dir, file))
}

// ParseFile creates ast file from path
func ParseFile(filePath string) (*ast.File, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return node, nil
}
