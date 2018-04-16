package ex06

import (
	"math"
	"testing"
)

var result int
var value uint64 = 100

func BenchmarkPopCountLowest100(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountLowest(value)
	}
	result = s
}

func BenchmarkPopCount100(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCount(value)
	}
	result = s
}

func BenchmarkPopCountShift100(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountShift(value)
	}
	result = s
}

func BenchmarkPopCountLowestMax(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountLowest(math.MaxInt64)
	}
	result = s
}

func BenchmarkPopCountMax(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCount(math.MaxInt64)
	}
	result = s
}

func BenchmarkPopCountShiftMax(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountShift(math.MaxInt64)
	}
	result = s
}

var value2 uint64 = 1024

func BenchmarkPopCountLowest1024(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountLowest(value2)
	}
	result = s
}

func BenchmarkPopCount1024(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCount(value2)
	}
	result = s
}

func BenchmarkPopCountShift1024(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountShift(value2)
	}
	result = s
}
