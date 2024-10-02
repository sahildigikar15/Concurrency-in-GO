package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}() // fork point
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exit")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some stuff")
}
