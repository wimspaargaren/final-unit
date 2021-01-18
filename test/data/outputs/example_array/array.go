package array

import "fmt"

func ArrayFunc() []int {
	x := []int{3}
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	return x
}
