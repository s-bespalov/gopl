package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"

	"github.com/s-bespalov/gopl/ch3/Exercise-3.5-3.9/fraktals"
)

func render(w io.Writer, f func(complex128) color.Color, xmin, ymin, xmax, ymax float64, width, height int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, f(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		render(rw, fraktals.Mandelbrot, -2., -2., +2., +2., 1024, 1024)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
