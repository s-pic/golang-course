package main

import "fmt"

// all functions, constants and variables declared outside a function
// are visible to all files of the same package

func sayHello(name string) {
	fmt.Printf("Hello %s! \n", name)
}

const BOILING_POINT_IN_DEG = 100

func celsiusToFahrenHeit(c float64) float64 {
	return c*1.8 + 32
}
