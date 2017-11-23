package main

import "fmt"

type Kilo float64
type Pound float64

type Converter interface {
	Convert(f float64) Converter
}

func (k Kilo) String() string {
	return fmt.Sprintf("%g kg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g Lb", p)
}

func KToP(k Kilo) Pound {
	return Pound(k / 0.453592)
}

func PToK(p Pound) Kilo {
	return Kilo(p * 0.453592)
}

type Metre float64
type Feet float64

func (m Metre) String() string {
	return fmt.Sprintf("%g m", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g ft", f)
}

func FToM(f Feet) Metre {
	return Metre(f * 0.3048)
}

func MToF(m Metre) Feet {
	return Feet(m / 0.3048)
}
