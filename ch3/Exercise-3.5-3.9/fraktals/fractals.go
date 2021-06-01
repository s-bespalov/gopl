package fraktals

import (
	"image/color"
	"math/cmplx"
)

var Mandelbrot Fractal

func init() {
	Mandelbrot = Fractal{-2, 2, -2, 2, fMandelbrot}
}

type Fractal struct {
	Xmin, Xmax, Ymin, Ymax float64
	F                      func(complex128) color.Color
}

// fMandelbrot returns colors of z complex number
func fMandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
