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

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type lissajousOptions struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func handler(w http.ResponseWriter, r *http.Request) {
	opts := &lissajousOptions{
		cycles:  20,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
	c := r.URL.Query().Get("cycles")
	if cyc, e := strconv.Atoi(c); e != nil && c != "" {
		opts.cycles = cyc
	}
	res := r.URL.Query().Get("res")
	if re, e := strconv.ParseFloat(res, 64); e != nil && res != "" {
		opts.res = re
	}
	lissajous(w, opts)
}

func lissajous(out io.Writer, opts *lissajousOptions) {
	var palette = []color.Color{color.White, color.Black}

	const (
		whiteIndex = 0
		blackIndex = 1
	)

	var (
		cycles  = opts.cycles
		res     = opts.res
		size    = opts.size
		nframes = opts.nframes
		delay   = opts.delay
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
