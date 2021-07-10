package main

import "fmt"

/*
Consider the following declarations:

var i int = 3
var f float64 = 3.2
Write a Go program that converts i to float64 and f to int.

Print out the type and the value of the variables created after conversion.
*/

func main() {
	var i int = 3
	var f float64 = 3.2

	var j = float64(i)
	var g = int(f)

	fmt.Printf("j type: %T value: %v", j, j)
	fmt.Printf("g type: %T value: %v", g, g)
}
