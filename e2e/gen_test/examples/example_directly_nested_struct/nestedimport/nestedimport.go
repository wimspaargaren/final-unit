package nestedimport

type Hello struct {
	X int
}

type Hi struct {
	*Hello
}

type Other struct {
	X map[int]*Hi
}
