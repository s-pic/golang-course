package main

import "fmt"

/*
Declare two defined types called mile and kilometer. Have the underlying type be an float64.

Declare a constant called m2km equals 1.609  (1mile=1.609km)

In function main:

declare a variable called mileBerlinToParis of type mile with value 655.3

declare a variable called kmBerlinToParis of type kilometer

calculate the distance between Berlin and Paris in km by multiplying mileBerlinToParis and m2km. Assign the result to kmBerlinToParis

print out the distance in km between Berlin and Paris
*/

type mile float64

type kilometer float64

const m2km = 1.609

func main() {
	var mileBerlinToParis = 655.3
	var kmBerlinToParis = mileBerlinToParis * m2km

	fmt.Println(kmBerlinToParis)
}
