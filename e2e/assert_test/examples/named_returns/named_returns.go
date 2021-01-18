package namedreturns

func NameReturn() (hi, bye int) {
	hi = 3
	bye = 4
	return
}

func MultiNamedReturn() (hi, bye int, x, y, z string) {
	hi = 3
	bye = 4
	x = "x"
	y = "y"
	return
}

func WeirdNamedReturn() (int, x, y, z string) {
	int = "hi"
	x = "x"
	y = "y"
	return
}
