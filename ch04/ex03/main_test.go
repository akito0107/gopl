package main

import "testing"

func Test_reverse(t *testing.T) {
	cases := []struct {
		in  [size]int
		out [size]int
	}{
		{
			in:  [size]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			out: [size]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, c := range cases {
		if reverse(&c.in); c.in != c.out {
			t.Errorf("expected: %s, but: %s", c.out, c.in)
		}
	}
}
