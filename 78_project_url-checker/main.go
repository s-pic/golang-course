package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("q is DOWN\n", url)
	} else {
		// close the connection to prevent leak
		defer resp.Body.Close() // do that just before the main function exits
		fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)
		// check for success response
		if resp.StatusCode == 200 {
			// save response body to textfile
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			fileName := strings.Split(url, "//")[1] // everything after protocol
			fileName += ".txt"

			fmt.Printf("Writing response body to file %q\n", fileName)

			err = ioutil.WriteFile(fileName, bodyBytes, 0644)

			if err != nil {
				log.Fatal("Failed to write response to file", err)
			}
		}
	}
}

func main() {
	// take in a couple of URLs and check if the server is up

	urls := []string{
		"https://golang.org",
		"https://golang-invalidurl.org",
		"https://www.google.com",
		"https://www.medium.com",
	}

	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
