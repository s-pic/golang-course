package main

import "fmt"

func main() {
	exOne()
	exTwo()
	exThree()
}

/*
Consider the following type and interface declaration.

type vehicle interface {
    License() string
    Name() string
}

type car struct {
    licenceNo string
    brand     string
}
1. Create a Go program where car type implements the vehicle interface.

2. Create a variable of type vehicle that holds a car struct value.

3. Call the methods (Licence and Name) on the interface value declared at step 2

4 and 5. Run the program without errors.
*/

type Vehicle interface {
	license() string
	name() string
}

type Car struct {
	licenceNo string
	brand     string
}

func (c Car) license() string {
	return c.licenceNo
}

func (c Car) name() string {
	return c.brand
}

func exOne() {
	var v Vehicle
	v = Car{
		licenceNo: "ABC123",
		brand:     "Honda",
	}
	fmt.Println("License:", v.license(), "| Name:", v.name())
}

// #######################

func exTwo() {
	var empty interface{}
	fmt.Printf("Type of empty: %T\n", empty) // nil
	empty = 5
	fmt.Printf("Type of empty: %T\n", empty) // int -> underlying type
	empty = 20.3
	fmt.Printf("Type of empty: %T\n", empty) // float
	empty = []int{4, 5, 6}
	fmt.Printf("Type of empty: %T\n", empty)         // []int
	empty = append(empty.([]int), 99)                // type assertion
	fmt.Printf("Slice after append is %#v\n", empty) // []int{4 and 5,5,6,99}

}

// #######################
type Cube struct {
	edge float64
}

func volume(c Cube) float64 {
	return c.edge * c.edge * c.edge
}

func exThree() {
	var x interface{}
	x = Cube{edge: 5}
	castedX, didCastSuccessfully := x.(Cube)
	if didCastSuccessfully == true {
		v := volume(castedX)
		fmt.Printf("Cube Volume: %v\n", v)
	}

}
