package ex03

import (
	"testing"

	"github.com/akito0107/gopl/ch02/popcount"
)

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(100)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(100)
	}
}
