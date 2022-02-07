package main

import "fmt"

func main() {
	name := "Andrei!"
	fmt.Println(&name) // prints memory address
	// & is an operator that can be also used with composite literals like
	// those for slices

	var x int = 2
	ptr := &x // if & -- when used in assignment expression --
	// this "Address Operator" creates a pointer

	fmt.Printf("Type %T, value%v\n", ptr, ptr)
	// Type *int, value0xc00011a010

	//  *int is pronounced "pointer to int"

	// we can print an address like so
	fmt.Printf("address of x is %p\n", &x)

	// a pointer has a value (the memory address of the value it is pointing to),
	// but also an own address:
	fmt.Printf("ptr is of type %T with a value of %v and address %p\n", ptr, ptr, &ptr)

	// we can declare a pointer without initializing it
	var ptr1 *float64
	fmt.Println(ptr1) // nil

	// another way to create a pointer
	p := new(int)

	x = 100
	p = &x

	// DE-REFERENCING-OPERATOR

	*p = 90 // equivalent to x = 90. *p means "the variable to which points"
	// or "the underlying value"
	fmt.Println(x, *p)

	// the star operator (de-referencing operator) says
	// take this memory address and get me the value at that address
	// -> the value of that variable

	fmt.Println("*p==x:", *p == x)

	// a star in front of a pointer and a start in front of a type mean
	// two different things!
	// first case means: create a pointer to a value of that type,
	// second case means: get me the value the pointer points to

	*p = 10 // -> same as x = 10
	*p = x / 2
	fmt.Println(x) // 5

	// key take-away:

	// turn a value into a pointer:
	// &value => pointer

	// turn a pointer into a value\
	// *pointer => value
}
