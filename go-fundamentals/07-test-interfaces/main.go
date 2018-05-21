package main

import (
	"fmt"
)

type shape interface {
	printArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		base:   10,
		height: 5,
	}
	s := square{
		sideLength: 10,
	}
	fmt.Println(t.printArea(), s.printArea())
}

func (t triangle) printArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) printArea() float64 {
	return s.sideLength * s.sideLength
}
