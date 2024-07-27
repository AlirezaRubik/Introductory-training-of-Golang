package main

import (
	"net/http"
	"time"
)

func main() {
	// Create a new HTTP server with specific configurations
	Server := http.Server{
		Addr:         ":5005",            // Set the address and port the server will listen on
		WriteTimeout: time.Second * 10,   // Set the maximum duration before timing out writes of the response
		ReadTimeout:  time.Second * 10,   // Set the maximum duration for reading the entire request, including the body
	}

	// Start the server and listen on the specified address and port
	// This function will block the execution and run the server until it is stopped
	Server.ListenAndServe()
}
