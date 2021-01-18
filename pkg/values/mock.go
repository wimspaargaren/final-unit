package values

import "github.com/stretchr/testify/mock"

// GenMock mocks the basic lit generator
type GenMock struct {
	mock.Mock
}

var _ IGen = &GenMock{}

// Error returns if error value
func (g *GenMock) Error() bool {
	args := g.MethodCalled("Error")
	return args.Get(0).(bool)
}

// DecoratorVal returns if decorator value
func (g *GenMock) DecoratorVal() bool {
	args := g.MethodCalled("DecoratorVal")
	return args.Get(0).(bool)
}

// DecoratorIndex returns decorator index for given length
func (g *GenMock) DecoratorIndex(length int) int {
	args := g.MethodCalled("DecoratorIndex", length)
	return args.Get(0).(int)
}

// Type returns type value
func (g *GenMock) Type() string {
	args := g.MethodCalled("Type")
	return args.Get(0).(string)
}

// Int returns int value
func (g *GenMock) Int() string {
	args := g.MethodCalled("Int")
	return args.Get(0).(string)
}

// Int8 returns int8 value
func (g *GenMock) Int8() string {
	args := g.MethodCalled("Int8")
	return args.Get(0).(string)
}

// Int16 returns int16 value
func (g *GenMock) Int16() string {
	args := g.MethodCalled("Int16")
	return args.Get(0).(string)
}

// Int32 returns int32 value
func (g *GenMock) Int32() string {
	args := g.MethodCalled("Int32")
	return args.Get(0).(string)
}

// Int64 returns int64 value
func (g *GenMock) Int64() string {
	args := g.MethodCalled("Int64")
	return args.Get(0).(string)
}

// UInt returns uint value
func (g *GenMock) UInt() string {
	args := g.MethodCalled("UInt")
	return args.Get(0).(string)
}

// UInt8 returns uint8 value
func (g *GenMock) UInt8() string {
	args := g.MethodCalled("UInt8")
	return args.Get(0).(string)
}

// UInt16 returns uint16 value
func (g *GenMock) UInt16() string {
	args := g.MethodCalled("UInt16")
	return args.Get(0).(string)
}

// UInt32 returns uint32 value
func (g *GenMock) UInt32() string {
	args := g.MethodCalled("UInt32")
	return args.Get(0).(string)
}

// UInt64 returns uint64 value
func (g *GenMock) UInt64() string {
	args := g.MethodCalled("UInt64")
	return args.Get(0).(string)
}

// UIntPtr returns uintptr value
func (g *GenMock) UIntPtr() string {
	args := g.MethodCalled("UIntPtr")
	return args.Get(0).(string)
}

// Bool returns bool value
func (g *GenMock) Bool() string {
	args := g.MethodCalled("Bool")
	return args.Get(0).(string)
}

// String returns string value
func (g *GenMock) String() string {
	args := g.MethodCalled("String")
	return args.Get(0).(string)
}

// Float64 returns float64 value
func (g *GenMock) Float64() string {
	args := g.MethodCalled("Float64")
	return args.Get(0).(string)
}

// Float32 returns float32 value
func (g *GenMock) Float32() string {
	args := g.MethodCalled("Float32")
	return args.Get(0).(string)
}

// Complex64 returns complex64 value
func (g *GenMock) Complex64() string {
	args := g.MethodCalled("Complex64")
	return args.Get(0).(string)
}

// Complex128 returns complex128 value
func (g *GenMock) Complex128() string {
	args := g.MethodCalled("Complex128")
	return args.Get(0).(string)
}

// Byte returns byte value
func (g *GenMock) Byte() string {
	args := g.MethodCalled("Byte")
	return args.Get(0).(string)
}

// Rune returns rune value
func (g *GenMock) Rune() string {
	args := g.MethodCalled("Rune")
	return args.Get(0).(string)
}

// ArrayLen returns length of arrays
func (g *GenMock) ArrayLen(maxLen int) int {
	args := g.MethodCalled("ArrayLen", maxLen)
	return args.Get(0).(int)
}

// MapLen returns length of a map
func (g *GenMock) MapLen() int {
	args := g.MethodCalled("MapLen")
	return args.Get(0).(int)
}
