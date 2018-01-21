package ex04

import "testing"

func TestIntSet_Elems(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(126)
	x.Add(356)

	res := x.Elems()
	if len(res) != 3 {
		t.Errorf("len must 3, but %d", len(res))
	}
}
