package main

import "fmt"

/*
There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

package main

import "fmt"

type duration int

func main() {
    var hour duration = 3600
    minute := 60
    fmt.Println(hour != minute)
}
*/

type duration int

func main() {
	var hour duration = 3600
	minute := 60
	fmt.Println(hour != duration(minute))
}
