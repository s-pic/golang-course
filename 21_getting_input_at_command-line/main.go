package main

import (
	"fmt"
	"os"
	"strconv"
)

// standard library package "os" provides
// operating system functionalities

func main() {
	fmt.Println("os.Args", os.Args)
	// prints [<Path-to-the-go-program>]
	// -> os.Args is a slice of strings

	// if we run the go program by passing a command line arg like so ..
	// `go run main.go linux macos windows`
	// .. the passed arguments are added after the path

	// so we can access command line args using indices

	fmt.Println("Path: ", os.Args[0])
	fmt.Println("First Arg: ", os.Args[1])
	fmt.Println("Second Arg: ", os.Args[2])
	fmt.Println("Third Arg: ", os.Args[3])
	fmt.Println("Number of Args", len(os.Args))
	/* -->
	First Arg:  linux
	Second Arg:  mac
	Second Arg:  windows
	Number of Args 5
	*/

	// another example, where we pass an argument
	// (which is always a string) to a float

	var result, err = strconv.ParseFloat(os.Args[4], 64)
	fmt.Println("result", result)
	fmt.Printf("%T\n", os.Args[4]) // -> string
	fmt.Printf("%T\n", result)     // -> float 64
	_ = err

	// run `go run main.go linux mac windows 4535.44`
}
