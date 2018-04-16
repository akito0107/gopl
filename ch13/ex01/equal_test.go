package ex01

import "testing"

func TestEqual(t *testing.T) {
	cases := []struct {
		x   float64
		y   float64
		out bool
	}{
		{
			x:   1.0,
			y:   1.0,
			out: true,
		},
		{
			x: 0.000000001,
			y: 0.000000002,
			out: false,
		},
		{
			x: 0.000000001,
			y: 0.0000000009,
			out: true,
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if Equal(c.x, c.y) != c.out {
				t.Errorf("must be same x: %G, y: %G, result: %t", c.x, c.y, c.out)
			}
		})
	}
}
