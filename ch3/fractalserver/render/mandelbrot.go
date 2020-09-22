// Package render renders fractals to output
package render

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

// Mandelbrot renders mandelbrot fractal to output
func Mandelbrot(out io.Writer, ix float64, iy float64, izoom int) {
	const (
		width, height = 1024, 1024
	)
	xmin, ymin, xmax, ymax := -ix, -iy, ix, iy

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
	//if izoom > 1 {
	img = zoom(img, izoom)
	//}
	png.Encode(out, img) // NOTE: ignoring errors
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

func zoom(img *image.RGBA, zoom int) *image.RGBA {
	r := image.NewRGBA(img.Bounds())
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	xstart := (w - w/zoom) / 2
	ystart := (h - h/zoom) / 2
	counterx, countery := 0, 0
	for xr := 0; xr < w; xr++ {
		ycurrent := ystart
		countery = 0
		for yr := 0; yr < h; yr++ {
			r.Set(xr, yr, img.At(xstart, ycurrent))
			//log.Print(xstart, ycurrent, img.At(xstart, ystart))
			countery++
			if countery >= zoom {
				ycurrent++
				countery = 0
			}
		}
		counterx++
		if counterx >= zoom {
			xstart++
			counterx = 0
		}
	}
	return r
}
