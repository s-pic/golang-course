package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("I am anon") // must be invoked now, since there is no name to refer it to later

	// use case: when we want to return a function or pass a function as arg
	// -> Go supports treats functions as first class citizens
	doIncrement := increment(10)
	fmt.Println(doIncrement()) // 11
	fmt.Println(doIncrement()) // 12
}
