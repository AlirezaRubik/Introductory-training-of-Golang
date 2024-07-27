package main

import (
	"fmt"
	"sync"
)

func main() {
	// Initialize a WaitGroup to wait for the completion of all goroutines
	wg := sync.WaitGroup{}

	// Initialize a Mutex to ensure exclusive access to the shared variable 'count'
	Mutex := sync.Mutex{}

	// Shared variable that will be incremented by goroutines
	count := 0

	// Loop to create 10,000,000 goroutines
	for i := 0; i < 10000000; i++ {
		// Add a count to the WaitGroup for each goroutine
		wg.Add(1)

		// Create a new goroutine
		go func() {
			// Lock the Mutex to ensure only one goroutine increments 'count' at a time
			Mutex.Lock()
			count++
			
			// Unlock the Mutex and signal that this goroutine is done
			defer Mutex.Unlock()
			defer wg.Done()
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the final value of 'count'
	fmt.Println(count)
}
