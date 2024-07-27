package main

import (
	"fmt"
	"sync"
)

var (
	Cond = sync.Cond{L: &sync.Mutex{}} // Create a condition variable with a mutex
)

func main() {
	wg := sync.WaitGroup{}
	// Launch 100 goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) { // Pass 'i' as a parameter to avoid closure capture issue
			defer wg.Done() // Ensure Done is called when the goroutine completes
			Increase(i)     // Call Increase with the current index
		}(i)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All Users Have Watched")
}

func Increase(Index int) {
	Cond.L.Lock() // Lock the mutex associated with the condition variable
	defer Cond.L.Unlock() // Ensure the mutex is unlocked when the function returns

	if Index > 50 {
		Cond.Broadcast() // Broadcast to all waiting goroutines
		Streaming(Index) // Execute Streaming if the condition is met
	} else {
		Cond.Wait() // Wait until a Broadcast signal is received
		fmt.Println("The User With Id", Index, "Waiting") // Print waiting message
	}
}

func Streaming(Index int) {
	fmt.Println("The User With Id", Index, "Watching") // Print watching message
}
