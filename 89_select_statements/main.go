package main

import (
	"fmt"
	"time"
)

func main() {
	// the select statements lets a groutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case.

	// It this allows us to wait on multiple channel operations.

	// Select is only used with channels

	c1 := make(chan string)
	c2 := make(chan string)

	start := time.Now().UnixNano() / 1000000

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Salut"
	}()

	// use select to await both channel events

	for i := 0; i < 2; i++ {
		select { // just like a switch, but only used with channels
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}

	end := time.Now().UnixNano() / 1000000

	fmt.Printf("Passed time in seconds %.2f", float64((end-start)/1000))
}
