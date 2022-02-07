package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// data races (or race conditions) in go happen because of un-synchronized access
	// to shared memory.

	// Lets see how to cause one

	const gr = 100 // 100 go routines will increment the same shared value
	// and 100 will decrement it at the same time

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	for i := 0; i < gr; i++ {
		go func() { // use anonymous function literal
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() { // same logic, but with decrement
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("n is", n) // close to 0, but RANDOM!!

	// The value of n depends on the order in which goroutines finish
	// ..and that is a data race!

	// go helps detecting them by running the program with
	// -race:
	//  go run -race main.go
}
