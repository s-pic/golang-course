package main

import f "fmt" // imports are file-scoped

const done = true // visible to all files of this package

var b = 10 // not used, but no error
// -> can be used in other files of same package

func main() { // also package scoped
	var task = "running" // block scoped

	const done = "valid go code" // block scope

	// -> local scope shadows package scope

	f.Println(task, done)

	f1()
}

func f1() {
	// block scope includes package scope
	f.Printf("done in f1(): %v\n", done)
}
