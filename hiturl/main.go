package main

import (
	"net/http"
)

type URLResult struct {
	url    string
	result string
}

func main() {
	channel := make(chan URLResult)

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	for _, url := range urls {
		go hitURl(url, channel)
	}

	for range urls {
		result := <-channel
		if result.result == "ok" {
			println(result.url + " is reachable")
		} else {
			println(result.url + " is not reachable")
		}
	}
}

func hitURl(url string, c chan<- URLResult) {

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- URLResult{url: url, result: "failed"}
	}

	if resp.StatusCode == 200 {
		c <- URLResult{url: url, result: "ok"}
	}
}
