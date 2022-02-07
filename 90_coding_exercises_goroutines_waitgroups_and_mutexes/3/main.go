package main

import (
	"fmt"
	"math"
	"sync"
)

/*
1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.

2. Launch the function as a goroutine and synchronize it with main using WaitGroups

Note: You calculate the square root of a float named f using the Sqrt() function from math package like this:
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(f float64) {
		sqrt := math.Sqrt(f)
		fmt.Println(sqrt)
		wg.Done()
	}(9)
	wg.Wait()
}
