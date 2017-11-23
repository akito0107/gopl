package main

import (
	"testing"

	"github.com/akito0107/gopl"
)

func TestPToK(t *testing.T) {
	if c := PToK(1); gopl.CompareFloat(float64(c), 0.4535) != 0 {
		t.Errorf("1 pound must be 0.4535kg, but %g\n", c)
	}

	if c := PToK(35); gopl.CompareFloat(float64(c), 15.8757) != 0 {
		t.Errorf("35 pound must be 15.8757kg, but %g\n", c)
	}
}

func TestKToP(t *testing.T) {
	if c := KToP(1); gopl.CompareFloat(float64(c), 2.2046) != 0 {
		t.Errorf("1kg must be 2.2046ld, but %g\n", c)
	}

	if c := KToP(35); gopl.CompareFloat(float64(c), 77.1618) != 0 {
		t.Errorf("35kg must be 77.1618ld, but %g\n", c)
	}
}

func TestFToM(t *testing.T) {
	if m := FToM(1); gopl.CompareFloat(float64(m), 0.3048) != 0 {
		t.Errorf("1 feet must be 0.3048m, but %g\n", m)
	}
}
func TestMToF(t *testing.T) {
	if m := MToF(25); gopl.CompareFloat(float64(m), 82.081) != 0 {
		t.Errorf("25 m must be 82.081ft, but %g\n", m)
	}
}
