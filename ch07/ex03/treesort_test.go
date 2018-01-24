package ex03

import "testing"

func TestTree_String(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		out  string
	}{
		{
			name: "basic case",
			in:   []int{1, 2, 3, 4, 5},
			out:  "[1 2 3 4 5]",
		},
		{
			name: "sorted",
			in:   []int{2, 3, 1, 4, 5, 3},
			out:  "[1 2 3 3 4 5]",
		},
		{
			name: "zero length",
			in:   []int{},
			out:  "[]",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tree := Sort(c.in)
			if act := tree.String(); act != c.out {
				t.Errorf("string must be %s, but actually %s", c.out, act)
			}
		})
	}
}
