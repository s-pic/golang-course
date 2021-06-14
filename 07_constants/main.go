package main

import "fmt"

// constants can be numbers, strings and booleans

func main() {

	// in contrast to variables
	// - must be initialized to a value
	// - must not be used in the remaining script
	const days = 7

	const pi float64 = 3.14
	const hourInSeconds = 3600

	durationInHours := 234

	fmt.Printf("Duration in seconds: %v\n", durationInHours*hourInSeconds)

	// using constants can prevent runtime errors like the one below

	// x, y := 5, 0
	// fmt.Println(x / y)
	// commenting the above in would lead to a Runtime
	// "panic: runtime error: integer divide by zero

	// const a, b = 5, 0
	// fmt.Println(a / b)
	// commenting the above in lets the IDE detect a compile error
	// "Division by 0"

	// we can declare multiple constants using a grouped constant

	const (
		x = 2
		y = 4
		z = 5
	)

	fmt.Println(x, y, z)

	// something strange: in a grouped constant
	const (
		m = 500
		n
		o
	)

	fmt.Println(m, n, o) // prints "500 500 500"

	// n and o get their type from the previous constant

	// constant rules
	// - cannot be reassigned a value
	// - cannot be initialized at run time
	//   (and thus *generally *not using a function call)
	// - cannot be initialized with a variable
	// - CAN be initialized with a function call like `len` (returning the number of elements in string)
	//   IF the argument passed is a string literal, since the literal is an unnamed constant

	const l1 = len("hello")
	println("l1 is ", l1)
}
