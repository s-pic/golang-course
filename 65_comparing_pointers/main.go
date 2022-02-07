package main

import "fmt"

func main() {
	// we can create a pointer to a pointer

	a := 5.5
	p1 := &a // p1 is a pointer to a, its stores a's addres
	pp1 := &p1

	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n", pp1, &pp1)

	// the value of pp1 is the address of p1

	// how to de-reference a pointer to a pointer?

	fmt.Printf("*p1 is %v\n", *p1)     // 5.5
	fmt.Printf("*pp1 is %v\n", *pp1)   // 0xc00001a0c8
	fmt.Printf("**pp1 is %v\n", **pp1) // 5.5

	**pp1++

	fmt.Println("a is ", a) // 6.5

	// COMPARING POINTERS

	// they are equal if both point to the same value
	y, z := 5, 5

	var p2 *int
	var p3 *int
	p2 = &y
	p3 = &z
	fmt.Println("p2 == p3", p2 == p3) // false
	// both pointers point to different variables (memory addresses)

	p4 := &y

	fmt.Println("p2 == p4", p2 == p4) // true, both point to y

}
