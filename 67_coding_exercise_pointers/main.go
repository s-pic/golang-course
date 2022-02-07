package main

import "fmt"

func main() {
	exOne()
	exTwo()
	exThree()
}

func exOne() {
	/*
		Consider the following variable declaration x := 10.10

		1. Print out the address of x

		2. Declare a pointer called ptr that stores the address of x.

		3. Print out the type and the value of ptr.

		4 and 5. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
	*/

	x := 10.10
	fmt.Printf("Memory address of x is %p\n", &x)
	ptr := &x
	fmt.Printf("ptr: Type %T, value:%v\n", ptr, ptr)
	fmt.Printf("ptr: Adress %p\n", &ptr)
	fmt.Printf("Value of x: %v\n", *ptr)

}

func exTwo() {
	/*
		Consider the following variable declarations:

		x, y := 10, 2
		ptrx, ptry := &x, &y
		Declare a new variable called z and initialize the variable by dividing x by y through the pointers.
	*/
	x, y := 10, 2
	ptrx, ptry := &x, &y

	z := *ptrx / *ptry
	fmt.Println("z:", z)

}

func exThree() {
	/*
		Consider the following variable declaration:x, y := 5.5, 8.8

		Write a function that swaps the values of x and y. After calling the function x will be 8.8 and y will 5.5
	*/

	x, y := 5.5, 8.8
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y)
}

func swap(a, b *float64) {
	*a, *b = *b, *a
}
