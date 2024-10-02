package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	regularMap := make(map[int]interface{}) //regular map do not support concurrent writes
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			regularMap[0] = i
// 		}()
// 	}
// }

func main() {
	regularMap := make(map[int]interface{}) //regular map do not support concurrent writes
	syncMap := sync.Map{}

	//put
	regularMap[0] = 0
	regularMap[1] = 1
	regularMap[2] = 2

	syncMap.Store(0, 0)
	syncMap.Store(1, 1)
	syncMap.Store(2, 2)

	//get
	regularValue, Ok := regularMap[0]
	fmt.Println(regularValue, Ok)

	syncValue, syncOk := syncMap.Load(0)
	fmt.Println(syncValue, syncOk)

	//delete
	regularMap[1] = nil
	syncMap.Delete(1)

	//get and delete
	syncValue, loaded := syncMap.LoadAndDelete(2)
	regularValue = regularMap[2]
	delete(regularMap, 2)

	fmt.Println(syncValue, loaded)

	// get and put
	syncValue, loaded = syncMap.LoadOrStore(1, 1)
	mu := sync.Mutex{}
	mu.Lock()
	regularValue, Ok = regularMap[1]
	if Ok {
		regularMap[1] = 1
		regularValue = regularMap[1]
	}
	mu.Unlock()

	//range
	for key, value := range regularMap {
		fmt.Println(key, value, "|")
	}

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value, " | ")
		return true
	})

}
