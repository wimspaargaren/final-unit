package funconstruct

type Simple struct {
	X int
}

func (s Simple) SimpleFuncOnStruct(x int) {}

type CustomType int

func (c CustomType) FuncOnCustomType(x int) {}
