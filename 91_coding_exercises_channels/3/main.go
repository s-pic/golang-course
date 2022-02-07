package main

import "fmt"

/*
package main

import (
    "fmt"
)

func main() {
    c := make(<-chan int)

    go func(n int) {
        c <- n
    }(100)

    fmt.Println(<-c)
}

*/
func main() {
	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}
