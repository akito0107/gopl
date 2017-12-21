package main

import "testing"

func Test_bitcount(t *testing.T) {
	a := [32]byte{1}
	b := [32]byte{0}
	if act := bitcount(a, b); act != 1 {
		t.Error("expected: %d, but %d", 1, act)
	}

	a = [32]byte{255}
	b = [32]byte{0}
	if act := bitcount(a, b); act != 8 {
		t.Error("expected: %d, but %d", 8, act)
	}
}
