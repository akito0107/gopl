package main

import "testing"

func Test_String(t *testing.T) {
	cases := []struct {
		name string
		expr string
		env  Env
	}{
		{
			name: "basic case",
			expr: "1 + 1",
			env:  nil,
		},
		{
			name: "with env",
			expr: "x + y * 2 + sin(z)",
			env:  Env{"x": 100, "y": 15, "z": 2.0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ex, _ := Parse(c.expr)
			want := ex.Eval(c.env)
			t.Log(ex.String())
			ex2, _ := Parse(ex.String())
			if ex2.Eval(c.env) != want {
				t.Errorf("must be same")
			}
		})
	}
}
