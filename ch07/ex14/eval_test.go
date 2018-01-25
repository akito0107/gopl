package main

import "testing"

func Test_String(t *testing.T) {
	cases := []struct {
		name string
		expr string
		want float64
		env  Env
	}{
		{
			name: "basic case",
			expr: "1 + 1 - abs(4)",
			want: -2,
			env:  nil,
		},
		{
			name: "with env",
			expr: "x + y * 2 + abs(z)",
			want: 133,
			env:  Env{"x": 100, "y": 15, "z": 3},
		},
		{
			name: "with negative value",
			expr: "x + y * 2 + abs(z)",
			want: 140,
			env:  Env{"x": 100, "y": 15, "z": -10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ex, _ := Parse(c.expr)
			if act := ex.Eval(c.env); act != c.want {
				t.Errorf("must match: expect %g, but %g", c.want, act)
			}
		})
	}
}
