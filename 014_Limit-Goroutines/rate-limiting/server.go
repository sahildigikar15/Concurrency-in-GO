package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not create listener: %v", err)
	}
	var connection int32
	for {
		conn, err := li.Accept()
		if err != nil {
			continue
		}
		connection++

		go func() {
			defer func() {
				_ = conn.Close()
				atomic.AddInt32(&connection, -1)
			}()
			if atomic.LoadInt32(&connection) > 3 {
				return
			}
			time.Sleep(time.Second)
			_, err := conn.Write([]byte("Success"))
			if err != nil {
				log.Fatalf("could not write to connection: %v", err)
			}
		}()
	}
}
