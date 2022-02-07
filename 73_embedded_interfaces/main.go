package main

import (
	"fmt"
)

func main() {
	// in go we cannot extend from an interface, but we can merge them
	// (which is known as interface embedding)

	type Shape interface {
		area() float64
	}

	type Object interface {
		volume() float64
	}

	// an interface can include (embed) other interface

	type Geometry interface {
		Shape  // all methods from Shape
		Object // all methods from Object
		getColor() string
	}

	type Cube struct {
		edge float64
	}

	// EMPTY INTERFACEs work like a top type (`any` in TS)
	// and thus should be handled with care (we use types for a reason hey)

	var empty interface{}

	empty = 5
	empty = "Go" // can take any value
	empty = []int{4, 5, 6}

	fmt.Println((empty))

	// fmt.Println(len(empty)) // compile time error

	fmt.Println(len(empty.([]int)))

	type Person struct {
		info interface{}
	}

	you := Person{info: "anything"}
	you.info = 5

}
