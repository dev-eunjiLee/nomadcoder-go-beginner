package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request Fail")

func main() {

	var results = make(map[string]string)

	urls := []string{
		"https://www.naver.com/",
		"https://nomadcoders.co/",
		"https://github.com/",
		"https://www.reddit.com/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}

}

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
