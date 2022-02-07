package main

import "fmt"

func foo() {
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is bar")
}

func foobar() {
	fmt.Println("This is foobar")
}

func main() {
	// the `defer` statement postpones the execution of a function until
	// the surrounding function returns

	defer foo()
	bar()

	fmt.Println("Just a string after deferring foo and calling bar")
	// -> foo is executed last
}
