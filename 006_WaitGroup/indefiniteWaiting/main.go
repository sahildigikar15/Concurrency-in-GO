package main

import "sync"

func main() {
	var beeper sync.WaitGroup
	beeper.Add(1)
	// if we add done twice then will give negative waitgroup error
	//beeper.Done()
	//beeper.Done()
	beeper.Wait() // before calling done we are calling wait do deadlock error
}
