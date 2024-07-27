package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Loop to send requests to the server Test 1
	for i := 1; i < 20; i++ {
		// Send request to the server using UrlRequest function
		err := UrlRequest("https://jsonplaceholder.typicode.com/todos/", i)
		if err != nil {
			fmt.Println(err)
		}
	}
	// Loop to send requests to the server Test 2
	for i := 0; i > -20; i-- {
		// Send request to the server using UrlRequest function
		err := UrlRequest("https://jsonplaceholder.typicode.com/todos/", i)
		if err != nil {
			fmt.Println(err)
		}
	}
	// Loop to send requests to the server Test 3
	for i := 0; i < 20; i++ {
		// Send request to the server using UrlRequest function
		err := UrlRequest("Nothing", i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// UrlRequest function to send HTTP request to the server
func UrlRequest(Url string, Index int) error {
	if Index <= 0 {
		// Error 1: Index out of range
		return fmt.Errorf("Index out of range: %w", errors.New("You must enter a valid index"))
	}
	if len(Url) <= 0 {
		// Error 2: Invalid URL
		return errors.New("Invalid URL")
	}
	// Form the complete URL using the Index
	Urls := fmt.Sprintf("%s%d", Url, Index)
	fmt.Println(Urls)

	// Send HTTP GET request to the server
	Request, err := http.Get(Urls)
	if err != nil {
		// Error 3: Failed to send request
		return fmt.Errorf("Error in sending request: %w", err)
	}
	defer Request.Body.Close()

	// Read the response body
	RequestReader, err := ioutil.ReadAll(Request.Body)
	if err != nil {
		// Error 4: Failed to read response body
		return fmt.Errorf("Error reading response body: %w", err)
	}

	// Open or create a file for writing response bodies
	RequestHandler, err := os.OpenFile("/home/alireza/Desktop/Learn-Golang-p1/Errors/CreateError/RequestHandles.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		// Error 5: Failed to open file
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer RequestHandler.Close()

	// Write the response body to the specified file
	Res := bytes.NewBuffer(RequestReader)
	_, err = RequestHandler.Write(Res.Bytes())
	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}

	return nil
}
