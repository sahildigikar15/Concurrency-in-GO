package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})

	go func() {
		work()
		done <- struct{}{}
	}() // fork point
	<-done
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exit")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some stuff")
}
