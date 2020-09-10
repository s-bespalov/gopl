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
	img = supersampling(img)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func supersampling(img *image.RGBA) *image.RGBA {
	r := image.NewRGBA(img.Bounds())
	ymax := r.Bounds().Max.Y
	xmax := r.Bounds().Max.X
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			x1 := x + 1
			y1 := y + 1
			if x == xmax-1 {
				x1 = x - 1
			}
			if y == ymax-1 {
				y1 = y - 1
			}
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := img.At(x1, y).RGBA()
			r3, g3, b3, _ := img.At(x, y1).RGBA()
			r4, g4, b4, _ := img.At(x1, y1).RGBA()
			R := uint8((r1 + r2 + r3 + r4) / 4)
			G := uint8((g1 + g2 + g3 + g4) / 4)
			B := uint8((b1 + b2 + b3 + b4) / 4)
			c := color.RGBA{R, G, B, 0xff}
			r.Set(x, y, c)
		}
	}
	return r
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
