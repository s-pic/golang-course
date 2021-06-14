package main

import "fmt"

func main() {
	const a float64 = 5.1 // typed constant
	const b = 6.7         // untyped constant

	// they behave differently

	// constant expressions are always evaluated at compile time
	const c float64 = a * b
	const str = "Hello " + "Go!" // same thing, since literals (anonymous constants)

	const d = 5 > 1
	fmt.Println("d:", d)

	// normally whe can not multiply values
	// of different types (e.g. int and float)

	// const x int = 5
	// const y float64 = 2.2 * x // Invalid operation: 2.2 * x (cannot convert the constant 2.2 to the type int)

	const x = 5
	const y = 2.2 * 5 // no error since using untyped constants

	fmt.Printf("%T %v\n", y, y) // prints float64

	// what happened here?
	// 1. When the variable was used for the first time, its type was determined
	// based on the expression it is used in
	// 2. When an untyped constant is used in a context where a type is
	// required, the type will be inferred.
	// Untyped constants exist in space next to the strong type world
	// of go constants, they give greater freedom.

	// another example

	var i int = x     // x is typeless changes to int
	var j float64 = x // behind the scenes: var j float64 = float64(x)
	var p byte = x

	fmt.Println(i, j, p)

	// all this is still a bit blurry:
	// https://blog.golang.org/constants
}
