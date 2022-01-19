package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          uint32
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact // Contact is embedded
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "john@foo.com",
			address: "Bar Street 3",
			phone:   0345443,
		},
	}

	fmt.Println(345443 == john.contactInfo.phone)
}
