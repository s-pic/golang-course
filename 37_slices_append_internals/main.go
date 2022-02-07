package main

import (
	"fmt"
	"strings"
)

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))
	// 0, 0

	nums = append(nums, 1, 2)

	// when the slice capacity is full (when the length of the backing array is not sufficient),
	// go creates a new backing array

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))
	// 3, 4 and 5

	nums = append(nums, 4)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	// creating new backing arrays is expensive,
	// so go increases the length of the new backing array exponentially
	// each time it "runs full"

	// just check the logs (printed statements)

	nums = append(nums, 6, 7, 8, 9) // 9 elements

	// capacity (length of backing array) increases to 16:
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	fmt.Println(strings.Repeat("#", 20))

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")

	fmt.Printf("Length: %d, Capacity: %d \n", len(letters), cap(letters))

	// 3, 6 -> capacity is still 6, the same backing array is used

}
