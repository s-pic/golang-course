package main

import "fmt"

func main() {

	// first way to declare a variable: var

	var age int = 30 // initial value is optional
	var age2 = 30    // type can also be inferred und thus omitted
	fmt.Println("Age:", age)
	fmt.Println("Age2:", age2)

	// unused variables will lead to a compile error.

	// If we want to declare a variable and *NOT* use it,
	// we can use the blank identifier that tell the compiler to
	// to mute the compile error. It should be used with caution
	// It always stays left of the equal sign
	var name = "Dan"
	_ = name

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender) // prints "0  false "
	// those are go default values.
	// All variables in go are initialized!

	var a, b, c int
	fmt.Println(a, b, c) // prints "0 0 0"

	var i, j int
	i, j = 5, 8
	_, _ = i, j

	j, i = i, j       // swap variables using multi assignment
	fmt.Println(i, j) // prints "8, 6"

	// second way to declare a variable: short declaration operator

	// only works in block scope (only inside main or another function)

	s := "Learning golang!"
	s = "Another String"
	fmt.Println(s) // prints "Another String"

	car, cost := "Audi", 50000
	fmt.Println(car, cost) // prints "Audi 50000"
	car, cost = "BWM", 6000
	fmt.Println(car, cost) // "BWM 6000"

	car, newVariable := "Mercedes", "something"
	_ = newVariable

	sum := cost + 40
	println(sum)
}
