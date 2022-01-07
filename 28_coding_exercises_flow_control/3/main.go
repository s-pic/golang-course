package main

import "fmt"

/*
Change the code from the previous exercise and use the break statement to
print out only the first 3 numbers divisible by 7 between 1 and 50.
*/

func main() {
	count := 0
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)
		count += 1
		if count == 3 {
			break
		}
	}
}
