package main

import "fmt"

/*
Declare a new type type called duration. Have the underlying type be an int.

Declare a variable of the new type called hour using the var keyword

In function main:

print out the value of the variable hour


print out the type of the variable hour

assign 3600 to the variable hour using the  = operator

print out the value of hour
*/

type duration int

var hour duration = 60

func main() {
	fmt.Printf("hour: value %d, type %T\n", hour, hour)
	hour = 3600
	fmt.Println(hour)
}
