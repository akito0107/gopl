package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type drawOpts struct {
	width   int
	height  int
	cells   int
	xyrange float64
	color   string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	opts := &drawOpts{
		width:   600,
		height:  320,
		cells:   100,
		xyrange: 30.0,
		color:   "white",
	}
	wid := r.URL.Query().Get("w")
	if wd, e := strconv.Atoi(wid); e == nil && wid != "" {
		opts.width = wd
	}
	ht := r.URL.Query().Get("h")
	if h, e := strconv.Atoi(ht); e == nil && ht != "" {
		opts.height = h
	}
	c := r.URL.Query().Get("c")
	if c != "" {
		opts.color = c
	}
	log.Println(ht, wid)
	log.Printf("%+v\n", opts)
	w.Header().Set("Content-Type", "image/svg+xml")
	draw(w, opts)
}

func draw(out io.Writer, opts *drawOpts) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", opts.color, opts.width, opts.height)
	for i := 0; i < opts.cells; i++ {
		for j := 0; j < opts.cells; j++ {
			ax, ay := corner(i+1, j, opts)
			bx, by := corner(i, j, opts)
			cx, cy := corner(i, j+1, opts)
			dx, dy := corner(i+1, j+1, opts)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, opts *drawOpts) (float64, float64) {
	x := opts.xyrange * (float64(i)/float64(opts.cells) - 0.5)
	y := opts.xyrange * (float64(j)/float64(opts.cells) - 0.5)

	z := f(x, y)

	xyscale := float64(opts.width) / 2 / float64(opts.xyrange)
	zscale := float64(opts.height) * 0.4

	sx := float64(opts.width)/2 + (x-y)*cos30*xyscale
	sy := float64(opts.height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	res := math.Sin(r) / r
	if math.IsNaN(res) {
		return 0.0
	}
	return res
}
