package ex07

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

var seed = time.Now().UTC().UnixNano()

func BenchmarkIntSet_Add(b *testing.B) {
	f := func() {
		x := IntSet{}
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 500; i++ {
			x.Add(rng.Intn(math.MaxInt16))
		}
	}
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkIntSet_Has(b *testing.B) {
	x := IntSet{}
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 500; i++ {
		x.Add(rng.Intn(math.MaxInt16))
	}
	f := func() {
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 500; i++ {
			x.Has(rng.Intn(math.MaxInt16))
		}
	}
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	f := func() {
		x := IntSet{}
		y := IntSet{}
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 500; i++ {
			x.Add(rng.Intn(math.MaxInt16))
			y.Add(rng.Intn(math.MaxInt16))
		}
		x.UnionWith(&y)
	}

	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkMapIntSet_Add(b *testing.B) {
	f := func() {
		x := NewMapIntSet()
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 500; i++ {
			x.Add(rng.Intn(math.MaxInt16))
		}
	}
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkMapIntSet_Has(b *testing.B) {
	x := NewMapIntSet()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 500; i++ {
		x.Add(rng.Intn(math.MaxInt16))
	}
	f := func() {
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 500; i++ {
			x.Has(rng.Intn(math.MaxInt16))
		}
	}
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkMapIntSet_UnionWith(b *testing.B) {
	f := func() {
		x := NewMapIntSet()
		y := NewMapIntSet()
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 500; i++ {
			x.Add(rng.Intn(math.MaxInt16))
			y.Add(rng.Intn(math.MaxInt16))
		}
		x.UnionWith(y)
	}

	for i := 0; i < b.N; i++ {
		f()
	}
}
