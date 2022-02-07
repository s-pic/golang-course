package main

import (
	"fmt"
)

/*
Create a goroutine named power() that has one parameter of type int, calculates the square value of its parameter and then sends  the result into a channel.

In the main function launch 50 goroutines that calculate the square values of all numbers between 1 and 50 included.

Print out the square values.

A square(or raising to power 2) is the result of multiplying a number by itself. e.g., 25 is the square of 5.
*/

func main() {
	const min = 1
	const max = 50

	ch := make(chan int)
	defer close(ch)

	for counter := min; counter <= max; counter++ {
		go func(i int) {
			sqr := i * i
			ch <- sqr
		}(counter)
	}

	for j := min; j <= max; j++ {
		fmt.Println(<-ch)
	}
}
