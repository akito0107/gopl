package ex07

import (
	"testing"
)

func BenchmarkIntSet_Add(b *testing.B) {
	x := IntSet{}
	for i := 0; i < b.N; i++ {
		x.Add(100)
	}
}

func BenchmarkIntSet_Has(b *testing.B) {
	x := IntSet{}
	for i := 0; i < b.N; i++ {
		x.Has(100)
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	x := &IntSet{}
	y := &IntSet{}
	for i := 0; i < b.N; i++ {
		x.UnionWith(y)
	}
}
