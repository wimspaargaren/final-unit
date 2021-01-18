package structexample

type Result struct {
	X int
	y string
}

func StructFunc() Result {
	return Result{
		X: 3,
		y: "hi",
	}
}

type X []Result

type Other struct {
	X X
}

func StructCustomTypeDef() *Other {
	return &Other{
		X: []Result{
			{
				X: 43,
			},
		},
	}
}

type Cycle struct {
	C *Cycle
	X int
}

func CycleStruct() *Cycle {
	return &Cycle{
		X: 1,
		C: &Cycle{
			X: 2,
			C: &Cycle{
				X: 3,
				C: &Cycle{
					X: 4,
					C: &Cycle{
						X: 5,
					},
				},
			},
		},
	}
}
