package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devine comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "Shakespeare", 1606

	fmt.Println("Book1", title1, author1, year1)
	fmt.Println("Book2", title2, author2, year2)

	// better approach to treat related values: structs

	// CREATING STRUCTS

	// 1. step: create type with field names and types

	type book struct {
		title  string
		author string
		// we could also write 		title, author string
		year uint32
	}

	// 2. step create the struct value (similar to instantiating a class
	// in other languages

	myBook := book{
		title:  title1,
		author: author1,
		// if a field is omitted, it is initialized with its zero value
	}

	fmt.Printf("%+v\n", myBook) // that plus modifier
	// prints out field and values

	// we can also assign values by stating field values in order
	// (not the recommended way)

	myOtherBook := book{title2, author2, 1606}
	fmt.Println(myOtherBook)

	// READING AND WRITING

	myOtherBook.title = "Anna Karenina"
	fmt.Println(myOtherBook.title)

	// COMPARING

	// Structs are comparable, using ++

	type One struct {
		a string
	}

	a := One{a: "val"}
	b := One{a: "val"}
	fmt.Println(a == b)               // true
	fmt.Println(a == One{a: author2}) // false

	// COPYING

	// in contrast to arrays or slices, we can duplicate a struct
	// simply using =

	c := One{a: "val"}
	d := c

	d.a = "modified"

	fmt.Println(c.a) // val -> not mutated
}
