package main

import (
	"fmt"
	"sync"
)

/*
1. Create a function called sum() that calculates and then prints out  the sum of 2 float numbers it receives as arguments.

Format the result with 2 decimal points.

2. From main launch 3 goroutines that execute the function you have just created (sum)

3. Synchronize the goroutines and the main function using WaitGroups
*/

func sum(a, b float64, wg *sync.WaitGroup) {
	fmt.Printf("Sum: %.2f\n", a+b)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go sum(11.4, 53.3, &wg)
	go sum(298.37, -34.3343, &wg)
	go sum(0, -1, &wg)
	wg.Wait()
}
