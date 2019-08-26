package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles float64) {
	var cy float64
	if cycles <= 0 {
		cy = 5
	} else {
		cy = cycles
	}
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cy*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var num int
		q := r.URL.Query()
		cycles, ok := q["cycles"]
		if ok {
			inty, err := strconv.Atoi(cycles[0]) // only parsing the first one
			if err != nil {
				fmt.Fprintf(w, "error parsing url query: %v", err)
				num = 5
			} else {
				num = inty
			}
		}
		lissajous(w, float64(num))
	})
	http.ListenAndServe(":8000", nil)
}
