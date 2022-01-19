package main

import "fmt"

func main() {
	diana := struct {
		firstName, lastName string
		age                 uint16 ``
	}{
		firstName: "Diana",
		lastName:  "Mueller",
		age:       12,
	}
	fmt.Printf("%+v\n", diana)

	// -> we created the struct without defining
	// the struct type alias.
	// Useful when we dont want to re-use the types

	// ANONYMOUS FIELDS
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{
		string:  "2343",
		float64: 0,
		bool:    false,
	}
	fmt.Printf("%+v\n", b1)

	// we did not specify field names, go automatically created
	// fields named after their type.

	// not sure why that would be useful though
}
