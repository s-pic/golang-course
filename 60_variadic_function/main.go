package main

import (
	"fmt"
	"strings"
)

func variadicFunc(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50
}

func SumAndProduct(a ...float64) (sum float64, product float64) {
	product = 1

	for _, val := range a {
		sum += val
		product *= val
	}
	return
}

func PersonInformation(nonVariadic int, variadicParams ...string) string {
	fullName := strings.Join(variadicParams, " ")

	return fmt.Sprintf("Age: %d, name: %s", nonVariadic, fullName)
}

func main() {
	variadicFunc(1, 2, 3, 4)
	variadicFunc() // valid call -> variadic means 0 or more args. slice contains nil

	// append is a variadic function
	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5, 6)

	// we can also pass in a slice to a variadic function
	variadicFunc(nums...)

	f2(nums...) // this modifies the input slice -> the function did not create a new slice because it received one
	fmt.Println("Nums", nums)

	s, p := SumAndProduct(2, 5, 10)
	fmt.Println("Sum", s, "Product", p)

	info := PersonInformation(30, "Wolfgang", "Amadeus")

	fmt.Println(info)
}
