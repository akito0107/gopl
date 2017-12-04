package main

import "testing"

func Test_f(t *testing.T) {
	type in struct {
		i float64
		j float64
	}
	cases := []struct {
		in  in
		out float64
	}{
		{in{0, 0}, 0.0},
	}

	for _, c := range cases {
		if act := f(c.in.i, c.in.j); act != c.out {
			t.Error(c, act)
		}
	}
}
