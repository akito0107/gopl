package main

import (
	"fmt"
	"math/big"

	"github.com/ALTree/bigfloat"
)

type BigComplex struct {
	real      *big.Float
	imaginary *big.Float
}

func NewBigComplex(real float64, imaginary float64) *BigComplex {
	return &BigComplex{
		big.NewFloat(real),
		big.NewFloat(imaginary),
	}
}

func (bc *BigComplex) String() string {
	return fmt.Sprintf("%s, %s", bc.real.String(), bc.imaginary.String())
}

func (bc *BigComplex) Add(a *BigComplex, b *BigComplex) *BigComplex {
	bc.real.Add(a.real, b.real)
	bc.imaginary.Add(a.imaginary, b.imaginary)
	return bc
}

func (bc *BigComplex) Sub(a *BigComplex, b *BigComplex) *BigComplex {
	bc.real.Sub(a.real, b.real)
	bc.imaginary.Sub(a.imaginary, b.imaginary)
	return bc
}

func (bc *BigComplex) Mul(a *BigComplex, b *BigComplex) *BigComplex {
	arbr := big.NewFloat(0.0).Mul(a.real, b.real)
	aibi := big.NewFloat(0.0).Mul(a.imaginary, b.imaginary)
	arbi := big.NewFloat(0.0).Mul(a.real, b.imaginary)
	aibr := big.NewFloat(0.0).Mul(a.imaginary, b.real)

	bc.real.Sub(arbr, aibi)
	bc.imaginary.Add(arbi, aibr)

	return bc
}

func (bc *BigComplex) Quo(a *BigComplex, b *BigComplex) *BigComplex {
	br2 := big.NewFloat(0.0).Mul(b.real, b.real)
	bi2 := big.NewFloat(0.0).Mul(b.imaginary, b.imaginary)
	d := big.NewFloat(0.0).Add(bi2, br2)

	arbr := big.NewFloat(0.0).Mul(a.real, b.real)
	aibi := big.NewFloat(0.0).Mul(a.imaginary, b.imaginary)

	aibr := big.NewFloat(0.0).Mul(a.imaginary, b.real)
	arbi := big.NewFloat(0.0).Mul(a.real, b.imaginary)

	bc.real.Quo(big.NewFloat(0.0).Add(arbr, aibi), d)
	bc.imaginary.Quo(big.NewFloat(0.0).Sub(aibr, arbi), d)
	return bc
}

func (bc *BigComplex) Abs() *big.Float {
	r2 := big.NewFloat(0).Mul(bc.real, bc.real)
	i2 := big.NewFloat(0).Mul(bc.imaginary, bc.imaginary)

	return bigfloat.Sqrt(r2.Add(r2, i2))
}

func (bc *BigComplex) Equals(a *BigComplex) bool {
	return bc.real.Cmp(a.real) == 0 && bc.imaginary.Cmp(a.imaginary) == 0
}
