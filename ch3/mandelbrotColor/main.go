// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 273
	const contrast = 15

	var v complex128
	for n := 0; n <= iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			c := n * contrast
			R := uint8((c>>6)&0x7) << 5
			G := uint8((c>>3)&0x7) << 5
			B := uint8(c&0x7) << 5
			//log.Print(R, G, B, c)
			return color.RGBA{R, G, B, 0xff}
		}
	}
	return color.Black
}
