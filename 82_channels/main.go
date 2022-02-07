package main

import "fmt"

func main() {
	// channels are used to communicate between goroutines

	var ch chan int // a channel sending int values
	fmt.Println(ch) // nil since uninitialized

	ch = make(chan int) // make is built in
	fmt.Println(ch)     // stores a memory address. indeed a channel works similar to a pointer

	c := make(chan int) // another channel

	// <- channel operator. The following statement sends a value 10 into the channel

	// SEND

	c <- 10

	// RECEIVE

	num := <-c // we wait for the channel to send a value and assign it
	_ = num

	fmt.Println(<-c) // or just use the value directly in a function call

	// CLOSE A CHANNEL (sending won't be possible after that)
	close(c) // fatal error: no goroutine used
}
