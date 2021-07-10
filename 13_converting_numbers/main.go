package main

import "fmt"

func main() {
	// in go we speak of converting, not casting

	// conversions can be explicit, or implicit depending on context

	var x = 3   // int inferred to be the type of the untyped constant 3
	var y = 3.9 // float inferred
	// x = x * y // compile time error

	x = x * int(y)
	println(x)                 // prints 9
	fmt.Printf("%T\n", y)      // float
	fmt.Printf("%T\n", int(y)) // int

	x = int(float64(x) * y)

	fmt.Printf("%v", x) // -> 35

	// -------- another example

	var a = 5       // int
	var b int64 = 2 // in64 -> another type than int

	//var c = a & b -> type mismatch, explicit conversion needed

	a = int(b) // conversion done, works

	_ = a
}
