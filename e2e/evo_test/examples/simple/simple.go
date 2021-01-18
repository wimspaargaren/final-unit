package simple

func SomeFunc(x int, y string) string {
	if x < 0 {
		return y
	}

	if x > 50 && x < 60 {
		return y + "hi"
	}

	if x > 40 && x < 45 {
		return y + "hi2"
	}

	if x > 25 && x < 35 {
		return y + "hi3"
	}

	if x > 75 && x < 80 {
		return y + "hi4"
	}

	if x > 80 && x < 85 {
		return y + "hi5"
	}
	if x > 90 && x < 95 {
		return y + "hi6"
	}

	for i := 0; i < x; i++ {
		y += "x"
	}
	return y + "bye"
}

type Client interface {
	Do(req string) error
}

type Object struct {
	X      int
	Client Client
}

func (o Object) SomeOtherFunc(x []int) []int {
	if o.X > 10 {
		return x
	}
	err := o.Client.Do("some request")
	if err != nil {
		return []int{}
	}
	if len(x) == 0 {
		return []int{}
	}
	for i := 0; i < len(x); i++ {
		x[i] = x[i] + 1
	}
	return x
}
