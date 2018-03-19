package ex02

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
