package arrays

func ArrayPointerInt() []*int {
	res := []*int{}
	for i := 0; i < 10; i++ {
		x := i + 1
		if i == 2 {
			res = append(res, nil)
		} else {
			res = append(res, &x)
		}
	}
	return res
}

type X struct {
	Y []*int
}

func ArrStructPointer() []*X {
	res := []*X{}
	for i := 0; i < 5; i++ {
		x := &X{
			Y: []*int{},
		}
		for j := 0; j < 5; j++ {
			y := j + 1
			x.Y = append(x.Y, &y)
		}
		res = append(res, x)
	}
	return res
}

func DoubleArray() [2][]int {
	return [2][]int{{3, 4}, {5, 6}}
}

func EmptyArray() []int {
	return []int{}
}
