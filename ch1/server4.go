// Server2 os a minimal "echo" and counter server
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0 //first color in palette
	blackIndex = 1 //next color in palette
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles := 5
		size := 100
		delay := 8
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			if k == "cycles" {
				if newCyc, err := strconv.Atoi(v[0]); err == nil {
					cycles = newCyc
				}
			}
			if k == "delay" {
				if newDelay, err := strconv.Atoi(v[0]); err == nil {
					delay = newDelay
				}
			}
			if k == "size" {
				if newSize, err := strconv.Atoi(v[0]); err == nil {
					size = newSize
				}
			}
		}
		lissajous(w, cycles, size, delay)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, acycles int, asize int, delay int) {
	const (
		res     = 0.001 // angular resolution
		nframes = 64    // number of animation frames
	)
	cycles := float64(acycles)
	log.Print(fmt.Sprintf("cycles: %d, delay: %d, size: %d\n", acycles, asize, delay))
	size := float64(asize)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*asize+1, 2*asize+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			color := uint8(rand.Intn(3) + 1)
			img.SetColorIndex(asize+int(x*size+5), asize+int(y*size+5), color)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

// counter echoes the number of calls so far.package ch1
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
