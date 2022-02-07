package main

import (
	"fmt"
	"mypackages/numbers"
) // imports are file-scoped.

// path is relative to go workspace (~/go on MacOS by default).
// When loading a package, we import the whole directory (all go files in it)

// -> a package is spread about multiple files

func main() {
	n := 40
	fmt.Printf("%d is even: %t \n",
		n,
		numbers.Even(n),
	)
	numbers.Bar()

	sayHello("Sascha") // belongs to the same package, so no namespace needed

	tf := celsiusToFahrenHeit(BOILING_POINT_IN_DEG)
	fmt.Println("Water boilingpoint in degrees Fahrenheit:", tf)
}
