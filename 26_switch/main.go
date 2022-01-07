package main

import (
	"fmt"
	"time"
)

func main() {
	language := "Golang"

	switch language {
	case "Python":
		fmt.Println("you are learning Python")
		// no break needed, go adds it automatically
	case "Go", "Golang": // works like *OR*
		fmt.Println("good to go")
	default:
		fmt.Println("any other lang works also")
	}

	n := 5
	switch true { // cases can only do comparisons with expressions
	// that evaluate to a bool
	case n%2 == 0:
		fmt.Println("Even num")
	case n%2 != 0:
		fmt.Println("Odd number")
	}

	hour := time.Now().Hour()
	fmt.Println(hour)

	switch { // missing expression means *true*
	case hour < 12:
		fmt.Println("good morning")
	case hour < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("Good evening")

	}
}
