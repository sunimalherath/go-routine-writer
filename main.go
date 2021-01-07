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

func main() {
	urls := []string{"http://www.apple.com", "http://www.microsoft.com", "http://www.samsung.com"}

	for _, url := range urls {
		pageSize, err := getPage(url)
		if err != nil {
			fmt.Printf("Could not get page size for %s", url)
		}
		fmt.Printf("%s, page size: %d\n", url, pageSize)
	}
}
