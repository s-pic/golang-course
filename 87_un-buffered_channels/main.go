package main

import (
	"fmt"
	"time"
)

func main() {
	// go has two types of channes with different behaviour: buffered an unbuffered ones

	// so far we have only seen unbuffered ones

	unbufferedChannel := make(chan int)

	bufferedChannel := make(
		chan int,
		3, // capacity arg
	)
	_ = bufferedChannel

	go func(c chan int) {
		fmt.Println("goroutine starts")
		c <- 10 // this blocks the goroutine until main wakes up
		fmt.Println("goroutine exits")

	}(unbufferedChannel)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")

	d := <-unbufferedChannel
	fmt.Println("main goroutine received data:", d)

	time.Sleep(time.Second)

	// After running the program we notice that the sender (the func goroutine) blocks on the channel
	// until the receiver (the main goroutine) receives the data from the channel.

}
