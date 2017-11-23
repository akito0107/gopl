package ex03

import (
	"testing"

	"github.com/akito0107/gopl/ch02/popcount"
)

func TestPopCountShift(t *testing.T) {
	if i := PopCountShift(1); i != 1 {
		t.Errorf("must be 1 but %d", i)
	}

	if i := PopCountShift(5); i != 2 {
		t.Errorf("must be 2 but %d", i)
	}

	if i := PopCountShift(255); i != 8 {
		t.Errorf("must be 8 but %d", 8)
	}

	if i := PopCountShift(256); i != 1 {
		t.Errorf("must be 1 but %d", i)
	}

	if i := PopCountShift(18446744073709551615); i != 64 {
		t.Errorf("must be 64 but %d", i)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(100)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(100)
	}
}
