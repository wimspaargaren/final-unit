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
