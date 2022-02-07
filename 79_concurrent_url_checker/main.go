package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, group *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("⚠️ %q is DOWN\n", url)
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
	group.Done()
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://golang-invalidurl.org",
		"https://www.google.com",
		"https://www.medium.com",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}

	wg.Wait()
}
