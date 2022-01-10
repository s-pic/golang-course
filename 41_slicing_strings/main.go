package main

import "fmt"

func main() {
	// slicing a string re-uses the same backing array

	s1 := "I love Golang"
	fmt.Println(s1[2:5])

	// slicing a string returns (the unicode representations of)
	// bytes, not runes!

	nonAsciiChars := "ƗØƀĐǤĦŁŦƵ"
	fmt.Println(nonAsciiChars[0:1]) // prints �

	// -> only one byte returned, which is not enough to represent the first rune

	// now how to return the first x runes?

	// 1. convert the string to a slice of runes
	rs := []rune(nonAsciiChars)
	fmt.Printf("Type %T, Value %#v\n", rs, rs)

	// convert back to string
	slicedSlice := rs[0:3]
	convertedString := string(slicedSlice)
	fmt.Println(convertedString) // ƗØƀ

}
