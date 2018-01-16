package main

import (
	"math"
)

//Shape is an embedded type
type Shape interface {
	area() float64
}

//Circle holds center and radius data
type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

//Rectangle contains ords for corners
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	return (r.x2 - r.x1) * (r.y2 - r.y1)
}
