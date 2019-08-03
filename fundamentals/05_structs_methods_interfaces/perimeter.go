package perimeter

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
