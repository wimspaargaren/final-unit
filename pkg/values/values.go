// Package values provides an interface for generating values
// nolint: gosec
package values

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/wimspaargaren/final-unit/pkg/chance"
)

// IGen the value generator interface
type IGen interface {
	Bool() string
	String() string
	Type() string

	Int() string
	Int8() string
	Int16() string
	Int32() string
	Int64() string

	UInt() string
	UInt8() string
	UInt16() string
	UInt32() string
	UInt64() string
	UIntPtr() string

	Byte() string
	Rune() string

	Float32() string
	Float64() string

	Complex64() string
	Complex128() string

	Error() bool
	DecoratorVal() bool
	DecoratorIndex(length int) int

	ArrayLen(maxLen int) int
	MapLen() int
}

// Gen IGen implementation
type Gen struct{}

// NewGenerator creates a new generator
func NewGenerator() IGen {
	rand.Seed(time.Now().UnixNano())
	return &Gen{}
}

// Int Generates an int value
func (g *Gen) Int() string {
	return IntVal()
}

// Type retrieves random basic lit type
func (g *Gen) Type() string {
	types := []string{
		"int",
		"bool",
		"string",
		"float32",
		"float64",
		"byte",
		"rune",
		"uintptr",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"int8",
		"int16",
		"int32",
		"int64",
		"complex64",
		"complex128",
	}

	return types[randomdata.Number(0, len(types))]
}

// Int8 Generates an int8 value
func (g *Gen) Int8() string {
	return IntVal()
}

// Int16 Generates an int16 value
func (g *Gen) Int16() string {
	return IntVal()
}

// Int32 Generates an int32 value
func (g *Gen) Int32() string {
	return IntVal()
}

// Int64 Generates an int64 value
func (g *Gen) Int64() string {
	return IntVal()
}

// UInt Generates an uint value
func (g *Gen) UInt() string {
	return UIntVal()
}

// UInt8 Generates an uint8 value
func (g *Gen) UInt8() string {
	return UIntVal()
}

// UInt16 Generates an uint16 value
func (g *Gen) UInt16() string {
	return UIntVal()
}

// UInt32 Generates an uint32 value
func (g *Gen) UInt32() string {
	return UIntVal()
}

// UInt64 Generates an uint64 value
func (g *Gen) UInt64() string {
	return UIntVal()
}

// UIntPtr Generates an uintptr value
func (g *Gen) UIntPtr() string {
	return IntVal()
}

// Bool Generates an bool value
func (g *Gen) Bool() string {
	val := randomdata.Boolean()
	return strconv.FormatBool(val)
}

// String Generates an bool value
func (g *Gen) String() string {
	val := randomdata.SillyName()
	return fmt.Sprintf(`"%s"`, val)
}

// Float64 Generates an float64 value
func (g *Gen) Float64() string {
	return FloatVal()
}

// Float32 Generates an float32 value
func (g *Gen) Float32() string {
	return FloatVal()
}

// Complex64 Generates an complex64 value
func (g *Gen) Complex64() string {
	return "42"
}

// Complex128 Generates an complex128 value
func (g *Gen) Complex128() string {
	return IntVal()
}

// Byte Generates an byte value
func (g *Gen) Byte() string {
	return UIntVal()
}

// Rune Generates an rune value
func (g *Gen) Rune() string {
	return IntVal()
}

// Error Indicates if an error should be returned or nil
func (g *Gen) Error() bool {
	val := randomdata.Boolean()
	return val
}

// DecoratorVal Indicates if a value should be used in the decorator spec
func (g *Gen) DecoratorVal() bool {
	const decoratorChance = 80
	return chance.IsChance(decoratorChance)
}

// DecoratorIndex returns random index for array length of decorators
func (g *Gen) DecoratorIndex(length int) int {
	return chance.GetIndex(length)
}

const (
	maxArrayLen     = 10
	changeVal       = 100
	notMaxLenChance = 10
)

// ArrayLen creates the length of an array
func (g *Gen) ArrayLen(maxLen int) int {
	if maxLen == 0 {
		return 0
	}
	change := rand.Intn(changeVal)
	if change < 1 {
		return 0
	}
	if maxLen != -1 {
		if change > notMaxLenChance {
			return maxLen
		}
		return rand.Intn(maxLen)
	}
	return rand.Intn(maxArrayLen)
}

// MapLen creates the length of a map
func (g *Gen) MapLen() int {
	return rand.Intn(maxArrayLen)
}

// Helpers

// IntVal create random int value and converts it to string
func IntVal() string {
	val := randomdata.Number(-100, 100)
	return strconv.Itoa(val)
}

// UIntVal create random uint value and converts it to string
func UIntVal() string {
	val := randomdata.Number(0, 100)
	return strconv.Itoa(val)
}

// FloatVal create random float value and converts it to string
func FloatVal() string {
	const maxAmountDecimalPoints = 5
	decimalPoints := randomdata.Number(maxAmountDecimalPoints)
	val := randomdata.Decimal(-100, 100, decimalPoints)
	return fmt.Sprintf("%f", val)
}
