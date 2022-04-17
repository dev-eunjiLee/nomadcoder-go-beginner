package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request Fail")

func main() {
	urls := []string{
		"https://www.naver.com/",
		"https://nomadcoders.co/",
		"https://github.com/",
	}

	for _, url := range urls {
		hitURL(url)
	}

}

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
