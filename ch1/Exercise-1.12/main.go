package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
}

func lissajous(out io.Writer, c, r float64, sz int) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    //number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	if c == 0 {
		c = cycles
	}
	if r == 0 {
		r = res
	}
	if sz == 0 {
		sz = size
	}
	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	var colorIdx uint8 = 1
	for i := 0.0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*sz+1, 2*sz+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < c*2*math.Pi; t += r {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(sz+int(x*float64(sz)+0.5), sz+int(y*float64(sz)+0.5), colorIdx)
		}
		phase += 0.1
		if colorIdx++; colorIdx >= uint8(len(palette)) {
			colorIdx = 1
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var c, res float64
		var sz int64
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		} else {
			for k, v := range r.Form {
				if k == "size" {
					if sz, err = strconv.ParseInt(v[0], 10, 64); err != nil {
						sz = 0
						log.Panicln("parsing form, wrong int:", v[0])
					}
				}
				if k == "cycles" {
					if c, err = strconv.ParseFloat(v[0], 64); err != nil {
						c = 0
						log.Panicln("parsing form, wrong float:", v[0])
					}
				}
				if k == "res" {
					if res, err = strconv.ParseFloat(v[0], 64); err != nil {
						res = 0
						log.Panicln("parsing form, wrong float:", v[0])
					}
				}
			}
		}
		lissajous(w, c, res, int(sz))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
