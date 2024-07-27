package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a background context
	ctx := context.Background()

	// Create a cancellable context with a 15-second timeout
	ctx, _ = context.WithTimeout(ctx, time.Second*15)

	// Start the Increase goroutine
	go Increase(0, ctx)

	// Start the Mines function
	Mines(0, ctx)

	// Sleep for a second to allow other goroutines to run
	time.Sleep(time.Second * 1)
}

// Increase increments the index every second until the context is cancelled
func Increase(index int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Check if the context is cancelled
			fmt.Println("Increase: Done")
			return
		default:
			index++
			fmt.Println("Increase:", index)
			time.Sleep(time.Second * 1)
		}
	}
}

// Mines decrements the index every second until the context is cancelled
func Mines(index int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Check if the context is cancelled
			fmt.Println("Mines: Done")
			return
		default:
			index--
			fmt.Println("Mines:", index)
			time.Sleep(time.Second * 1)
		}
	}
}
