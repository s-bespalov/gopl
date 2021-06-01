package main

import (
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/s-bespalov/gopl/ch3/Exercise-3.5-3.9/fraktals"
)

func render(w io.Writer, f fraktals.Fractal, width, height int, ox, oy float64, zoom float64) {
	xmin := f.Xmin / zoom
	ymin := f.Ymin / zoom
	xmax := f.Xmax / zoom
	ymax := f.Ymax / zoom
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin + oy
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin + ox
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, f.F(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var ox, oy float64
		zoom := 1.
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		} else {
			for k, v := range r.Form {
				if k == "x" {
					if ox, err = strconv.ParseFloat(v[0], 10); err != nil {
						log.Println("error parsing float parametr:", err)
						ox = 0
					}
				}
				if k == "y" {
					if oy, err = strconv.ParseFloat(v[0], 10); err != nil {
						log.Println("error parsing float parametr:", err)
						oy = 0
					}
				}
				if k == "zoom" {
					if zoom, err = strconv.ParseFloat(v[0], 10); err != nil {
						log.Println("error parsing float parametr:", err)
						zoom = 1
					}
				}
			}
		}
		render(rw, fraktals.Mandelbrot, 1024, 1024, ox, oy, zoom)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
