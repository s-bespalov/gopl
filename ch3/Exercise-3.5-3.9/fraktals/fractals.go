package fraktals

import (
	"image/color"
	"math/cmplx"
)

var Mandelbrot Fractal
var MandelbrotColor Fractal

var colors []color.Color

func init() {
	Mandelbrot = Fractal{-2, 2, -2, 2, fMandelbrot}
	MandelbrotColor = Fractal{-2, 2, -2, 2, fMandelbrotColor}
	colors = []color.Color{
		color.RGBA{66, 30, 15, 255},
		color.RGBA{25, 7, 26, 255},
		color.RGBA{9, 1, 47, 255},
		color.RGBA{4, 4, 73, 255},
		color.RGBA{0, 7, 100, 255},
		color.RGBA{12, 44, 138, 255},
		color.RGBA{24, 82, 177, 255},
		color.RGBA{57, 125, 209, 255},
		color.RGBA{134, 181, 229, 255},
		color.RGBA{211, 236, 248, 255},
		color.RGBA{241, 233, 191, 255},
		color.RGBA{248, 201, 95, 255},
		color.RGBA{255, 170, 0, 255},
		color.RGBA{204, 128, 0, 255},
		color.RGBA{153, 87, 0, 255},
		color.RGBA{106, 52, 3, 255},
	}
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

func fMandelbrotColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colors[int(n)%len(colors)]
		}
	}
	return color.Black
}
