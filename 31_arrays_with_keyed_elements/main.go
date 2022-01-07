package main

import "fmt"

func main() {
	grades := [3]int{
		// we can state indices at any order
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) // prints [5 10 7]

	accounts := [3]int{2: 50} // only last element stated
	fmt.Println(accounts)     // prints [0,0,50]

	names := [...]string{
		5: "Dan",
	}
	fmt.Println(names, len(names)) // prints [     Dan] 6

	cities := [...]string{
		5: "Paris",
		// an unkeyed element gets its key from the last keyed element
		"London",
		1: "NYC",
	}
	fmt.Printf("%#v\n", cities) // prints [7]string{"", "NYC", "", "", "", "Paris", "London"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // [false false false false false true true]

}
