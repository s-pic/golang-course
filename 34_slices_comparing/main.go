package main

import "fmt"

func main() {
	var n []int // nil
	_ = n

	m := []int{} // empty but initialized
	_ = m

	// slices can only be compared to nil

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	//fmt.Println(a == b) // Runtime error. invalid operation: a == b (slice can only be compared to nil)

	// we need to iterate and compare value by value
	var eq bool = len(a) == len(b)
	if eq {
		for idx, valA := range a {
			valB := b[idx]
			if valA != valB {
				eq = false
				break
			}
		}
	}
	fmt.Println("eq:", eq)
}
