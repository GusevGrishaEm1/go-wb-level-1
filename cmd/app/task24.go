package main

import (
	"math"
)

type Point struct {
    x, y float64
}

func NewPoint(x, y float64) Point {
    return Point{x, y}
}

func (p Point) DistanceTo(p2 Point) float64 {
    dx := p.x - p2.x
    dy := p.y - p2.y
    return math.Sqrt(dx*dx + dy*dy)
}