package main

import "testing"

func TestBigComplex_Add(t *testing.T) {
	type in struct {
		x *BigComplex
		y *BigComplex
	}
	cases := []struct {
		in  *in
		out *BigComplex
	}{
		{&in{
			NewBigComplex(1.0, 1.0),
			NewBigComplex(-2.0, 0.0),
		},
			NewBigComplex(-1.0, 1.0),
		},
	}

	for _, c := range cases {
		x := c.in.x
		y := c.in.y
		i := NewBigComplex(0.0, 0.0)
		if i.Add(x, y); !c.out.Equals(i) {
			t.Errorf("Must be equal expect: %+v, actual %+v", c.out, i)
		}
	}
}

func TestBigComplex_Sub(t *testing.T) {
	type in struct {
		x *BigComplex
		y *BigComplex
	}
	cases := []struct {
		in  *in
		out *BigComplex
	}{
		{&in{
			NewBigComplex(1.0, 0.0),
			NewBigComplex(-2.0, 2.0),
		},
			NewBigComplex(3.0, -2.0),
		},
	}

	for _, c := range cases {
		x := c.in.x
		y := c.in.y
		i := NewBigComplex(0.0, 0.0)
		if i.Sub(x, y); !c.out.Equals(i) {
			t.Errorf("Must be equal expect: %+v, actual %+v", c.out, i)
		}
	}
}

func TestBigComplex_Mul(t *testing.T) {
	type in struct {
		x *BigComplex
		y *BigComplex
	}
	cases := []struct {
		in  *in
		out *BigComplex
	}{
		{&in{
			NewBigComplex(24.0, 12.0),
			NewBigComplex(37.0, -13.0),
		},
			NewBigComplex(1044.0, 132.0),
		},
	}

	for _, c := range cases {
		x := c.in.x
		y := c.in.y
		i := NewBigComplex(0.0, 0.0)
		if i.Mul(x, y); !c.out.Equals(i) {
			t.Errorf("Must be equal expect: %+v, actual %+v", c.out, i)
		}
	}
}

func TestBigComplex_Quo(t *testing.T) {
	type in struct {
		x *BigComplex
		y *BigComplex
	}
	cases := []struct {
		in  *in
		out *BigComplex
	}{
		{&in{
			NewBigComplex(-1.0, 1.0),
			NewBigComplex(2.0, -4.0),
		},
			NewBigComplex(-0.3, -0.1),
		},
	}

	for _, c := range cases {
		x := c.in.x
		y := c.in.y
		i := NewBigComplex(0.0, 0.0)
		if i.Quo(x, y); !c.out.Equals(i) {
			t.Errorf("Must be equal expect: %+v, actual %+v", c.out, i)
		}
	}
}
