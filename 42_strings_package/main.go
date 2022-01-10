package main

import (
	"fmt"
	"strings"
)

func printBreak() {
	line := strings.Repeat("#", 20)
	fmt.Println(line)
}

func main() {
	p := fmt.Println

	// Contains
	boolResult := strings.Contains("I love Go Programming", "love")
	p(boolResult) // -> true -> contains substring
	p(strings.Contains("I love Go Programming", "Paula"))
	printBreak()

	// ContainsRune
	boolResult = strings.ContainsRune("I love Go Programming", 'G')
	p(boolResult)
	printBreak()

	// ContainsAny
	boolResult = strings.ContainsAny("success", "xy")
	p(boolResult)                            // false -> contains neither x or y
	p(strings.ContainsAny("success", "sce")) // true
	printBreak()

	// Count
	count := strings.Count("Giraffe", "f")
	p(count) // 2
	printBreak()
	// watch out
	count = strings.Count("Five", "")
	p(count) // 5 -> num of runes + 1

	// ToLower
	strResult := strings.ToLower("PAULA")
	p(strResult) // paula
	printBreak()

	// ToUpper
	strResult = strings.ToUpper("paula")
	p(strResult) // PAULA
	printBreak()

	// EqualFold
	p("go" == "go")                                  // true
	boolResult = strings.EqualFold("Paula", "paula") // true
	p(boolResult)
	p(strings.EqualFold("Paula", "PAULA")) // true
	p(strings.EqualFold("Paula", "Finja"))

	// Replace
	strResult = strings.Replace("192.168.0.1", ".", ":", 2)
	// replace first two dot characters with :
	p(strResult)
	// pass -1 so that all occurrences are replaces
	p(strings.Replace("192.168.0.1", ".", ":", -1))

	// easier solution:
	p(strings.ReplaceAll("192.168.0.1", ".", ":"))

	printBreak()

	// Split
	sliceResult := strings.Split("a,b,c", ",")
	fmt.Printf("%#v\n", sliceResult)
	sliceResult = strings.Split("Go fo go!", "") // split after each rune
	fmt.Printf("%#v\n", sliceResult)

	// convenient way to split a string at all white spaces and line breaks:
	s := "Orange Green \n Blue  Yellow"
	sliceResult = strings.Fields(s)
	fmt.Printf("%#v\n", sliceResult)

	printBreak()

	// Join
	sl := []string{"I", "learn", "go"}
	fmt.Printf("%#v\n", strings.Join(sl, "-"))

	// Trim
	s = strings.TrimSpace("\t Goodbye Windows, Welcome Linux \n")
	// -> Remove all leading and trailing whitespaces, tabs or new lines
	fmt.Printf("%q\n", s)

	s = strings.Trim("Hello Gophers!!!?", ".!?")
	p(s) // prints "Hello Gophers"

	printBreak()

	sub()
}

func sub() {
	s1, s2 := "gO", "Go"

	result1 := strings.ToLower(s1) == strings.ToLower(s2)
	result2 := strings.EqualFold(s1, s2)

	fmt.Println(result1, result2)
}
