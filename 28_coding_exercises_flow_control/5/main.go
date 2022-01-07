package main

import (
	"fmt"
	"time"
)

/*
Using a for loop print out all the years from your birthday to the current year.
Use a variant of for loop where the post statement is moved inside the for block of code.
*/

func main() {
	for i := 1988; i <= time.Now().Year(); {
		fmt.Println(i)
		i++
	}
}
