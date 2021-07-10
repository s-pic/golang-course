package main

import "fmt"

/*
Write the correct logical operator (&&, || , !) inside the expression so that result1 will be false and result2 will be true.

Program source code: https://play.golang.org/p/74SCleChC20

package main

import "fmt"

func main() {
    x, y := 0.1, 5
    var z float64

    // Write the correct logical operator (&&, || , !)
    // inside the expression so that result1 will be false and result2 will be true.

    result1 := x < z         int(x) != int(z)
    fmt.Println(result1)

    result2 := y == 1 * 5        int(z) == 0
    fmt.Println(result2)

}
*/
func main() {
	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := x < z && int(x) != int(z)
	fmt.Println(result1) // -> false

	result2 := y == 1*5 || int(z) == 0
	fmt.Println(result2) // -> true
}
