package main

import (
	"fmt"
	"math"
)

//interface Shape definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

//Shape's implementations: Rectangle, Circle
type (
	Rectangle struct {
		Width  float64
		Height float64
	}
	Circle struct {
		Radius float64
	}
)

//Rectangle implements Shape ...
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2.0 * (r.Width + r.Height)
}

// ... as well as Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2.0 * math.Pi * c.Radius
}

func main() {
	var s Shape
	s = Rectangle{Width: 5.0, Height: 7.0}
	fmt.Printf("%T %v\n", s, s)
	fmt.Printf("area = %g, perimeter = %g\n", s.Area(), s.Perimeter())
	s = Circle{Radius: 3.5}
	fmt.Printf("%T %v\n", s, s)
	fmt.Printf("area = %g, perimeter = %g\n", s.Area(), s.Perimeter())
}
