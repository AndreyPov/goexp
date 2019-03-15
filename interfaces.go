package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectan struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectan) area() float64 {
	return r.height * r.width
}

func (r rectan) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectan{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
