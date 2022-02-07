package main

import (
	"fmt"
	"strings"
)

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	c <- f
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)
	// the main go-routine sleeps (blocks)
	// until it receives a value via the channel
	f := <-ch
	fmt.Println(f)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}

	fmt.Println(strings.Repeat("@", 20))

	// Spawning another 10 goroutines this time as anonymous functions
	for i := 5; i < 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}

			// sending the value f into the channel
			c <- f
		}(i, ch)
		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}

	fmt.Println("Main exited")
}
