package main

import (
	"fmt"
	"math"
)

type IShape interface {
	// contains only method signatures
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

// adding a new method to circle
func (c Circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// both types Circle and Rectangle implement the shape interface implicitly
// since they have all of its methods declared via receiver functions

func printShape(s IShape) { // any type that implements the interface
	fmt.Println("Shape:", s)
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {

	// TYPE ASSERTIONS get the underlying type of an interface
	var s IShape = Circle{radius: 20}
	fmt.Printf("%T\n", s)

	// s.volume() -> error, since no method on interface type
	s.(Circle).volume() // type assertion -> extract the dynamic value of the interface value type-safe

	ball, ok := s.(Circle)
	if ok == true {
		fmt.Printf("Ball Volume: %v\n", ball.volume())
	}

	// A Type Switch runs a couple of type assertions in series and runs the first case with the
	// matching type

	switch value := s.(type) { // like a regular switch statement, but the cases specify types, to which the interface value is compared
	case Circle:
		fmt.Printf("%#v has circle type\n", value)
	case Rectangle:
		fmt.Printf("%#v has Rectangle type\n", value)

	}
}
