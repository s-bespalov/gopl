package fraktals

import (
	"image/color"
	"log"
	"math"
	"math/cmplx"
)

var Mandelbrot Fractal
var MandelbrotColor Fractal
var Newton Fractal
var NewtonBW Fractal

var mdlColors []color.Color

func init() {
	Mandelbrot = Fractal{-2, 2, -2, 2, fMandelbrot}
	MandelbrotColor = Fractal{-2, 2, -2, 2, fMandelbrotColor}
	mdlColors = []color.Color{
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
	Newton = Fractal{-5, 5, -5, 5, fNewton}
	NewtonBW = Fractal{-5, 5, -5, 5, fNewtonbw}
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

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return mdlColors[int(n)%len(mdlColors)]
		}
	}
	return color.Black
}

func fNewton(z complex128) color.Color {
	const iterations = 40

	for n := uint8(0); n < iterations; n++ {
		v := z
		z = z - ((cmplx.Pow(z, 4) - 1) / (4 * cmplx.Pow(z, 3)))
		if cmplx.Abs(z-v) < 0.1 {
			contrast := uint8(255 - (n * 6))
			r := int(cmplx.Phase(z) / (math.Pi / 4))
			if r == 0 {
				return color.RGBA{contrast, 0, 0, 0xff}
			}
			if r == 1 || r == 2 {
				return color.RGBA{0, contrast, 0, 0xff}
			}
			if r >= 3 || r <= -3 {
				return color.RGBA{0, 0, contrast, 0xff}
			}
			if r == -1 || r == -2 {
				return color.RGBA{contrast, contrast, 0, 0xff}
			}
			log.Print(r)
			return color.White
		}
	}
	return color.Black
}

func fNewtonbw(z complex128) color.Color {
	const iterations = 40

	for n := uint8(0); n < iterations; n++ {
		v := z
		z = z - ((z*z*z*z)-1)/(4*z*z*z)
		if cmplx.Abs(z-v) < 0.1 {
			contrast := uint8(255 - (n * 6))
			return color.Gray{contrast}
		}
	}
	return color.Black
}
