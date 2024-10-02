package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock   sync.Mutex
	rwlock sync.RWMutex
	count  = 0
)

func main() {
	basics()
}

func readWrite() {
	go read()
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read() {
	rwlock.RLock()
	defer rwlock.Unlock()

	fmt.Println("Reading locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Reading unlocking")
}

func write() {
	rwlock.Lock()
	defer rwlock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlocking")
}

func basics() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is", count)
}

func increment() {
	lock.Lock() // so only 1 goroutine has access to count variable at a time
	count++
	lock.Unlock()
}
