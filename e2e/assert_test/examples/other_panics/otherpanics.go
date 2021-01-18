package otherpanics

type X struct{}

func (x X) DefTwo() int {
	panic("AAH")
	return 43
}
