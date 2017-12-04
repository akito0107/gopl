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
	max           = 1.0
	min           = -0.3
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmls='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			//_, _, az := corner(i+1, j)
			//_, _, bz := corner(i, j)
			//_, _, cz := corner(i, j+1)
			//_, _, dz := corner(i+1, j+1)

			color := colorScale((az + bz + cz + dz) / 4)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	res := math.Sin(r) / r
	if math.IsNaN(res) {
		return 0.0
	}
	return res
}

func colorScale(z float64) string {
	h := z - min
	u := 512 / (math.Abs(max) + math.Abs(min))
	if scale := h * u; scale > 256.0 {
		return fmt.Sprintf("#%s0000", fmtHex(int(math.Ceil(scale)-256)))
	} else {
		return fmt.Sprintf("#0000%s", fmtHex(int(256-math.Ceil(scale))))
	}
}

func fmtHex(hex int) string {
	if hex < 16 {
		return fmt.Sprintf("0%x", hex)
	}
	return fmt.Sprintf("%x", hex)
}
