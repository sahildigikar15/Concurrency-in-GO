package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		defer wg.Done()
		// count += 10
		atomic.StoreInt32(&count, 10)
	}()
	go func() {
		defer wg.Done()
		// count -= 15
		atomic.StoreInt32(&count, -15)
	}()
	go func() {
		defer wg.Done()
		// count++
		atomic.StoreInt32(&count, 1)
	}()
	go func() {
		defer wg.Done()
		// count = 0
		atomic.StoreInt32(&count, 0)
	}()
	go func() {
		defer wg.Done()
		// count = 100
		atomic.StoreInt32(&count, 100)
	}()
	wg.Wait()
	fmt.Println("count", count)
}
