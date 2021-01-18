package chanother

func ChanOther(x func(x chan (int))) {
}

type X struct {
	x chan (int)
}

func (x X) ChanOnReceiver() {
}

func ChanReceiveStructChan(x X) {
}

type Y interface {
	SomeFunc(chan (int))
}

func ChanReceiveInterfaceChan(x Y) {
}

type Z chan (X)

func ChanReceiveCustomTypeChan(x Z) {
}

func ChanPanics(x chan (int)) {
	panic("AAAH")
}

func ChanNested(x chan (chan (int))) {
	panic("AAAH")
}
