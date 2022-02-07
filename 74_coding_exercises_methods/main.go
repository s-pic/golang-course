package main

import "fmt"

func main() {
	exOne()
	exTwo()
	exThree()
	exFour()
}

/*
	1. Create a new type called money. Its underlying type is float64.
	2. Create a method called print that prints out the money value with exact 2 decimal points.
*/

type Money float64

func (m Money) print() {
	fmt.Printf("%.2f\n", m)
}

func exOne() {
	money := Money(40)
	money.print()
}

// #############

func (m Money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)

}

func exTwo() {
	money := Money(40)
	fmt.Println(
		money.printStr(),
	)
}

// #############

type Book struct {
	price float64
	title string
}

func (b Book) vat() float64 {
	vatInPercent := float64(9)
	return (b.price * vatInPercent) / 100
}

func (b *Book) discount() {
	(*b).price = (*b).price - (*b).price/10
}

func exThree() {
	book := Book{
		price: 40.20,
		title: "Go to go",
	}
	fmt.Printf("The book is %#v\n", book)
	fmt.Printf("Calculated Vat for %q is %.2f\n", book.title, book.vat())
	book.discount()
	fmt.Printf("New Price for %q is %.2f\n", book.title, book.price)
	fmt.Printf("Re-Calculated Vat for %q is %.2f\n", book.title, book.vat())

}

// #################

/*
	A junior Go Programmer has just developed the following Go Program.   You want the change the book's price by calling changePrice method. However, you notice that after calling the method the price is not changed.

	You task is to change the code in order for the changePrice method to work as expected.

	package main

	import "fmt"

	type book struct {
	    title string
	    price float64
	}

	// method for book type
	func (b book) changePrice(p float64) {
	    b.price = p
	}

	func main() {
	    // book value
	    bestBook := book{title: "The Trial by Kafka", price: 9.9}

	    // changing the price
	    bestBook.changePrice(11.99)

	    fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
	}
*/

func (b *Book) changePrice(p float64) {
	b.price = p
}

func exFour() {

	bestBook := Book{title: "The Trial by Kafka", price: 9.9}

	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price)
}
