package main

import (
	"fmt"
	"strings"
)

func main() {
	exOne()
	fmt.Println(strings.Repeat("#", 20))
	exTwo()
	fmt.Println(strings.Repeat("#", 20))
	exThree()
	fmt.Println(strings.Repeat("#", 20))
	exFour()
}

func exFour() {
	type ColorList []string
	type Grades struct {
		grade  uint8
		course string
	}
	type Person struct {
		name   string
		age    uint8
		colors ColorList
		grades Grades
	}
	me := Person{
		name:   "sascha",
		age:    33,
		colors: ColorList{"red, green"},
		grades: Grades{
			grade:  5,
			course: "golang",
		},
	}
	fmt.Printf("%+v\n", me)
}

func exThree() {
	/*
		Consider the code from Exercise #1.
		Iterate and print out the favorite colors of the struct value called me.
		Are you stuck? Do you want to see the solution for this exercise? Click here.
	*/

	type ColorList []string
	type Person struct {
		name   string
		age    uint8
		colors ColorList
	}
	me := Person{
		name:   "sascha",
		age:    33,
		colors: ColorList{"red, green"},
	}

	for _, val := range me.colors {
		fmt.Println(val)
	}
}

func exTwo() {
	/*
		1. Change the name or the struct value called me to "Andrei"
		2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
		3. Add a new favorite color to the second struct value called you.
		4 and 5. Print out the struct values.
	*/
	type ColorList []string
	type Person struct {
		name   string
		age    uint8
		colors ColorList
	}
	me := Person{
		name:   "sascha",
		age:    33,
		colors: ColorList{"red, green"},
	}
	me.name = "Andrei"

	you := Person{
		name:   "foo",
		age:    22,
		colors: ColorList{"yellow", "pink"},
	}
	favColors := you.colors
	fmt.Printf("Type of favColors: %T. Value: %v\n", favColors, favColors)

	you.colors = append(you.colors, "brown")
	fmt.Println(you.colors)
}

func exOne() {
	/*
		1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
		2. Declare and initialize two values of type person, one called me and another called you.
		3. Print out the struct values.
	*/

	type ColorList []string
	type Person struct {
		name   string
		age    uint8
		colors ColorList
	}
	me := Person{
		name:   "sascha",
		age:    33,
		colors: ColorList{"red, green"},
	}
	you := Person{
		name:   "foo",
		age:    22,
		colors: nil,
	}

	fmt.Printf("me: %+v\n", me)
	fmt.Printf("me: %+v\n", you)
}
