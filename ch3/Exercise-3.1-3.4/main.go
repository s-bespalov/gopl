package main

import (
	"log"
	"net/http"

	"github.com/s-bespalov/gopl/ch3/Exercise-3.1-3.4/surface"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "image/svg+xml")
		surface.Render(w, surface.Sinr)
	})
	http.HandleFunc("/egg", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "image/svg+xml")
		surface.Render(w, surface.EggBox)
	})
	http.HandleFunc("/saddle", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "image/svg+xml")
		surface.Render(w, surface.Saddle)
	})
	http.HandleFunc("/moguls", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "image/svg+xml")
		surface.Render(w, surface.Moguls)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
