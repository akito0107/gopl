package ex06

import "testing"

func BenchmarkPopCountLowest100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowest(100)
	}
}

func BenchmarkPopCount100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(100)
	}
}

func BenchmarkPopCountShift100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(100)
	}
}

func BenchmarkPopCountLowestMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowest(18446744073709551615)
	}
}

func BenchmarkPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(18446744073709551615)
	}
}

func BenchmarkPopCountShiftMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(18446744073709551615)
	}
}

func BenchmarkPopCountLowest1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowest(1024)
	}
}

func BenchmarkPopCount1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1024)
	}
}

func BenchmarkPopCountShift1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(100)
	}
}
