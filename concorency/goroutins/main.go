package main

import (
	"fmt"
	"time"
)

func main() {
	// Launch an anonymous goroutine that prints "Testing"
	go func() {
		fmt.Println("Testing")
	}()
	
	// Launch additional test functions as goroutines
	go Test1()
	go Test2()
	go Test3()
	go Test4()
	go Test5()

	// Sleep for 1 second to allow all goroutines to complete
	time.Sleep(time.Second * 1)
}

// Test1 prints "Test function 1"
func Test1() {
	fmt.Println("Test function 1")
}

// Test2 prints "Test function 2"
func Test2() {
	fmt.Println("Test function 2")
}

// Test3 prints "Test function 3"
func Test3() {
	fmt.Println("Test function 3")
}

// Test4 prints "Test function 4"
func Test4() {
	fmt.Println("Test function 4")
}

// Test5 prints "Test function 5"
func Test5() {
	fmt.Println("Test function 5")
}
