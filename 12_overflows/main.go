package main

import (
	"fmt"
	"math"
)

func main() {
	// Overflows and Underflows of integers and floats
	/////////////////////////////

	// If the result of an arithmetic operation has more bits
	// than can be represented in the result type, it is called
	// an overflow. High order bits are then silently discarded

	// For efficiency reason, the go runtime does not check for
	// overflows (just like in C, C++ or Java).

	var x uint8 = 255
	x++

	fmt.Println(x) // equals 0, without an error
	// --> reset to min value for type

	// another example

	var b int8 = 127        // max value
	fmt.Printf("%d\n", b+1) // expression is evaluated at run time
	// -> overflow not detected

	// an overflow can be detected at compile time
	// only if the expression generating the result is evaluated
	// at compile time

	//a := int8(255 + 1)

	// overflows of floats

	f := float32(math.MaxFloat32) // prints 3.4028234663852886e+38

	fmt.Println(f)

	fmt.Println(f * 1.2) // -> +Inf

	// -> floats overflow to infinity !

	// reminder: constants dont overflow!

}
