package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {

	c := make(chan result)
	//var results = make(map[string]string)

	urls := []string{
		"https://www.naver.com/",
		"https://nomadcoders.co/",
		"https://github.com/",
		"https://www.reddit.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- result) {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{
		url:    url,
		status: status,
	}

}
