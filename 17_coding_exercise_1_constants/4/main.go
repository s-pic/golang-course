package main

import "fmt"

/*
There are an error in the following Go program. Try to identify the error, change the code and run the program without errors.

package main

func main() {
    const x int = 10

    // declaring a constant of type slice int ([]int)
    const m = []int{1: 3, 4: 5, 6: 8}
    _ = m
}
*/
func main() {
	const x int = 10

	// declaring a constant of type slice int ([]int)
	var m = []int{1: 3, 4: 5, 6: 8} // cannot declare a slice constant
	_ = m

	fmt.Println(m) // -> [0 3 0 0 5 0 8]
	// that syntax above is "new" -> values are stated only for certain indices,
	// gaps are filled with the minimum value of the type
}
