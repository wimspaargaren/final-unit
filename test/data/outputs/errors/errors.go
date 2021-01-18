package errors

import "fmt"

func ErrorFunc() error {
	return fmt.Errorf("hi")
}
