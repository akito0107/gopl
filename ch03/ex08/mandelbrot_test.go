package main

import (
	"io/ioutil"
	"testing"
)

func Benchmark_drawBig1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		draw(ioutil.Discard, mandelbrotBigComplex, 1024, 1024)
	}
}

func Benchmark_draw1281024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		draw(ioutil.Discard, mandelbrot128, 1024, 1024)
	}
}

func Benchmark_draw641024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		draw(ioutil.Discard, mandelbrot64, 1024, 1024)
	}
}
