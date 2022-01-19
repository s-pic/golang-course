package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// the os package provides basic io functionality, offering
	// the same windows to operate on files in linux, windows and mac

	// CREATE

	var newFile *os.File // a pointer to *os.File.
	// A pointer stores the memory address to another variable
	fmt.Printf("%T\n", newFile)

	// go has no exceptions
	var err error

	newFile, err = os.Create("a.txt") // created in current working dir
	// creates a file if not yet exists.
	// returns a file descriptor and eventually an error value

	if err != nil {
		//fmt.Println(err)
		//os.Exit(1) // exit with exit code 1

		// OR (recommended since idiomatic)
		log.Fatal(err) // adds timing info and has some other benefits
	}

	// TRUNCATE OR EMPTY

	err = os.Truncate("a.txt", 0) // second param means shrink file to x bytes (everything above is lost)
	if err != nil {
		log.Fatal(err)
	}

	// When done, close the file

	newFile.Close()

	// OPEN EXISTING FILE

	file, err := os.Open("a.txt") // opened in read-only mode
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// `OpenFile` gives us more options
	file2, err := os.OpenFile(
		"a.txt",
		os.O_CREATE|os.O_APPEND, // opening mode -> Create or (if exists) append to. Chained with an OR
		0644,                    // https://stackoverflow.com/a/18415918/5418403
	) // opened with permissions
	if err != nil {
		log.Fatal(err)
	}
	file2.Close()

	// GET FILE INFORMATION

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("File size in bytes: ", fileInfo.Size())
	fmt.Println("Last modified: ", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Println("Permissions", fileInfo.Mode())

	fileInfo, err = os.Stat("non-existing-file.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			log.Fatal(err)
		}
	}

	// Rename or Move a file

	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}
}
