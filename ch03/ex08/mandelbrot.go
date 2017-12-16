package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
)

type mandelbrot func(complex128) color.Color

func main() {
	draw(os.Stdout, mandelbrotBigComplex, 1024, 1024)
}

func draw(out io.Writer, f mandelbrot, width, height int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			img.Set(px, py, f(complex(x, y)))
		}
	}
	png.Encode(out, img)
}

func mandelbrotBigComplex(z complex128) color.Color {
	c := NewBigComplex(float64(real(z)), float64(imag(z)))

	const iterations = 200
	const contrast = 15

	v := NewBigComplex(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v.Mul(v, v)
		v.Add(v, c)
		if v.Abs().Cmp(big.NewFloat(2)) >= 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot64(z complex128) color.Color {
	c := complex64(z)
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + c
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
