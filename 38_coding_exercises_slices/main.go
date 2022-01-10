package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1. Using a composite literal declare and initialize a slice of type string with 3 elements
	compositeLiteral := []string{"A", "B", "C"}
	_ = compositeLiteral

	for idx, el := range compositeLiteral {
		fmt.Printf("Element: %s, Index: %d \n", el, idx)
	}

	fmt.Println(strings.Repeat("#", 20))

	/*
		2. There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
		```
		package main

		import "fmt"

		func main() {
			mySlice := []float64{1.2, 5.6}

			mySlice[2] = 6

			a := 10
			mySlice[0] = a

			mySlice[3] = 10.10

			mySlice = append(mySlice, a)

			fmt.Println(mySlice)

		}
		```
	*/

	mySlice := []float64{1.2, 5.6}

	mySlice[1] = 6 // index out of range

	a := 10.0 // types must match
	mySlice[0] = a

	mySlice[1] = 10.10 // index out of range

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)

	fmt.Println(strings.Repeat("#", 20))

	/*
		exercise #3
		1. Declare a slice called nums with three float64 numbers.
		2. Append the value 10.1 to the slice
		3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
		4. Print out the slice
		5. Declare a slice called n with two float64 values
		6. Append n to nums
		7. Print out the nums slice
	*/

	nums := []float64{1.0, 2.1, 3.1}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Printf("%#v\n", nums)

	n := []float64{3.3, 4.4}
	nums = append(nums, n...)
	fmt.Printf("%#v\n", nums)

	fmt.Println(strings.Repeat("#", 20))

}
