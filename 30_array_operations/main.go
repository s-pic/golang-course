package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	// we cannot add or remove elements from the array, since it has a fixed length
	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	//numbers[5] = 2  100 compile time error

	// for loop using *range*
	for idx, val := range numbers {
		fmt.Println("index:", idx, "val:", val)
	}

	fmt.Println(strings.Repeat("#", 20))

	// imperative
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "val:", numbers[i])
	}
	fmt.Println(strings.Repeat("#", 20))

	// multi-dimensional array
	balances := [2][3]int{ // two lines 3 columns
		{5, 6, 7},
		{8, 9, 10},
	}
	fmt.Println(balances)
	fmt.Println(strings.Repeat("#", 20))

	// array equality: same length + same elements in same order
	// note that this is not true for slices!

	m := [3]int{1, 2, 3}
	n := m                                  // two different arrays, no referential equality
	fmt.Println("n is equal to m:", n == m) // prints true

	m[0] = -1
	fmt.Println("n is equal to m:", n == m) // prints false

}
