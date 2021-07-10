package main

import "fmt"

type km float64
type mile float64

func main() {
	type speed uint
	var s1 speed = 10
	var s2 speed = 20 // uint is still the underlying type

	fmt.Println(s2 - s1) // same type, so arithmetic operation is valid

	var x uint
	//x = s1 // error since different types

	x = uint(s1)
	fmt.Printf("type of x is %T, its value is %v\n", x, x)

	// another example ..

	var s3 speed = speed(x)
	fmt.Println(s3)

	// a conversion from one type to another can be done
	// if both types share the same underlying type

	// another example ..

	var parisToLondon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLondon / 0.621)

	fmt.Println(distanceInMiles)

	// aliases

	var a uint8 = 10
	var b byte

	b = a // byte is an *alias* for uint8, so the assignment is valid
	_ = b

	// we can also declare our own aloases

	type second = uint // note the equal sign

	var hour second = 3600
	fmt.Printf("Minutes per hour: %d\n", hour/60)

}
