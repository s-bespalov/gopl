// Newton emits a PNG image of the Newton fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -5, -5, +5, +5
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func newton(z complex128) color.Color {
	const iterations = 40
	const aroach = 0.1

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
