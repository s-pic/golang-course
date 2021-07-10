package main

import "fmt"

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher"

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z
	sub()
	sub2()
}

func sub() {
	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false

	_, _, _ = a, x, y
}

const version = "3.1"

func sub2() {
	name := "Golang"
	fmt.Println(name)
}

/*
Coding Exercise #1

Using the var keyword, declare 3 variables called a, b, c, d of type int, float64, bool and string.

Using short declaration syntax declare and assign these values to variables x, y and z:

- 20

- 15.5

- "Gopher!"

Using fmt.Println() print out the values of a, b, c, d, x, y and z.

Try to run the program without error.

Do you wonder what Gopher is?  Check it here: https://blog.golang.org/gopher

Are you stuck? Do you want to see the solution for this exercise? Click here.



Coding Exercise #2

Change the code from the previous exercise and

1. Declare a, b, c, d on a single line for better readability -> multiple declarations

2. Declare x, y and z on a single line -> multiple short declarations

3. Remove the statement that prints out the variables. See the error!

4. Change the program to run without error using the blank identifier (_)

Are you stuck? Do you want to see the solution for this exercise? Click here.



Coding Exercise #3

There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

package main

func main() {
    var a float64 = 7.1

    x, y := true, 3.7

    a, x := 5.5, false

    _, _, _ = a, x, y
}
Are you stuck? Do you want to see the solution for this exercise? Click here.



Coding Exercise #4

There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

package main

version := "3.1"

func main() {
	name := 'Golang'
	fmt.Println(name)
}


Are you stuck? Do you want to see the solution for this exercise? Click here.



*/
