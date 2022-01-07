package main

import "fmt"

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600

	// all slices share the same backing array,
	// so a change to s1 also mutates s3 and s4

	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s3)
	fmt.Printf("%#v\n", s4)

	// -> when a slice is created by slicing an array,
	// the array becomes the backing array of the new slice

	// to create a complete new slice from an existing slice
	// (one that is not connected to the original slice),
	// we can use append()

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"

	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

	// -> only cars has been modified, it is disconnected from `newCars`

	sli1 := []int{10, 20, 30, 40, 50}
	newSli := sli1[0:3]                   // shared backing array
	fmt.Println(len(newSli), cap(newSli)) // 3, 5

}
