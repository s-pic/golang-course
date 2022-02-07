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

// both types Circle and Rectangle implement the shape interface implicitly
// since they have all of its methods declared via receiver functions

func printShape(s IShape) { // any type that implements the interface
	fmt.Println("Shape:", s)
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {
	c1 := Circle{radius: 5}
	r1 := Rectangle{
		width:  10,
		height: 10,
	}
	printShape(c1)
	printShape(r1)
}
