package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := []int{2, 3}
	newSlice := append(numbers, 5) // does not mutate
	fmt.Println("newSlice", newSlice)

	newSlice = append(newSlice, 45, 56) // variadic function
	fmt.Println("newSlice", newSlice)

	n := []int{100, 200}
	numbers = append(numbers, n...) // ellipsis operator
	fmt.Println("numbers", numbers)

	fmt.Println(strings.Repeat("#", 20))

	// copy slice to another slice
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)

	fmt.Println("number of copied elements", nn)
	fmt.Println("src", src)
	fmt.Println("dst", dst)

	// copies until destination slice is full
}
