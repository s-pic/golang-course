package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too expensive!")
	}

	if price <= 100 && inStock == true {
		fmt.Println("Buy it!")

	}

	// in golang, there is no "Truthiness" as other languages
	// like JS, Java or Python have it.
	//We can only use Booleans as conditions for an if clause

	//if price {
	//	fmt.Println("price is truthy")
	//} --> compiler error: non-bool price (type int) used as if condition

	if price < 80 {
		fmt.Println("Its cheap!")
	} else if price == 100 {
		fmt.Println("On The edge!")
	} else {
		fmt.Println("too expensive")
	}

	age := 50

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote!. Please return in %d years\n", 18-age)
	} else if age == 18 {
		fmt.Println("This is your first vote!")
	} else if age > 18 && age <= 150 {
		fmt.Println("Please vote, it's important!")
	} else {
		fmt.Printf("Invalid age!")
	}
}
