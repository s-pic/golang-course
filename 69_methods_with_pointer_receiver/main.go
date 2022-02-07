package main

import "fmt"

type car struct {
	brand string
	price int
}

func (c *car) changeCar(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

func main() {
	// Reasoning: calling a function always copies the args.
	// Sometimes we want to save memory, by not doing that copying.
	// Sometimes we want to mutate the input struct/array or whatever.

	// In such situations, we can pass a pointer type as receiver
	c := car{
		brand: "Audi",
		price: 1000,
	}
	//(&c).changeCar("Mazeratti", 50000)
	// or shorter (using an implicit & operator)
	c.changeCar("Mazeratti", 50000)

	fmt.Println(c)

	var yourCar *car
	yourCar = &c
	yourCar.changeCar("VW", 3000)
	fmt.Println(*yourCar)

}
