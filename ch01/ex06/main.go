package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"

	"log"

	"os"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles   = 5
		res      = 0.001
		size     = 100
		nframes  = 64
		delay    = 8
		colorlen = 255
	)
	var palette = []color.Color{color.Black}

	for i := 0; i < colorlen; i++ {
		palette = append(palette, genColor())
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		usage := getCpuUsage()
		var cnt int
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			cnt++
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if cnt%1000 == 0 {
				usage = getCpuUsage()
				cnt = 0
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), scale(usage))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

var i uint8

func genColor() *color.RGBA {
	i++
	if i > 255 {
		i = 255
	}
	return &color.RGBA{R: i, G: 0x00, B: 0x00, A: 0xff}
}

func getCpuUsage() uint8 {
	usage, err := cpu.Percent(5*time.Millisecond, false)
	if err != nil {
		log.Fatal(err)
	}
	return uint8(usage[0])
}

func scale(i uint8) uint8 {
	i = i * (255 / 100)
	if i > 255 {
		return 255
	}
	return i
}
