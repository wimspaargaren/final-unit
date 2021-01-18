package ignore

func IgnoreChan() chan (int) {
	x := make(chan int)
	return x
}

func IgnoreFunc() func(x int) {
	return func(x int) {
	}
}

func IgnoreInterfaceEmpty() interface{} {
	return nil
}

type Normal interface {
	X()
}

func IgnoreInterface() Normal {
	return nil
}

func FuncReturnFunc() func() {
	return func() {}
}
