package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func panicOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("my_file.txt")
	panicOnError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// here, right before the first scan, we could also specify what delimiter
	// will lead to the scanner splitting the files in chunks
	// e.g. using:
	// scanner.Split(bufio.ScanWords)

	success := scanner.Scan()

	if success == false {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached EOF")
		} else {
			panicOnError(err)
		}
	}

	fmt.Println("First line found:", scanner.Text())

	// Read whole remainin file

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if err := scanner.Err(); err != nil {
			panicOnError(err)
		}
	}
}
