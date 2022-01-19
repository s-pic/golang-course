package main

import "fmt"

func main() {
	var employees map[string]string // keys and values are strings; initialized with its zero value (nil)
	fmt.Printf("%#v\n", employees)

	// print num of key-value pairs
	fmt.Printf("%d\n", len(employees))

	// we can already try to get the value for a key, even if the key is not used
	fmt.Printf("the value for key %q is %q\n", "Dan", employees["Dan"])
	// the zero-value for the respective type is returned (in this case string)

	var accounts map[string]float64
	fmt.Printf("%f\n", accounts["any string"]) // 0.000000

	// we can only use "comparable" keys (so e.g. no slices):

	//var m1 map[[]int]string --> compile error:
	//Invalid map key type: comparison operators == and != must be fully defined for the key type

	var m1 map[[5]int]string // arrays are valid keys though
	_ = m1

	// adding entries
	//employees["Dan"] = "Programmer" -> Runtime Error: panic: assignment to entry in nil map

	// IN GO, IT IS ILLEGAL TO ADD TO AN UNINITIALIZED MAP

	people := map[string]float64{} // initialized but empty

	people["John"] = 21.4
	people["John"] = 23.4444 // value updated
	people["Marry"] = 10
	fmt.Println(people)

	// second way to create an initialized, but empty map
	map1 := make(map[int]int)
	map1[4] = 8
	fmt.Println(map1)

	// another way to init the map: a map literal
	balances := map[string]float64{
		"USD": 323.23,
		"EUR": 232.12,
		"PUN": 0.000000,
	}
	fmt.Println(balances)

	// As learned, accessing a mao value by a key that is not set
	// returns the zero value.
	// In order to distinguish such a missing entry
	// from an existing entry that is a zero value,
	// Go offers the so called "comma ok Idiom"

	e1, ok := balances["Ron"]
	if ok {
		fmt.Printf(`Entry for key "Ron" is %f`, e1)
	} else {
		fmt.Println("The \"Ron\" key does not exist in the map")
	}

	e2, ok := balances["PUN"] // existing entry
	if ok {
		fmt.Printf(`Entry for key "PUN" is %f`, e2)
		fmt.Println("")
	}

	// iterate over map
	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %f\n", k, v)
	}

	// delete from map
	delete(balances, "PUN")
	fmt.Println(balances)
}
