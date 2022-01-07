package main

import "fmt"

func main() {
	// Basics
	var numbers [4]int
	fmt.Printf("%v\n", numbers)  // prints [0 0 0 0]
	fmt.Printf("%#v\n", numbers) // prints [4]int{0, 0, 0, 0}

	// other ways to declare arrays
	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1) // prints [4]float64{0, 0, 0, 0}

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2) // [3]int{-10, 1, 100}

	a3 := [4]string{"Dan", "Diana", "John"} // last arg left blank
	// -> all elements not stated in the literal expression
	// are initialized to the zero-value for the given type
	fmt.Printf("%#v\n", a3) // [4]string{"Dan", "Diana", "John", ""}

	// Ellipsis Operator
	a4 := [...]int{
		1,
		2,
		3,
		-10,
		66, // mandatory comma -> multi-line literal
	} // length inferred from literal
	fmt.Printf("The length is %d\n", len(a4))
}
