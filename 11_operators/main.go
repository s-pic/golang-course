package main

import "fmt"

func main() {
	a, b := 4, 2

	// 1. arithmetic operators
	//////////////////////

	// as usual, parantheses can change te order of
	// operations
	r := (a + b) / (a - b) * 2 // -> 6

	fmt.Println(r)

	r = 9 % a // -> 1

	fmt.Println(r)

	// 2. assignment operators
	//////////////////////

	c, d := 2, 3

	// increment assignment

	c += d
	fmt.Println(c) // -> 5

	// decrement assignment
	d -= 2 // -> 1

	// multiplication assignment
	//////////////////////

	e := 4
	e *= 10
	fmt.Println(e) // -> 40

	// division assignment
	//////////////////////
	e /= 2
	fmt.Println(e) // -> 20

	// modulus assignment
	//////////////////////
	e %= 8
	fmt.Println(e) // -> 4 and 5

	// increment and decrement statements
	// (not operators)
	x := 1
	x += 1
	x++
	fmt.Println(x) // -> 3
	x--
	fmt.Println(x) // -> 2

	// 3. comparison operators
	//////////////////////

	// as usual, they yield a boolean
	//types on both sides must match

	f, g := 5, 10
	fmt.Println(f == g) // -> false
	fmt.Println(f != g) // -> true

	fmt.Println(f > g)  // -> false
	fmt.Println(f <= g) // -> true

	// 3. logical operators
	//////////////////////

	// apply to booleans

	// && logical and

	h, i := 5, 10

	fmt.Println(h > 1 && i == 1) // -> true
	fmt.Println(h < 1 && i == 1) // -> false
	// as in js, && works as short circuit rule: the second expression is only evaluated,
	// if the first one returns true

	// || logical or

	fmt.Println(h == 5 || b == 8) // -? true

	// ! logical negation
	fmt.Println(!(h > 0)) // -> false
}
