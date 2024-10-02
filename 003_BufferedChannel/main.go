package main

import "fmt"

func main() {
	channel := make(chan string, 1)
	channel <- "First message"
	fmt.Println(<-channel) // we are sending as well as receiving message in same goroutine because of buffered channel
	// if we send two message - will again give error as capacity is 1
	// channels are FIFO queue
}
