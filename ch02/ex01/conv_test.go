package main

import (
	"testing"

	"github.com/akito0107/gopl"
)

func TestKToC(t *testing.T) {
	if c := KToC(0); gopl.DiffAbs(float64(c), -273.15) > gopl.EPSILON {
		t.Errorf("0K must be -273.15℃, but %g\n", c)
	}

	if c := KToC(325); gopl.DiffAbs(float64(c), 51.85) > gopl.EPSILON {
		t.Errorf("325K must be 51.85℃, but %g\n", c)
	}
}

func TestCToK(t *testing.T) {
	if k := CToK(0); gopl.DiffAbs(float64(k), 273.15) > gopl.EPSILON {
		t.Errorf("0℃ must be 273.15k, but %g\n", k)
	}

	if k := CToK(51.85); gopl.DiffAbs(float64(k), 325) > gopl.EPSILON {
		t.Errorf("51.85℃ must be 325k, but %g\n", k)
	}
}
