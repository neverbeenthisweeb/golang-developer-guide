package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{
		sideLength: 7,
	}

	tr := triangle{
		height: 5,
		base:   9,
	}

	// fmt.Printf("sq: %v\ntr: %v\n", sq.getArea(), tr.getArea())
	printArea(sq)
	printArea(tr)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Printf("Area of shape: %v\n", s.getArea())
}
