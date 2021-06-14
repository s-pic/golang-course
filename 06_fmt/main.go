package main

import "fmt"

func main() {
	// print concatenated output using Println

	name := "sascha"
	fmt.Println("Hello, my name is ", name)

	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Mean Value:", (a+b)/2) // adds whitespace between each arg and a /n at the end

	// print formatted output
	fmt.Printf("Your age is %d\n", 21) // newline must be added manually, in contrast to Println
	fmt.Printf("x is %d, y is %f", 5, 6.8)

	fmt.Printf("He says: \"Hello go!\"\n") // special chars like quotes must be excaped

	figure, radius := "Circle", 5

	fmt.Printf("The diameter of a %s with a Radius is %d\n", figure, radius)

	fmt.Printf("This is a %q\n", figure) // quoted strings, useful!

	// "swiss-army-placeholder" %v

	fmt.Printf("The diameter of a %v with a Radius is %v\n", figure, radius)

	// %T -> type

	fmt.Printf("figure is of type %T, Radius is of type %T\n", figure, radius)

	// %t -> boolean

	fileClosed := true
	fmt.Printf("File closed: %t\n", fileClosed)

	// %b -> base 2

	fmt.Printf("16 in bytes is %b \n", 4)

	// control shown number of decimal points

	x := 3.4
	y := 6.9

	fmt.Printf("x * y = %.3f\n", x*y) // use only 3 decimal points

}
