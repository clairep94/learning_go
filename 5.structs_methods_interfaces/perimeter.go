package main

import (
	"math"
)

type Shape interface { // the 'meta' of Rectangle and Circle
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2*(r.Width + r.Height)
}

type Circle struct {
	Radius float64
}
func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}
func (c Circle) Perimeter() float64 {
	return (2 * math.Pi * c.Radius)
}

type Triangle struct {
	Height float64
	Base float64
}

func (tr Triangle) Area() float64 {
	return (tr.Height * tr.Base)/2
}