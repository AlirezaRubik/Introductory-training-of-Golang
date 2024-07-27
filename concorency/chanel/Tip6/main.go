package main

import (
	"fmt"
	"time"
)

// Function that sends data to a send-only channel
func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	// Close the channel after sending all data
	close(ch)
}

// Function that receives data from a receive-only channel
func receiver(ch <-chan int) {
	//tip 1
	for i:=0; i < 5; i++ {
		fmt.Println("Received:", <-ch)
	}
	//tip 2
	for v := range ch {
		fmt.Println("Received:", v)
	}
}

func main() {
	// Create a bidirectional channel
	ch := make(chan int)

	// Start the sender goroutine
	go sender(ch)

	// Start the receiver goroutine
	go receiver(ch)

	// Give some time for the goroutines to finish processing
	// (In a real program, use synchronization mechanisms to handle this properly)
    time.Sleep(time.Second*1)
}
