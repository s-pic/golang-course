package main

import "fmt"

func main() {
	// arrays, slices and strings are slicable

	a := [5]int{1, 2, 3, 4, 5}
	slice := a[0:3] // like [].slice in js
	fmt.Printf("%v, %T\n", slice, slice)

	// -> returns a slice, not an array

	//fmt.Println(a[1:10]) runtime error: out of bounds for 5-element array

	fmt.Println("A string"[2:]) // returns 'string'
	fmt.Println("A string"[:3]) // returns A s

	// if start index is omitted -> defaults to 0
	// if end index is omitted -> defaults to len

	s1 := []int{1, 2, 3, 4, 5, 6, 7}

	s1 = append(s1[:3], 100)
	fmt.Println(s1)
}
