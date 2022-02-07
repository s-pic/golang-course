package main

import "fmt"

/*
1. Using the var keyword, declare a bidirectional unbuffered channel called c1 that works with values of type float64

2. Using the make() built-in function declare and initialize a receive-only channel called c2 and a send-only channel called c3. Both work with data of type rune.

3. Declare a bidirectional buffered channel  called c4 with a capacity of 10 ints.

4 and 5. Print out the type of all the channels declared.
*/
func main() {
	var c1 = make(chan float64)
	c2 := make(chan<- rune)
	c3 := make(<-chan rune)
	c4 := make(chan int, 10)

	fmt.Printf("Type of c1 is %T\n", c1)
	fmt.Printf("Type of c2 is %T\n", c2)
	fmt.Printf("Type of c3 is %T\n", c3)
	fmt.Printf("Type of c4 is %T\n", c4)
}
