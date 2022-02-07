package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		s := fmt.Sprintf("⚠️ %q is DOWN\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s
	} else {
		defer resp.Body.Close() // do that just before the main function exits
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			fileName := strings.Split(url, "//")[1]
			fileName += ".txt"

			s += fmt.Sprintf("Writing response body to file %q\n", fileName)

			err = ioutil.WriteFile(fileName, bodyBytes, 0644)

			if err != nil {
				//log.Fatal("Failed to write response to file", err)
				s += "Error writing file\n"
				c <- s
			}
			s += fmt.Sprintf("%q is UP", url)
			c <- s
		}
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://golang-invalidurl.org",
		"https://www.google.com",
		"https://www.medium.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}
