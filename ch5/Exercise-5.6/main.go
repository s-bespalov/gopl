package main

import (
	"log"
	"net/http"

	"com.github/s-bespalov/gopl/ch5/Exercise-5.6/surface"
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
