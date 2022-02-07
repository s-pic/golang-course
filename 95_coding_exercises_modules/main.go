package main

import (
	"fmt"
	"mypackages/arithmetic"
)

func main() {
	result := arithmetic.IsPrime(13) == true

	fmt.Println("13 is a prime number", result)
	fmt.Println("Factorial of 7 is", arithmetic.Factorial(7))
}
