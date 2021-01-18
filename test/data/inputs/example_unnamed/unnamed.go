package array

type Other struct {
	Y string
}

type Simple struct {
	X struct {
		Y string
	}
	Z Other
}

func SimpleUnnamed(x Simple) {}

type SimpleInterface struct {
	X interface {
		Hi(x int)
	}
}

func SimpleInterfaceUnnamed(x SimpleInterface) {}
