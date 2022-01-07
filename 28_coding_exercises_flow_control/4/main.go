package main

import "fmt"

/*
Using a for loop, an if statement and the logical and
operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.
*/

func main() {
	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}
}
