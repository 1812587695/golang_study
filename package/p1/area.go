package main

import (
	"math"
)

type Circle struct {
	Radius float64
}

func newCircle() *Circle {
	return &Circle{Radius: 1}
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
