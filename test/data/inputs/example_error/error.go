package error

func ErrorFunc(x error) {
}

func ErrorPointerFunc(x *error) {
}

func main() {
	err := func() error {
		return nil
	}()
	ErrorFunc(err)
}
