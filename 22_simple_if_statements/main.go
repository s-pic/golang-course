package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
	Simple/Short statements are statements
	which are allowed in another (If, Switch) Statements.
	Variables declared in a simple statement are scoped
	to the branch of that statement.
	Most often they are used to handle errors.

*/

func main() {
	i, err := strconv.Atoi("45")
	if err != nil {
		fmt.Println(err) // --> strconv.Atoi: parsing "45s": invalid syntax
	} else {
		fmt.Println(i) // -> 45
	}

	// no the same with the Simple syntax

	// the part after the if is called "Initialization Statement"
	// the second part is the Boolean expression determining
	// which IF-Else branch is selected
	if j, err := strconv.Atoi("45"); err == nil {
		fmt.Println("No error, j is ", j)
	} else {
		fmt.Println("There is an error:", err)
	}

	// fmt.Println(j) --> error, since j is scoped to the branches
	// of the If statement

	// second example

	if args := os.Args; len(args) != 2 {
		fmt.Println("Please pass kilometers as argument")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Please pass an integer value")
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
	}
}
