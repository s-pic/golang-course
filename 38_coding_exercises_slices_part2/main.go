package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at command line.

The user should give between 2 and 10 numbers.

Expected output: Sum: 10, Product: 30

*/
func main() {
	nums := os.Args[1:]

	if len(nums) < 2 || len(nums) > 10 {
		err := errors.New("Expected between 2 and 10 numbers")
		log.Fatal(err)
	}
	fmt.Println("Received numbers ", nums)

	var sum float64 = 0
	var product float64 = 1

	for _, numAsStr := range nums {
		num, err := strconv.ParseFloat(numAsStr, 32)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
		product *= num
	}

	fmt.Printf("Sum: %.1f, Product: %.1f", sum, product)

	fmt.Println(strings.Repeat("#", 20))
	runExerciseFive()

	fmt.Println(strings.Repeat("#", 20))
	runExerciseSix()

	fmt.Println(strings.Repeat("#", 20))
	runExerciseSeven()

	fmt.Println(strings.Repeat("#", 20))
	runExerciseEight()
}

func runExerciseEight() {
	/*
		Consider the following slice declaration:
		 years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
		Using a slice expression and append() function create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.  newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}
	*/
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}

	newYears = append(years[:3], years[len(years)-3:]...)

	fmt.Printf("%#v\n", newYears)
}

func runExerciseSeven() {
	/*
	   	Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
	   Using append() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.*
	*/

	friends := []string{"Marry", "John", "Paul", "Diana"}

	var friendsCopy []string
	friendsCopy = append(friendsCopy, friends...)

	friendsCopy[0] = "brbebeberb"

	fmt.Println("friends", friends)
	fmt.Println("friendsCopy", friendsCopy)
}

func runExerciseSix() {
	/*
		Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

		Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	*/

	friends := []string{"Marry", "John", "Paul", "Diana"}

	friendsCopy := make([]string, len(friends))
	copy(friendsCopy, friends)

	friendsCopy[0] = "brbebeberb"

	fmt.Println("friends", friends)
	fmt.Println("friendsCopy", friendsCopy)

}

func runExerciseFive() {
	/*
		Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
		Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
		Print those elements and their sum.
	*/

	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	sum := 0
	for _, val := range nums[1 : len(nums)-2] {
		fmt.Println(val)
		sum += val
	}

}
