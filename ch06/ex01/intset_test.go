package ex01

import "testing"

func TestIntSet_Len(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var x IntSet
		x.Add(1)
		x.Add(2)
		x.Add(92)
		x.Add(128)

		if x.Len() != 4 {
			t.Errorf("length must 4 but %d\n", x.Len())
		}
	})
	t.Run("null value", func(t *testing.T) {
		var x IntSet

		if x.Len() != 0 {
			t.Errorf("length must 4 but %d\n", x.Len())
		}
	})
}

func TestIntSet_Remove(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var x IntSet
		x.Add(1)
		x.Add(67)
		x.Add(259)
		if !x.Has(1) {
			t.Fatal("unexpcted removal")
		}
		x.Remove(1)

		if x.Has(1) {
			t.Error("1 must be removed")
		}
		x.Remove(67)

		if x.Has(67) {
			t.Error("67 must be removed")
		}
	})
	t.Run("null", func(t *testing.T) {
		var x IntSet
		x.Remove(1)
		if x.Has(1) {
			t.Error("1 must be exists")
		}
	})
}

func TestIntSet_Clear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Clear()

	if x.Len() != 0 {
		t.Error("must be cleared")
	}
}

func TestIntSet_Copy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(10)
	x.Add(100)

	y := x.Copy()

	if !y.Has(1) {
		t.Error("y must have a 1")
	}
	if !y.Has(10) {
		t.Error("y must have a 10")
	}
	if !y.Has(100) {
		t.Error("y must have a 100")
	}
}
