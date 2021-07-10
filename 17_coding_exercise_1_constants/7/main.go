package main

import "fmt"

/*
Using Iota declare the following months of the year: Jun, Jul and Aug

Jun, Jul and Aug are constant and their value is 6, 7 and 8.
*/

func main() {
	const (
		Jun = iota + 6
		Jul
		Aug
	)
	fmt.Printf("Jun: %v, Jul: %v, ug: %v", Jun, Jul, Aug)
}
