package ch06

import "testing"

func TestIntSet_AddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 10, 100, 1000)

	if !x.Has(1) {
		t.Error("x must have a 1")
	}
	if !x.Has(10) {
		t.Error("x must have a 10")
	}
	if !x.Has(100) {
		t.Error("x must have a 100")
	}
	if !x.Has(1000) {
		t.Error("x must have a 1000")
	}
}
