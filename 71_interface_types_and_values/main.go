package main

import (
	"fmt"
	"math"
)

type IShape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}
func main() {
	var s IShape // zero value -> nil
	fmt.Printf("Type of IShape: %T\n", s)

	ball := Circle{radius: 5}

	s = ball // Circle is the underlying, concrete type
	fmt.Printf("Type of s: %T\n", s)

	s = Rectangle{
		width:  10,
		height: 20,
	}

	// -> values typed as Interface are dynamically typed,
	// their underlying types can change at runtime

	// (like a Union in TS)

}
