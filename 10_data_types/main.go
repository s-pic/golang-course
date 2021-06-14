package main

import "fmt"

func main() {

	// NUMERIC TYPES
	///////////////

	// INT
	var i1, i2 int8 = -128, 127 // min, max
	fmt.Printf("%T, %v\n", i1, i1)
	fmt.Printf("%T, %v\n", i2, i2)

	var j1, j2 uint16 = 0, 65535 // min, max
	fmt.Printf("%T, %v\n", j1, j1)
	fmt.Printf("%T, %v\n", j2, j2)

	// FLOAT

	var f1, f2, f3 float64 = 1.1, -.2, 5. // 1.1, -0.2, 5.0
	fmt.Printf("%T, %v\n", f1, f1)
	fmt.Printf("%T, %v\n", f2, f2)
	fmt.Printf("%T, %v\n", f3, f3)

	// BYTE and RUNE
	// -> both are used to represent characters in go

	var r rune = 'f'
	fmt.Printf("%T, %v\n", r, r) // prints "int32, 102" -> decimal asci code

	////////////////////

	// BOOL

	var b bool = true
	fmt.Printf("%T, %v\n", b, b)

	// STRING

	var s = "Hello Go!"
	_ = s

	// all those above where the basic types in go

	// Composite types
	/////////////////

	// ARRAY

	var numbers = [4]int{4, 5, 9, 100} // length and type must be stated
	fmt.Printf("%T, %v\n", numbers, numbers)

	// SLICE

	// like an array, but can shrink/grow

	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T, %v\n", cities, cities)

	// MAP

	// unordered group of elements of one type indexed by a set of unique keys of another type

	balances := map[string]float64{
		"USD": 2334.2,
		"EUR": 232.3,
	}

	fmt.Printf("%T, %v\n", balances, balances)

	// STRUCT

	// sequence of named elements ("fields"), each fields has a name and a type.
	// can be compared to a class in OOP

	type Car struct {
		brand string
		price int
	}
	var car Car
	fmt.Printf("%T, %v\n", car, car)

	// POINTER

	// stores the memory address of another variable. initialized with nil

	var x int = 2
	ptr := &x // the adress of x
	//ptr2 := &y fails since y is unknown
	fmt.Printf("%T, %v\n", ptr, ptr) // prints "*int, 0xc000130068"

	// FUNCTION
	fmt.Printf("%T, %v\n", f, f) // prints "func(), 0x10a5640"

}
func f() {

}
