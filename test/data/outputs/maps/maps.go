package maps

func MapFunc(x map[int]string) map[int]string {
	return x
}

func MapUnSupportedKeyFunc(x map[chan int]string) map[chan int]string {
	return x
}

func MapUnSupportedValFunc(x map[int]chan int) map[int]chan int {
	return x
}
