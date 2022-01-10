package main

import "fmt"

func main() {
	s1 := "Hi there Go!" // what is between quotes is called string literal
	fmt.Println(s1)

	// raw strings
	fmt.Println("He says: \"Hello\"\n\n..yeah")
	fmt.Println(`He says: "Hello" \n`) // raw string -> no special chars like \ or "

	// printing over multiple lines: the following 2 expressions print the same thing
	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`Price 100:
Brand: Nike
	`)

	// string concatenation
	var s3 = "I love " + "Go"
	fmt.Println(s3 + "!")

	// index access
	fmt.Println(s3[0]) // 73 -> A string is a slice of bytes, here we got the first byte as decimal value

	// a string is immutable
	//s3[5] = x compile error

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3) // with quotes
	fmt.Printf("%T\n", s3)

}
