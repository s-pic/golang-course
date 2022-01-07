package main

import "fmt"

func main() {
	var cities []string // uninitialized -> nil
	fmt.Println("cities is equal to nil:", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities)) // 0

	//cities[0] = "London" runtime error: index out of range
	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	// another way to create a slice
	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums) // two elements initialized with 0
	nums[0] = 1
	fmt.Printf("%#v\n", nums) // now the index access works

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My best friend is:", myFriend)

	for idx, value := range friends {
		fmt.Printf("index: %d, value: %v\n", idx, value)
	}

	var n []int
	n = numbers
	fmt.Println(n)
}
