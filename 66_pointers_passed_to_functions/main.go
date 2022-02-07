package main

import "fmt"

// in go there is no "pass by ref" https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
// but we can differentiate between
// pass by value and
// pass bu pointer
// still, the variable points to a memory address, so we can change the value under
// that address
func change(a *int) *float64 { // a is created as copy, a variable local to the function.
	*a = 100 // changes passed pointer -> affects outside world

	b := 5.5

	return &b // valid go code
}

func changeVar(a int) {
	a = 66 // change scoped to the function. same is true for floats, strings and bools
}

type Product struct {
	price float64
	name  string
}

func changeStruct(product Product) { // just changes the local copy of the struct
	product.price = 33
	product.name = "Byclicle"
}

func changeStructByPointer(product *Product) { // just changes the local copy of the struct
	(*product).price = 33 // both notations are equal
	product.name = "Byclicle"
}

func changeSlice(sl []int) {
	for i := range sl {
		sl[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
}

func main() {
	x := 8
	p := &x

	fmt.Println("Val of x *before* calling change()", x)

	change(p)

	fmt.Println("Val of x *after* calling change()", x)

	changeVar(x)

	fmt.Println("Val of x *after* calling changeVar()", x)

	gift := Product{
		price: 100,
		name:  "Watch",
	}
	changeStruct(gift)
	fmt.Println("Gift", gift)

	changeStructByPointer(&gift)

	fmt.Println("Gift after", gift)

	// TLDR: ints, floats, bools, strings, arrays and structs can only be changed
	// by a function, when passed as pointers.
	// THIS IS DIFFERENT FOR SLICES AND MAPS!

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println("Prices after calling func", prices)

	someMap := map[string]int{
		"x": 34,
	}
	changeMap(someMap)
	fmt.Println("someMap after calling func", someMap)

	// SLICES AND MAPS ARE ALREADY POINTERS TO AN UNDERLYING ARRAY

}
