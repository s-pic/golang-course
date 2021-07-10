package main

import "fmt"

/*
Consider the following variable declarations:

x, y, z := 10, 15.5, "Gophers"
score := []int{10, 20, 30}
Using fmt.Printf():

Print each variable using a specific verb for its type

Print the string value enclosed by double quotes ("Gophers")

Print each variable using the same verb

Print the type of y and score
*/

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	// 1
	fmt.Printf("x is %d\n", x)
	fmt.Printf("y is %f\n", y)
	fmt.Printf("z is %s\n", z)
	fmt.Printf("score is %v\n", score)
	// 2
	fmt.Printf("z is %q\n", z)
	// 3
	fmt.Printf("x is %v\n", x)
	fmt.Printf("y is %v\n", y)
	fmt.Printf("z is %v\n", z)
	fmt.Printf("score is %v\n", score)
	// 3
	fmt.Printf("type of y is %T\n", y)

	fmt.Printf("type of score is %T\n", score)

}
