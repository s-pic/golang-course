package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func panicOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("../test.txt")
	panicOnError(err)
	defer file.Close()

	byteSlice := make([]byte, 2) // size of 2 bytes

	numOfBytesRead, err := io.ReadFull(file, byteSlice)
	panicOnError(err)
	fmt.Println("Num of bytes read:", numOfBytesRead)
	fmt.Printf("Data read: %s\n", byteSlice) // He

	// Now read Entire Content -> Returns slice of unknown size

	file, err = os.Open("main.go")
	panicOnError(err)

	data, err := ioutil.ReadAll(file)
	panicOnError(err)

	fmt.Printf("data as string: %s\n", data)
	fmt.Println("Num of bytes read:", len(data))

	// yet another way -,-

	data, err = ioutil.ReadFile("main.go")
	panicOnError(err)
	fmt.Printf("data as string: %s\n", data)

}
