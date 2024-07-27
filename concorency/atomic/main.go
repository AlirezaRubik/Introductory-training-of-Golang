package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Initialize the salary variable with an initial value
var Salary int64 = 11000000

func main() {
	// Create a WaitGroup to wait for the completion of all goroutines
	wg := sync.WaitGroup{}

	// Loop to create 10 goroutines
	for i := 0; i < 10; i++ {
		// Add a count to the WaitGroup for each goroutine
		wg.Add(1)
		// Create a new goroutine
		go func() {
			// Call the function to pay the salary and pass the WaitGroup
			GiveSalarys(&Salary, &wg)
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the final value of 'Salary' after all payments
	fmt.Println("After paying for 10 people we have:", Salary)
}

// Function to pay the salary to one person
func GiveSalarys(Salary *int64, wg *sync.WaitGroup) {
	// Define the amount to be paid per person
	PerPerson := 1000000

	// Atomically subtract the salary amount from the total Salary
	Res := atomic.AddInt64(Salary, -int64(PerPerson))

	// Print the remaining Salary after the payment
	fmt.Println("After payment we have:", Res)

	// Signal that this goroutine is done
	wg.Done()
}
