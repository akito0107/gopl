package ex02

import "testing"

func TestIntSet(t *testing.T) {
	cases := []struct {
		name string
		adds []int
	}{
		{
			name: "only one value",
			adds: []int{1},
		},
		{
			name: "multiple value",
			adds: []int{1, 2, 4, 6},
		},
	}

	for _, c := range cases {
		x := IntSet{}
		y := NewMapIntSet()
		for _, a := range c.adds {
			x.Add(a)
			y.Add(a)
		}

		/*
			if x.String() != y.String() {
				t.Errorf("must be same with expect %s, but %s", y.String(), x.String())
			}
		*/
		for _, a := range c.adds {
			if x.Has(a) != y.Has(a) {
				t.Errorf("must be same with expect %v, but %v", x.Has(a), y.Has(a))
			}
		}
	}
}

func TestMapIntSet_UnionWith(t *testing.T) {
	cases := []struct {
		name   string
		src    []int
		target []int
	}{
		{
			name:   "only one value",
			src:    []int{1},
			target: []int{2},
		},
		{
			name:   "multiple value",
			src:    []int{1, 2, 4, 6},
			target: []int{10, 12},
		},
	}

	for _, c := range cases {
		xs := IntSet{}
		ys := NewMapIntSet()
		xt := IntSet{}
		yt := NewMapIntSet()

		for _, a := range c.src {
			xs.Add(a)
			ys.Add(a)
		}
		for _, a := range c.target {
			xt.Add(a)
			yt.Add(a)
		}

		xs.UnionWith(&xt)
		ys.UnionWith(yt)
		/*
			if xs.String() != ys.String() {
				t.Errorf("must be same with expect %s, but %s", ys.String(), xs.String())
			}
		*/
		for _, a := range xs.Elems() {
			if !ys.Has(a) {
				t.Errorf("must be same with expect %v, but %v", xs.Has(a), ys.Has(a))
			}
		}

		for k, _ := range ys.set {
			if !xs.Has(k) {
				t.Errorf("must be same with expect %v, but %v", xs.Has(k), ys.Has(k))
			}
		}
	}

}
