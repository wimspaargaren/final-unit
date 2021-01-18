package decorator

import (
	"errors"
	"go/ast"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DecoratorTestSuite struct {
	suite.Suite
}

func (s *DecoratorTestSuite) TestSimpleDeco() {
	res, err := GetDecorators("testdata/simple")
	s.Require().NoError(err)

	file, ok := res.Files["password.go"]
	s.Require().True(ok)
	function, ok := file.Funcs["ComparePassword"]
	s.Require().True(ok)
	s.Require().Equal(1, len(function.ReceiverValues))
	s.True(res.HasReceiverVal("password.go", "ComparePassword"))
	paramDef, ok := function.Params["password"]
	s.Require().True(ok)
	s.Require().Equal(2, len(paramDef.Values))
	t, ok := paramDef.Values[0].Type.(*ast.Ident)
	s.Require().True(ok)
	s.Equal("string", t.Name)
	values := res.GetVal("password.go", "ComparePassword", "password")
	s.Equal(2, len(values))
	values = res.GetVal("password.go", "ComparePassword", "x")
	s.Equal(0, len(values))
	values = res.GetVal("password.go", "x", "x")
	s.Equal(0, len(values))
	values = res.GetVal("x", "x", "x")
	s.Equal(0, len(values))
}

func (s *DecoratorTestSuite) TestIgnoreFile() {
	res, err := GetDecorators("testdata/ignorefile")
	s.Require().NoError(err)
	s.NotNil(res)
	file, ok := res.Files["ignore.go"]
	s.Require().True(ok)
	s.True(file.Ignore)
	s.True(res.ShouldIgnoreFile("ignore.go"))
	s.False(res.ShouldIgnoreFile("dont_ignore.go"))
}

func (s *DecoratorTestSuite) TestIgnoreFunc() {
	res, err := GetDecorators("testdata/ignorefunc")
	s.Require().NoError(err)
	s.NotNil(res)
	file, ok := res.Files["ignore.go"]
	s.Require().True(ok)
	s.False(file.Ignore)
	function, ok := file.Funcs["IgnoreMe"]
	s.Require().True(ok)
	s.True(function.Ignore)
	s.True(res.ShouldIgnoreFunc("ignore.go", "IgnoreMe"))
	s.False(res.ShouldIgnoreFunc("ignore.go", "DontIgnoreMe"))
	s.False(res.ShouldIgnoreFunc("x.go", "DontIgnoreMe"))
}

func (s *DecoratorTestSuite) TestNoDeco() {
	res, err := GetDecorators("testdata/nodeco")
	s.Require().NoError(err)
	s.NotNil(res)
	s.NotNil(res.Files)
	values := res.GetVal("password.go", "ComparePassword", "x")
	s.Equal(0, len(values))
}

func (s *DecoratorTestSuite) TestIncorrect() {
	_, err := GetDecorators("testdata/incorrect")
	s.Require().Error(err)
	var pathError *os.PathError
	ok := errors.As(err, &pathError)
	s.Require().True(ok)
}

func (s *DecoratorTestSuite) TestAbsolutePath() {
	x, err := filepath.Abs("testdata/simple")
	s.Require().NoError(err)
	_, err = GetDecorators(x)
	s.Require().NoError(err)
}

func (s *DecoratorTestSuite) TestIncorrectSpec() {
	_, err := GetDecorators("testdata/incorrectspec")
	s.Require().Error(err)
}

func TestDecoratorTestSuite(t *testing.T) {
	suite.Run(t, new(DecoratorTestSuite))
}
