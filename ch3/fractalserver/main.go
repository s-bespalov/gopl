// Server renders and sends an image with fractal in response
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/s-bespalov/gopl/ch3/fractalserver/render"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var x, y float64 = 5, 5
		zoom := 1

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			if k == "x" {
				if nwx, err := strconv.ParseFloat(v[0], 64); err == nil {
					x = nwx
				}
			}
			if k == "y" {
				if nwy, err := strconv.ParseFloat(v[0], 64); err == nil {
					y = nwy
				}
			}
			if k == "zoom" {
				if nwzoom, err := strconv.Atoi(v[0]); err == nil {
					zoom = nwzoom
				}
			}
		}
		render.Mandelbrot(w, x, y, zoom)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
