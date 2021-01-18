package maps

type X struct {
	Y *string
}

func MapFunc() map[int]*X {
	res := make(map[int]*X)
	lkj := "asdf"
	res[3] = &X{
		Y: &lkj,
	}
	return res
}

func MapNestedFunc() map[int]*map[string]X {
	res := make(map[int]*map[string]X)
	lkj := "asdf"
	asdf := make(map[string]X)
	asdf[lkj] = X{
		Y: &lkj,
	}
	res[3] = &asdf
	return res
}

func MapNestedNilFunc() map[int]*map[string]X {
	res := make(map[int]*map[string]X)
	res[3] = nil
	return res
}
