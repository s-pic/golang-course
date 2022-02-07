package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

/*

Coding Exercise #1

Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3).




Coding Exercise #2

Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:

a) the factorial of n

b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)

Test the program by calling the function.




Coding Exercise #3

Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes (this is in fact a string). If the argument is integer 'n', the function should return the result of n + nn + nnn

Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107




Coding Exercise #4 and 5

Create a function with the identifier sum that takes in a variadic parameter of type int and returns the sum of all values of type int passed in.




Coding Exercise #5

Change the function from the previous exercise and use a `naked return`.




Coding Exercise #6

Create a function called searchItem() that takes 2 parameters: a) a string slice and b) a string.

The function should search for the string (the second parameter) in the slice (the first parameter) and returns true if it finds the string in the slice and false otherwise. Do function does an case-sensitive search.

Call the function and see how it works.

Example:

animals := []string{"lion", "tiger", "bear"}
result := searchItem(animals, "bear")
fmt.Println(result) // => true
result = searchItem(animals, "pig")
fmt.Println(result)     // => false



Coding Exercise #7

Change the function from the previous exercise to do a case-insensitive search.

Example:

animals := []string{"Lion", "tiger", "bear"}
result := searchItem(animals, "beaR")
fmt.Println(result) // => true
result = searchItem(animals, "lion")
fmt.Println(result)     // => true



Coding Exercise #8

Consider the following Go program that prints out:

The Go gopher is the iconic mascot of the Go project.

Hello, Go playground!



package main

import "fmt"

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}
Modify only the line in the main() body function where the print() function is invoked so that the program will print out Hello, Go playground! and then The Go gopher is the iconic mascot of the Go project.


Coding Exercise #9

Create a function that takes in an int value and prints out that value.

Assign the function to a variable, print out the type of the variable and then call that function through the variable name.

*/

func main() {
	exOne()
	extwo()
	exThree()
	exFour()
	exFive()
	exSix()
	exSeven()
	exEight()
	exNine()
}

func cube(num float64) float64 {
	return math.Pow(num, 3)
}

func exOne() {
	fmt.Println(cube(4.2))
}

func extwo() {
	for n := uint(0); n < 100; n++ {
		factorial, sum := f1(n)
		fmt.Printf("%d -> Sum %d, Facorial %d\n", n, sum, factorial)
	}
}

// factorial(n): n * (n -1) * (n - 2) .. (n - i) while i < n)
// 5! =5 * 4 and 5 * 3 * 2 * 1= 5 * 24= 120
// sum:  the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)
func f1(num uint) (factorial uint, sum uint) { // TODO: fix
	if num < 2 {
		factorial, sum = num, num
		return
	}

	factorial = num
	sum = num
	for n := uint(1); n < num; n++ {
		sum += n
		factorial *= num - n
	}
	return
}
func myFunc(intInQuotes string) int {
	sum := 0
	parsedInt, err := strconv.Atoi(intInQuotes)
	sum += parsedInt
	parsedInt, err = strconv.Atoi(intInQuotes + intInQuotes)
	sum += parsedInt
	parsedInt, err = strconv.Atoi(intInQuotes + intInQuotes + intInQuotes)
	sum += parsedInt

	if err != nil {
		log.Fatal("Failed to parse ", intInQuotes)
	}
	return sum
}
func exThree() {
	fmt.Println(myFunc("5"))
}

func exFour() {
	fmt.Println(
		sum(1, 3, 5, 12),
	)
}

func sum(nums ...int) int {
	returnedSum := 0
	for _, summand := range nums {
		returnedSum += summand
	}
	return returnedSum
}

func exFive() {
	fmt.Println(
		sumWithNakedReturn(1, 3, 5, 12),
	)
}

func sumWithNakedReturn(nums ...int) (returnedSum int) {
	returnedSum = 0
	for _, summand := range nums {
		returnedSum += summand
	}
	return
}

func exSix() {
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
}

func searchItem(strSlice []string, str string) bool {
	found := false
	for _, val := range strSlice {
		if val == str {
			found = true
		}
	}
	return found
}

func exSeven() {
	animals := []string{"Lion", "tiger", "bear"}
	result := searchItemCaseInsensitive(animals, "beaR")
	fmt.Println(result) // => true
	result = searchItemCaseInsensitive(animals, "lion")
	fmt.Println(result) // => true

}

func searchItemCaseInsensitive(strSlice []string, str string) bool {
	found := false
	for _, val := range strSlice {
		if strings.ToLower(val) == strings.ToLower(str) {
			found = true
		}
	}
	return found
}

func print(msg string) {
	fmt.Println(msg)
}

func exEight() {
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}

func someFunc(val int) {
	fmt.Println(val)
}

func exNine() {
	funcVar := someFunc

	fmt.Printf("Tyoe of funcVar is %T", funcVar)

	funcVar(4)
}
