package main

import (
	"fmt"
)

// var errRequestFailed = errors.New("request failed")

func main() {
	channel := make(chan string)

	go sampleFunction("John Doe", channel)
	go sampleFunction("Jane Smith", channel)
	person := <-channel
	fmt.Println("Received from channel:", person)
	person = <-channel
	fmt.Println("Received from channel:", person)
}

func sampleFunction(person string, c chan string) {
	c <- person + " is perpect"
}
