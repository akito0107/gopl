package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const iterations = 128
	const contrast = 10
	var epsillon = math.Nextafter(1, 2) - 1

	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}

	fp := func(z complex128) complex128 {
		return 4 * z
	}

	for i := uint8(0); i < iterations; i++ {
		v := z - f(z)/fp(z)
		if cmplx.Abs((v-z)/z) < epsillon {
			return color.RGBA{i * contrast, 0, 0, 255}
		}
		z = v
	}

	return color.Black
}
