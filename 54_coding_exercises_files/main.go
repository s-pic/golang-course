package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	/*
		Coding Exercise #1
		Create a new file in the current working directory called info.txt and then close the file. If the file cannot be created (no permissions, wrong path etc) then print out the error and stop the program (do error handling).
	*/
	fileName := "info.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	/*
		Rename the file created at Exercise #1 to data.txt
		Check if file exists before renaming it. If it doesn't exist print a message and stop the program.
	*/
	newFileName := "data.txt"
	err = os.Rename(fileName, newFileName)
	if err != nil {
		log.Fatal(err)
	}
	/*
		Remove the file created at exercise #1. Take care that the file is now called data.txt (it has been renamed at exercise #2).
		Perform error handling.
	*/
	err = os.Remove(newFileName)
	if err != nil {
		log.Fatal(err)
	}
	/*
		Create a Go Program that reads the entire contents of a file called info.txt into a string. You can use ioutil.ReadAll() or any other function you want.
		The file exists in the current working directory.
	*/

	bytes, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bytes)

	fmt.Println(strings.Repeat("#", 25))

	exerciseFive()
	exerciseSix()
}

func exerciseSix() {
	/*
		1. Using single function creates a file called info.txt in the current directory. If the file already exists it will truncate it to zero size.

		2. Write the string "The Go gopher is an iconic mascot!" to the file
	*/

	str := "The Go gopher is an iconic mascot!"
	sl := []byte(str)
	name := "info.txt"

	err := ioutil.WriteFile(name, sl, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func exerciseFive() {
	/*
		Create a Go Program that reads the entire contents of a file called info.txt  using a scanner (bufio package) line by line.
		The file exists in the current working directory.
	*/
	file, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	success := scanner.Scan()

	if !success || scanner.Err() != nil {
		return
	}

	fmt.Println(scanner.Text())

	// Read whole remainin file

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

}
