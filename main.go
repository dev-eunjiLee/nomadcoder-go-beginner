package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {

	c := make(chan requestResult)
	results := make(map[string]string)

	urls := []string{
		"https://www.naver.com/",
		"https://nomadcoders.co/",
		"https://github.com/",
		"https://www.reddit.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{
		url:    url,
		status: status,
	}

}
