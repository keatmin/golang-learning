package structs

import "math"

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Length)
}

func (r Rectangle) Area() float64 { // receiver variable to be first letter of the type
	return r.Length * r.Width
}

func (c Circle) Area() float64 { // method can be used instead of declaring a separate package to allow Area() to be defined again
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 { // method can be used instead of declaring a separate package to allow Area() to be defined again
	return 0.5 * t.Base * t.Height
}
