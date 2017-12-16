package main

import "testing"

func Test_reverse(t *testing.T) {
	type in struct {
		s []int
		i int
	}
	cases := []struct {
		in  in
		out []int
	}{
		{
			in: in{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			out: []int{3, 4, 5, 1, 2},
		},
		{
			in: in{
				[]int{1, 2, 3, 4, 5},
				4,
			},
			out: []int{5, 1, 2, 3, 4},
		},
	}
	for _, c := range cases {
		if act := rotate(c.in.s, c.in.i); !equals(act, c.out) {
			t.Errorf("expected: %v, but: %v", c.out, act)
		}
	}
}

func equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
