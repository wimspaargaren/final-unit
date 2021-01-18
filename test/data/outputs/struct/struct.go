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
