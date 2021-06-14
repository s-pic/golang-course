// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main

import "fmt"

// package scoped variables and constants
var x int = 100

const secondsInHour = 3600

// a function declaration. main is a special function name
// it is the entry point for the executable program
func main() {
	// Local Scoped Variables and Constants Declarations, calling functions etc
	var a int = 7
	var b float64 = 3.5
	const c int = 10

	// 06_fmt is a standard library package used
	// to print to standard output, pronounced famt

	fmt.Println(a, b, c)

	fmt.Println("Hello Go World!")
	distance := 60.8

	fmt.Printf("The distance in miles is %f n", distance*0.62137)
}
