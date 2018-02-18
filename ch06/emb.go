package main

import (
	"fmt"
	"log"
)

type Point struct {
	A, B string
}

func (p *Point) String() string {
	return fmt.Sprintf("%s, %s\n", p.A, p.B)
}

type Square struct {
	A string
}

func (s *Square) String() string {
	return fmt.Sprintf("%s\n", s.A)
}

func (s *Square) Override(p *Point) {
	s.A = p.A
}

type Mixture struct {
	A string
	Point
	P Point
	Square
}

type Stringer func() string

func main() {
	m := &Mixture{Point: Point{A: "fuga", B: "hoge"}}
	fmt.Printf(m.Point.String())
	fmt.Printf(m.Point.A)
	m.Override(&m.Point)

	fmt.Printf("%s\n", m.Point)
	fmt.Printf("%s\n", m.Square)

	executor := func(stringer Stringer) {
		fmt.Println(stringer())
	}

	partial := func(p *Point) interface{} {
		return p.String
	}
	p := &Point{A: "hoge", B: "fuga"}

	executor(partial(p).(Stringer))
	partial(p).(Stringer)()
	var words []int
	log.Println("%d\n", len(words))

}
