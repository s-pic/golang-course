package main

import "fmt"

/*
Consider the following Go program: https://play.golang.org/p/Wo5fIlurnX6

```golang
package main
import "fmt"

func main() {
	age := -9
	if age < 0 || age > 100 {
		fmt.Println("Invalid Age")
	} else if age < 18 {
		fmt.Println("You are minor!")
	} else if age == 18 {
		fmt.Println("Congratulations! You've just become major!")
	} else {
		fmt.Println("You are major!")
	}
}
```

Change the code and use a switch statement instead of an if statement.
*/
func main() {
	age := -9
	switch {
	case age > 0 && age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	case age > 18 && age <= 100:
		fmt.Println("You are major!")
	default:
		fmt.Println("Invalid Age")
	}
}
