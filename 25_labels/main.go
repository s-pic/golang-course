package main

import "fmt"

func main() {
	people := [5]string{"Hellen", "Mark", "Brenda", "Antonion", "Michael"}
	friends := [2]string{"Mark", "May"}

outer: // that is the label for the *outer loop*
	for index, name := range people { // iterate over array
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after loop")

	outer := "not conflicting with label"
	// labels live in another space,
	// they dont conflict with other names
	_ = outer
}
