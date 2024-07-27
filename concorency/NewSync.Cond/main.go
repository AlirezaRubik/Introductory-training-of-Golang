package main

import (
	"fmt"
	"sync"
	"time"
)

var broadcast = false
var Users = []int{}

func main() {
	// Create a new condition variable with a mutex
	cond := sync.NewCond(&sync.Mutex{})
	// Start 100 goroutines to simulate users trying to stream
	for i := 0; i < 100; i++ {
		go Streaming(i, cond)
	}
	// Wait for a while to allow goroutines to run
	time.Sleep(time.Second * 1)
}

func Streaming(index int, cond *sync.Cond) {
	fmt.Println("The", index, "Waiting Stream")
	// Check the view count and potentially broadcast
	CheckViewCount(cond, index)
	cond.L.Lock()
	// Wait until broadcast signal is received
	if !broadcast {
		cond.Wait()
	} else {
		fmt.Println("The", index, "Watching Stream")
	}
	defer cond.L.Unlock()
}

func CheckViewCount(cond *sync.Cond, index int) {
	Users = append(Users, index)
	cond.L.Lock()
	// Broadcast signal if user count exceeds 50
	if len(Users) > 50 {
		broadcast = true
		cond.Broadcast()
	}
	fmt.Println("The", index, "Waiting")
	defer cond.L.Unlock()
}
