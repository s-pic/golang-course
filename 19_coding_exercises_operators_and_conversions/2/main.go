package main

import (
	"fmt"
	"strconv"
)

/*
Consider the following declarations:

var i = 3
var f = 3.2
var s1, s2 = "3.14", "5"
Write a Go program that converts:

1. i to string (int to string)

2. s2 to int (string to int)

3. f to string (float64 to string)

4 and 5. s1 to float64  (string to float64)

6. Print the value and the type for each variable created after conversion.
*/
func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	var iConverted = strconv.Itoa(i)
	fmt.Printf("iConverted type: %T value %q\n", iConverted, iConverted)

	var s2Converted, atoiErr = strconv.Atoi(s2)
	if atoiErr != nil {
		fmt.Println(atoiErr)
		return
	}
	fmt.Printf("s2Converted type: %T value %d\n", s2Converted, s2Converted)

	var fConverted = fmt.Sprintf("%f", f)
	fmt.Printf("fConverted type: %T value %q\n", fConverted, fConverted)

	var s1Converted, parseFloatErr = strconv.ParseFloat(s1, 64)
	if parseFloatErr != nil {
		fmt.Println(parseFloatErr)
		return
	}
	fmt.Printf("s1Converted type: %T value %f\n", s1Converted, s1Converted)

}
