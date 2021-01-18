package function

func FuncFunc(x func(x int)) {
}

type FuncType func(x int)

func FuncTypeFunc(x FuncType) {}

func FuncFuncWithReturn(x func(x int) int) {
}
