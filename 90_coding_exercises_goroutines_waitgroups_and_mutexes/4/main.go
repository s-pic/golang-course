package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Change the code from Exercise #3 and launch 50 goroutines that calculate
concurrently the square root of all the numbers between 100 and 149 (both included).
*/
func main() {
	const numOfRoutines = 50
	var wg sync.WaitGroup
	wg.Add(numOfRoutines)
	for i := 100; i < 100+numOfRoutines; i++ {
		go func(f float64) {
			sqrt := math.Sqrt(f)
			fmt.Println(sqrt)
			wg.Done()
		}(float64(i))
	}

	wg.Wait()
}
