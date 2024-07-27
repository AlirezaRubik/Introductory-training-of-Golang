package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a new ServeMux to handle routing
	Mux := http.NewServeMux()
	
	// Register the handler function for the "/test" route
	Mux.HandleFunc("/test", SetRequest)
	
	// Create a new HTTP server with specific configurations
	Server := http.Server{
		Addr:         ":5005",            // Set the address and port the server will listen on
		WriteTimeout: time.Second * 10,   // Set the maximum duration before timing out writes of the response
		ReadTimeout:  time.Second * 10,   // Set the maximum duration for reading the entire request, including the body
		Handler:      Mux,                // Assign the ServeMux as the handler
	}
	
	// Start the server and listen on the specified address and port
	// This function will block the execution and run the server until it is stopped
	Server.ListenAndServe()
}

// Handler function for the "/test" route
func SetRequest(w http.ResponseWriter, r *http.Request) {
	// Set the response status code to 200 OK
	w.WriteHeader(http.StatusOK)
	
	// Send a response with formatted text (tip 1)
	fmt.Fprintf(w, "Hello This is Test [Test 1]")
	
	// Send another response using the Write method (tip 2)
	w.Write([]byte("This is Test Tip2 "))
}
