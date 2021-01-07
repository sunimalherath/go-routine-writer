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

func worker(urlCh chan string, sizeCh chan string, id int) {
	for {
		url := <-urlCh
		length, err := getPage(url)
		if err == nil {
			sizeCh <- fmt.Sprintf("%s => page size: %d (%d)", url, length, id)
		} else {
			sizeCh <- fmt.Sprintf("%s => could not get size - error %s", url, err)
		}
	}
}

func main() {
	urls := []string{"http://www.apple.com", "http://www.microsoft.com", "http://www.samsung.com"}

	sizeCh := make(chan string)
	urlCh := make(chan string)

	for i := 0; i < 10; i++ {
		go worker(urlCh, sizeCh, i)
	}

	for _, url := range urls {
		urlCh <- url
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}
}