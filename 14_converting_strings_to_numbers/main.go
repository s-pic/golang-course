package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99)
	fmt.Println(s) // -> c. 99 is the ascii code for c

	//s1 := string(44.2) -> error

	var myStr = fmt.Sprintf("%f", 44.2) // -> calculates string based on stated format and value
	println(myStr)                      // -> 44.200000

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1) // -> 34234

	fmt.Println(string(34234)) // -> 薺, (unicode character for that integer)

	// the package strconv offers conversion methods to and from string as well as other types

	s1 := "3.123"
	fmt.Printf("%T\n", s1) // -> string

	var f1, err = strconv.ParseFloat(s1, 64) // second argument is the precision, ..
	_ = err
	// so a float64 is returned
	fmt.Println(f1 * 2) // -> 6.246

	i, err := strconv.Atoi("-50") // string (asci) to int
	fmt.Printf("i type is %T, i value is %v\n", i, i)
	// -> i type is int, i value is -50

	s2 := strconv.Itoa(20)                                // int to asci
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s2) // q == quoted string
	// -> s2 type is string, s2 value is "20"

}
