package ex07

import "testing"

func BenchmarkMapIntSet_Add(b *testing.B) {
	x := NewMapIntSet()
	for i := 0; i < b.N; i++ {
		x.Add(100)
	}
}

func BenchmarkMapIntSet_Has(b *testing.B) {
	x := NewMapIntSet()
	for i := 0; i < b.N; i++ {
		x.Has(100)
	}
}

func BenchmarkMapIntSet_UnionWith(b *testing.B) {
	x := NewMapIntSet()
	y := NewMapIntSet()
	for i := 0; i < b.N; i++ {
		x.UnionWith(y)
	}
}
