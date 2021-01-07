package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// getPage -> accepts url in string format and returns the size of the page.
func getPage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func getter(url string, size chan string) {
	length, err := getPage(url)
	if err == nil {
		size <- fmt.Sprintf("%s => page size: %d", url, length)
	}
}

func main() {
	urls := []string{"http://www.apple.com", "http://www.microsoft.com", "http://www.samsung.com"}

	size := make(chan string)

	for _, url := range urls {
		go getter(url, size)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-size)
	}
}