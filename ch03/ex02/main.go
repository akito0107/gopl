package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	alpha         = 10
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var form string = "magura"

func main() {
	fmt.Printf("<svg xmls='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	var z float64
	if form == "mogul" {
		z = mogul(x, y)
	} else if form == "magura" {
		z = hyperbolicParaboloid(x, y)
	} else {
		x = f(x, y)
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func mogul(x, y float64) float64 {
	res := (math.Sin(x) * math.Cos(y)) / alpha
	if math.IsNaN(res) {
		return 0.0
	}
	return res
}

func hyperbolicParaboloid(x, y float64) float64 {
	a := 25.0
	b := 15.0
	a2 := a * a
	b2 := b * b

	return y*y/a2 - x*x/b2
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	res := math.Sin(r) / r
	if math.IsNaN(res) {
		return 0.0
	}
	return res
}
