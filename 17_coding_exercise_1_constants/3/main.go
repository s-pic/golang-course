package main

import "fmt"

/*

Calculate how many seconds are in a year.

 STEPS:

 1. Declare secPerDay constant and initialize it to the number of seconds in a day

 2. Declare daysYear constant and initialize it to 365

 3. Use fmt.Printf() to print out the total number of seconds in a year.

EXPECTED OUTPUT:

There are 31536000 seconds in a year.

*/
func main() {
	const secPerDay = 24 * 60 * 60
	const daysPerYear = 365

	fmt.Printf("There are %d seconds in a year\n", secPerDay*daysPerYear)
}
