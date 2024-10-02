package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attack(evilNinja, smokeSignal)
	//smokeSignal <- false  cannot do this for stopping. Will result in deadlock. Solved using buffered channel
	fmt.Println(<-smokeSignal)
}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
	attacked <- true
}
