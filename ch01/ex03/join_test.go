package main

import (
	"strings"
	"testing"
)

func BenchmarkStringAdd(b *testing.B) {
	s1, s2 := "test1", "test2"
	for i := 0; i < b.N; i++ {
		_ = s1 + " " + s2
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	s := []string{"test1", "test2"}
	for i := 0; i < b.N; i++ {
		strings.Join(s, " ")
	}
}
