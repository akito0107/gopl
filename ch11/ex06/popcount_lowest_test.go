package ex06

import (
	"testing"
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

func BenchmarkPopCountMid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(100)
	}
}

func BenchmarkPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowest(18446744073709551615)
	}
}

func BenchmarkPopCountBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1024)
	}
}
