// Server renders and sends svg image in response
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/s-bespalov/gopl/ch2/surface/render"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		width, height := 600, 320
		pcolor, vcolor := "#ff0000", "#0000ff"

		// check parameters
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			if k == "width" {
				if nwidth, err := strconv.Atoi(v[0]); err == nil {
					width = nwidth
				} else {
					log.Print(err)
				}
			}
			if k == "height" {
				if nhight, err := strconv.Atoi(v[0]); err == nil {
					height = nhight
				} else {
					log.Print(err)
				}
			}
			if k == "pcolor" {
				pcolor = v[0]
			}
			if k == "vcolor" {
				vcolor = v[0]
			}
		}
		w.Header().Set("Content-type", "image/svg+xml")

		//render SVG and write response
		svg := render.Surf(height, width, pcolor, vcolor)
		w.Write([]byte(svg))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
