package main

import "fmt"

// a goroutine that sends values into the channel

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	c := make(chan int)

	// btw: there are also uni-directional channels

	// only for receiving
	c1 := make(<-chan string)
	// only for sending
	c2 := make(chan<- string)

	fmt.Println("The types are")
	fmt.Printf("c: %T\n", c)
	fmt.Printf("c1: %T\n", c1)
	fmt.Printf("c2: %T\n", c2)

	go f1(10, c) // pass the bidirectional channel

	n := <-c

	fmt.Println("Received value is", n)
	fmt.Println("Exiting main")
}
