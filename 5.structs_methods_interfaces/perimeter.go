package main

import (
	"math"
)

func main() {
	rectangle := Rectangle{Width: 5, Height: 2}

}

type Shape interface { // the 'meta' of Rectangle and Circle -- 
// contract of behaviour (fulfills method with the same function signature)
// keep as small as possible, flexible as possible
// all methods are mandatory to fulfill to satisfy the interface
// can be combined
	Area() float64
}

type Rectangle struct { //implementor of the behaviour
	Width float64
	Height float64
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2*(r.Width + r.Height)
}
func (r *Rectangle) SetHeight(h float64) {
	r.Height = h
}

var somevariable := 0


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