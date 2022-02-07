package main

import (
	"fmt"
	"time"
)

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day) // Time.duration

	seconds := day.Seconds()    // a method (a receiver function) on the type
	fmt.Printf("%T\n", seconds) // float64
	fmt.Printf("Seconds in a day %v\n", seconds)

	// -> named types can have methods (also called receiver functions)
	// attached to them ..

	friends := names{"John", "Silla", "Ben"}
	friends.print() // dot is the selection operator,
	// it selects a name from a namespace (struct, package or type). Here we select a function

	// another way to call a method:

	names.print(friends) // not common, but possible
}

// ..this is how they are declared:

type names []string

func (n names) print() { // n is the function receiver!
	for _, val := range n {
		fmt.Println(val)
	}
}
