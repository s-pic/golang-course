package main

import "fmt"

var numbers [10]int // package scoped, initialized with 0 value

// entry point for an executable application, automatically called
func main() {
	fmt.Println("this is main")
	// init() cannot be called explicitly

	fmt.Printf("numbers: %#v", numbers)
}

// also automatically called, but BEFORE main
func init() {
	fmt.Println("this is init")

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
}

// we can have multiple init files, executed in order of declaration
func init() {
	fmt.Println("this is another init file")
}
