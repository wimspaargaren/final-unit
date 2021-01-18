package chanexample

func ChanFunc(x chan (int)) int {
	counter := 0
	for {
		val, ok := <-x
		if ok {
			counter += val
		} else {
			return counter
		}
	}
}

func ChanPointerFunc(x chan (*int)) int {
	counter := 0
	for {
		val, ok := <-x
		if ok {
			counter += *val
		} else {
			return counter
		}
	}
}

type X struct {
	x int
}

type Y struct {
	x int
}

func chanPointerStructFunc(x chan (*X), y chan (*Y), amount int) {
	counter := 0
	for {
		val, ok := <-x
		if ok {
			counter += val.x
		} else {
			return
		}
	}
}
