package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (c *circle) area() float64  {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (r *rect) perim() float64 {
	return 2 * r.height + 2 * r.width
}

func geometryArea(geometry2 geometry) float64 {
	return geometry2.area()
}

func main() {
	c1 := &circle{radius: 4}
	fmt.Println(geometryArea(c1))

	r1 := &rect{
		width:  4,
		height: 3,
	}

	fmt.Println(geometryArea(r1))
}
