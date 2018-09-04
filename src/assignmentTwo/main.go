package main

import (
	"fmt"
	"math"
)

type triangle struct {
	hight float64
	base  float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{
		side: 10,
	}

	t := triangle{
		hight: 10,
		base:  5.5,
	}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}

func (t triangle) getArea() float64 {
	return t.base * t.hight / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
