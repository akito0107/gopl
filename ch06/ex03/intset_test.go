package ex01

import "testing"

func TestIntSet_IntersectWith(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(128)

	var y IntSet
	y.Add(1)
	y.Add(10)

	x.IntersectWith(&y)

	if !x.Has(1) {
		t.Errorf("x must have 1")
	}

	if x.Has(2) {
		t.Errorf("x must not have 2")
	}

	if x.Has(3) {
		t.Errorf("x must not have 3")
	}
	if x.Has(128) {
		t.Errorf("x must not have 128")
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(128)

	var y IntSet
	y.Add(1)
	y.Add(10)

	x.DifferenceWith(&y)

	if x.Has(1) {
		t.Errorf("x must not have 1")
	}
	if !x.Has(2) {
		t.Errorf("x must have 2")
	}
	if !x.Has(3) {
		t.Errorf("x must have 3")
	}
	if x.Has(10) {
		t.Errorf("x must not have 10")
	}
	if !x.Has(128) {
		t.Errorf("x must have 128")
	}
}

func TestIntSet_SymmetricDifferenceWith(t *testing.T) {
	t.Run("x len is greater than y", func(t *testing.T) {
		var x IntSet
		x.Add(1)
		x.Add(2)
		x.Add(3)
		x.Add(128)

		var y IntSet
		y.Add(1)
		y.Add(10)

		x.SymmetricDifferenceWith(&y)

		if x.Len() != 4 {
			t.Errorf("x must 2, 3, 10, 128")
		}
		if !x.Has(2) || !x.Has(3) || !x.Has(10) || !x.Has(128) {
			t.Errorf("x must 2, 3, 10, 128")
		}
	})
	t.Run("x len is greater than y", func(t *testing.T) {
		var x IntSet
		x.Add(1)
		x.Add(2)
		x.Add(3)
		x.Add(128)

		var y IntSet
		y.Add(1)
		y.Add(128)
		y.Add(256)

		x.SymmetricDifferenceWith(&y)

		if x.Len() != 3 {
			t.Errorf("x must 2, 3, 256: but %s", x.String())
		}
		if !x.Has(2) || !x.Has(3) || !x.Has(256) {
			t.Errorf("x must 2, 3, 256")
		}
	})
}
