package main

import "fmt"

func main() {

	// intro:
	// 	 go is a statically typed language

	var a = 4
	var b = 5.2

	// a = b <-- will throw, type of a variable is static and thus cannot be changed
	a = int(b)

	fmt.Println(a, b) // prints "5 5.2"

	/*
		core statement:
			in go there is no such thing like an
			uninitialized variable
	*/

	var value int     // initialized with 0
	var price float64 // initialized with 0.0
	var name string   // initialized with ""
	var done bool     // initialized with false

	fmt.Println(value, price, name, done) // prints "0 0  false"
	// go initialized variables to zero values to prevent certain errors

}
