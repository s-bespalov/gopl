package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/s-bespalov/gopl/ch3/Exercise-3.5-3.9/fraktals"
)

func render(w io.Writer, f fraktals.Fractal, width, height int, ox, oy float64, zoom float64, qualityup bool) {
	xmin := f.Xmin / zoom
	ymin := f.Ymin / zoom
	xmax := f.Xmax / zoom
	ymax := f.Ymax / zoom
	if qualityup {
		width *= 2
		height *= 2
	}
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
	if qualityup {
		img = supersampling(img)
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func supersampling(i *image.RGBA) *image.RGBA {
	result := image.NewRGBA(image.Rect(0, 0, i.Rect.Size().X/2, i.Rect.Size().Y/2))
	for x := 0; x < result.Rect.Size().X; x++ {
		for y := 0; y < result.Rect.Size().Y; y++ {
			ix, iy := x*2, y*2
			r, g, b, a := i.At(ix, iy).RGBA()
			/*
				var r1, g1, b1, a1 uint32
				if ix+1 < i.Rect.Size().X {
					r1, g1, b1, a1 = i.At(ix+1, y).RGBA()
				} else if ix-1 >= 0 {
					r1, g1, b1, a1 = i.At(ix-1, y).RGBA()
				} else {
					r1, g1, b1, a1 = r, g, b, a
				}

				var r2, g2, b2, a2 uint32
				if iy+1 < i.Rect.Size().Y {
					r2, g2, b2, a2 = i.At(ix, iy+1).RGBA()
				} else if iy-1 >= 0 {
					r2, g2, b2, a2 = i.At(ix, iy-1).RGBA()
				} else {
					r2, g2, b2, a2 = r, g, b, a
				}

				var r3, g3, b3, a3 uint32
				if iy+1 < i.Rect.Size().Y && ix+1 < i.Rect.Size().X {
					r3, g3, b3, a3 = i.At(ix+1, iy+1).RGBA()
				} else if iy-1 >= 0 && ix-1 >= 0 {
					r3, g3, b3, a3 = i.At(ix-1, iy-1).RGBA()
				} else {
					r3, g3, b3, a3 = r, g, b, a
				}

				r, g, b, a = (r+r1+r2+r3)/4, (g+g1+g2+g3)/4, (b+b1+b2+b3)/4, (a+a1+a2+a3)/4
			*/
			result.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}
	return result
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var ox, oy float64
		zoom := 1.
		q := false
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
				if k == "q" {
					if q, err = strconv.ParseBool(v[0]); err != nil {
						log.Println("error parsing bool parametr:", err)
						q = false
					}
				}
			}
		}
		render(rw, fraktals.Mandelbrot, 1024, 1024, ox, oy, zoom, q)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
