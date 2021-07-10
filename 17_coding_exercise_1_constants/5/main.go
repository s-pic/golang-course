package main

import "math"

/*

There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

	const a int = 7
	const b float64 = 5.6
	const c = a * b

	x := 8
	const xc int = x

	const noIPv6 = math.Pow(2, 128)

*/

func main() {
	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b // type mismatch

	x := 8
	var xc int = x // variable has been assigned to constant
	_ = xc

	var noIPv6 = math.Pow(2, 128)
	// same thing as above: We cannot initiate a constant at runtime,
	// but function calls belong to runtime
	_ = noIPv6
}
