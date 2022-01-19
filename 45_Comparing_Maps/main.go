package main

import "fmt"

func main() {
	// Maps are not comparable and can thus only compared to nil

	a, b := map[string]string{"A": "X"}, map[string]string{"B": "Y"}

	// what we can do is to create and compare string representations
	// of these maps

	s1 := fmt.Sprintf("%s", a) // dont print to stdout, but write to memory (assign to var)
	s2 := fmt.Sprintf("%s", b) // dont print to stdout, but write to memory (assign to var)

	fmt.Println(s1, s2)

	fmt.Println("Both maps are equal", s1 == s2)
}
