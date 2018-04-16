package ex02

import "testing"

type link struct {
	tail *link
}

type dup struct {
	a *link
	b *link
}

func TestIsCyclic(t *testing.T) {
	a, b, c := &link{}, &link{}, &link{}
	a.tail, b.tail, c.tail = b, a, c
	d := &link{}
	dup := &dup{d, d}

	cases := []struct {
		name string
		in  interface{}
		out bool
	}{
		{
			name: "int",
			in:  123,
			out: false,
		},
		{
			name: "link a",
			in:  a,
			out: true,
		},
		{
			name: "link b",
			in:  b,
			out: true,
		},
		{
			name: "link c",
			in:  c,
			out: true,
		},
		{
			name: "dup",
			in:  dup,
			out: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if act := IsCyclic(c.in); act != c.out {
				t.Errorf("must be same vale with %v, %t", c.in, c.out)
			}
		})
	}
}
