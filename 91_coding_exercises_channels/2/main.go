package main

import "fmt"

func main() {
	ch := make(chan string)

	go func(str string, c *chan string) {
		(*c) <- str
	}("foo", &ch)

	fmt.Println(<-ch)
}
