package main

import "fmt"

func main() {
	// similar to slices:
	// when creating a Map,
	// Go returns a pointer to a Map header in Memory

	// The key value pairs are stored at the address referenced by the map header

	// This must be payed attention to when copying maps

	friends := map[string]int{"Dan": 40, "Maria": 25}
	neighbors := friends // shared reference
	neighbors["Dan"] = 100
	fmt.Println(friends) // mutated

	// how to clone properly
	people := make(map[string]int)
	for k, v := range friends {
		people[k] = v
	}

}
