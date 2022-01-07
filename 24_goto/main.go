package main

import "fmt"

func main() {
	// goto statements allow us to jump to specific labels

	// using these statements are discouraged

	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}
}
