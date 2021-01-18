package basic

import "fmt"

// Foo example struct
type Foo struct{}

// Bar example function
func (f *Foo) Bar(x int) string {
	if x < 0 {
		return "x is a negative integer"
	} else if x > 0 && x < 10 {
		return "x between zero and 10"
	} else if x > 10 && x < 20 {
		return "x between 10 and 20"
	} else if x > 30 && x < 50 {
		return "x between 30 and 50"
	} else {
		return fmt.Sprintf("x not recognized: %d", x)
	}
}
