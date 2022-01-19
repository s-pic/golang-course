package main

import (
	"fmt"
	"strings"
)

func printBreak() {
	line := strings.Repeat("#", 20)
	fmt.Println(line)
}

func main() {
	solutionOne()
	printBreak()
	solutionTwo()
	printBreak()
	solutionThree()
}

func solutionOne() {
	/*
	   	1. Declare a map called m1 which value is nil. Print out its type and value.

	      2. Declare a map called m2. It's keys are of type int and values of type string.  Initialize the map using  a map literal with two key:value pairs.

	      3. Add the following key: value to the map: 10: "Abba"

	      4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.

	*/

	var m1 map[string]string
	fmt.Printf("Type of m1 is %T,  \nits value is %#v\n", m1, m1)

	m2 := map[int]string{
		1: "X",
		2: "Y",
	}
	m2[10] = "Abba"

	entry1, ok := m2[10]
	if ok {
		fmt.Println(entry1)
	} else {
		fmt.Println("no value at key 10")
	}

	entry2, ok := m2[6]
	if ok {
		fmt.Println(entry2)
	} else {
		fmt.Println("no value at key 6")
	}
}

func solutionTwo() {
	/*
		There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

		package main

		import "fmt"

		func main() {
			var m1 map[int]bool
			m1[5] = true

			m2 := map[int]int{3: 10, 4: 40}
			m3 := map[int]int{3: 10, 4: 40}

			fmt.Println(m2 == m3)
		}
	*/

	//var m1 map[int]bool Must be initialized in order to add values
	var m1 = make(map[int]bool)
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	//fmt.Println(m2 == m3) // Maps can only be compared to nil

	fmt.Println(
		fmt.Sprintf("%s", m2) == fmt.Sprintf("%s", m3),
	)
}

func solutionThree() {
	/*
		Consider the following map declaration: m := map[int]bool{"1": true, 2: false, 3: false}

		1. The above map declaration has an error. Solve the error!

		2. Delete a key:value pair from the map.

		3. Iterate over the map and print out each key and value.
	*/

	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 1)

	for k, v := range m {
		fmt.Println("key:", k, "val: ", v)
	}
}
