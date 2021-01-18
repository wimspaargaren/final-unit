package structpackage

type Simple struct {
	Var int
}

func StructFuncSimple(x Simple) {
}

type Composite struct {
	VarInt    int
	VarString *string
}

func StructFuncComposite(x Composite) {
}

type CompositeOther struct {
	x, y int
}

func StructFuncCompositeOther(x CompositeOther) {
}

type Nested struct {
	X Simple
}

func StructFuncNested(x Nested) {
}

type NestedNested struct {
	Nested
	x int
}

func StructFuncNestedNested(x NestedNested) {
}

type NestedInt struct {
	int
	x string
}

func StructFuncNestedInt(x NestedInt) {
}

type CustomType func(x int) string

func StructFuncNestedCustomType(x NestedCustomType) {
}

type NestedCustomType struct {
	CustomType
	x string
}
