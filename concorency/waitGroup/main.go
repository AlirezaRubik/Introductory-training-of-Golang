package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to synchronize goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1) // Add a goroutine to the WaitGroup for each request
		go func() {
			SendRequestIndex(i, &wg) // Launch a goroutine to send HTTP requests concurrently
		}()
	}
	wg.Wait() // Wait for all goroutines to finish
	log.Println("Done sending requests")
}

// SendRequestIndex sends an HTTP GET request to a URL based on the given index
func SendRequestIndex(index int, wg *sync.WaitGroup) {
	Url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", index)
	Request, err := http.Get(Url) // Send HTTP GET request
	if err != nil {
		log.Printf("Error sending request %d: %v", index, err)
		return
	}
	defer Request.Body.Close() // Ensure the response body is closed after function exits

	ResponseBody, err := ioutil.ReadAll(Request.Body) // Read response body
	if err != nil {
		log.Printf("Error reading response body for request %d: %v", index, err)
		return
	}

	// Open or create a file to write response data
	File, err := os.OpenFile("//home//alireza//Desktop//Learn-Golang-p1//concorency//waitGroup//ResponseLog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening file for request %d: %v", index, err)
		return
	}
	defer File.Close() // Ensure the file is closed after function exits

	Res, err := File.Write([]byte(ResponseBody)) // Write response body to file
	if err != nil {
		log.Printf("Error writing to file for request %d: %v", index, err)
		return
	}

	// Log successful request details
	log.Printf("Request %d completed with status code %d and response length %d bytes", index, Request.StatusCode, Res)

	wg.Done() // Signal completion to WaitGroup
}
