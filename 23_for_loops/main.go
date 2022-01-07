package main

import "fmt"

func main() {
	// simple example
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// also works: omitting the optional, last "post statement"
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}
	// this means that golang does not need/offer an extra `while` loop
	j := 10
	for j >= 2 { // while j is positive
		fmt.Println(j)
		j--
	}

	// infinite loop
	// sum := 0
	// for {
	//  sum++
	// }
	// fmt.Println(sum) //this line is never reached

	// the `continue` statement

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue // end iteration go to post statement
		}
		fmt.Println(i)
	}

	// the break statement terminates the innermost `for` or `switch` statement
	count := 0

	for i := 0; true; i++ {
		// without a break statement, this will be an infinite loop

		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}
		if count == 10 {
			break // jump out of the loop
		}
	} // program will continue hee after the `break` statement

}
