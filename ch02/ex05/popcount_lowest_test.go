package ex03

import (
	"testing"

	"github.com/akito0107/gopl/ch02/popcount"
)

func TestPopCountLowest(t *testing.T) {
	if i := PopCountLowest(1); i != 1 {
		t.Errorf("must be 1 but %d", i)
	}

	if i := PopCountLowest(5); i != 2 {
		t.Errorf("must be 2 but %d", i)
	}

	if i := PopCountLowest(255); i != 8 {
		t.Errorf("must be 8 but %d", 8)
	}

	if i := PopCountLowest(256); i != 1 {
		t.Errorf("must be 1 but %d", i)
	}

	if i := PopCountLowest(18446744073709551615); i != 64 {
		t.Errorf("must be 64 but %d", i)
	}
}

func BenchmarkPopCountLowest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowest(100)
	}
}

func BenchmarkPopCountLowestMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowest(18446744073709551615)
	}
}

func BenchmarkPopCountLowestBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowest(1024)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(100)
	}
}

func BenchmarkPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowest(18446744073709551615)
	}
}

func BenchmarkPopCountBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(1024)
	}
}
