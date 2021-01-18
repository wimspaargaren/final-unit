package color

import "image/color"

func X(c color.Color) {
}

func Y(x func() (a, b int)) func() (c, d string) {
	return func() (c, d string) {
		return "hi", "hi"
	}
}
