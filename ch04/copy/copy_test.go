package copy

import (
	"math/rand"
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	b.StopTimer()
	s := make([]interface{}, b.N+1)
	for i := 0; i < b.N+1; i++ {
		s[i] = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s = append(s[:1], s[2:]...)
	}
}

func BenchmarkCopy(b *testing.B) {
	b.StopTimer()
	s := make([]interface{}, b.N+1)
	for i := 0; i < b.N+1; i++ {
		s[i] = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(s[1:], s[2:])
		s = s[:len(s)-1]
	}
}
