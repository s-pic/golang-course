package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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
	printBreak()
	solutionFour()
	printBreak()
	solutionFive()
	printBreak()
	solutionSix()
	printBreak()
}

func solutionSix() {
	/*
		Consider the following string declaration:s := "你好 Go!"

		1. Convert the string to a rune slice.

		2. Print out the rune slice

		3. Iterate over the rune slice and print out each index and rune in the rune slice
	*/

	s := "你好 Go!"

	runeSlice := []rune(s)

	for idx, byte := range runeSlice {
		fmt.Printf("Index %d, Byte: %v\n", idx, byte)
	}
}

func solutionFive() {
	/*
		Consider the following string declaration:s := "你好 Go!"

		1. Convert the string to a byte slice.

		2. Print out the byte slice

		3. Iterate over the byte slice and print out each index and byte in the byte slice

	*/

	s := "你好 Go!"

	byteSlice := []byte(s)

	for idx, byte := range byteSlice {
		fmt.Printf("Index %d, Byte: %v\n", idx, byte)
	}
}

func solutionFour() {
	/*
		There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

		package main

		import "fmt"

		func main() {
			s1 := "Go is cool!"
			x := s1[0]
			fmt.Println(x)

			s1[0] = 'x'

			// printing the number of runes in the string
			fmt.Println(len(s1))
		}
	*/

	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	//s1[0] = 'x' -> index access returns byte at given index, not a rune

	// printing the number of runes in the string
	fmt.Println(utf8.RuneCountInString(s1))
}

func solutionThree() {
	/*
		Consider the following string declaration: s1 := "țară means country in Romanian"

		1. Iterate over the string and print out byte by byte

		2. Iterate  over the string and print out rune by rune and the byte index where the rune starts in the string
	*/

	s1 := "țară means country in Romanian"

	for _, val := range s1 {
		fmt.Println(val)
	}

	// decoding a string rune by rune manually:
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:]) // it returns the rune in string in variable r
		//and its size in bytes in variable size

		// printing out each rune
		fmt.Printf("%c ", r)
		i += size // incrementing i by the size of the rune in bytes
	}
	fmt.Printf("\n")
}

func solutionTwo() {
	/*
		1. Declare a rune called r that stores the non-ascii letter ă

		2. Print out the type of r

		3. Declare 2 strings that contains the values ma and m

		4 and 5. Concatenate the strings and the rune in a new string (the new string will hold the value mamă ).

		BTW: mamă means mother in Romanian.

		Note: You should convert rune to string to contatenate it to another string.
	*/

	r := 'ă'
	fmt.Printf("%T\n", r)

	s1, s2 := "ma", "m"

	var concatenated string = s1 + s2 + string(r)
	fmt.Println(concatenated)
}

func solutionOne() {
	/*
		1. Using the var keyword declare a string called name and initialize it with your name.

		2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.

		3. Print the following string on multiple lines like this:

		Your name: `here the value of name variable`

		Country: `here the value of country variable`

		4 and 5. Print out the following strings:

		a) He says: "Hello"

		b) C:\Users\a.txt
	*/

	var name = "sascha"
	country := "germany"

	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)

	fmt.Println(`a) He says: "Hello"`)
	fmt.Println("b) C:\\Users\\a.txt")

}
