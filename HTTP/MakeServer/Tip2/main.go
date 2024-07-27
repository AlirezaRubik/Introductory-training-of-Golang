package main

import (
	"net/http"
)

func main() {
	// Start the HTTP server and listen on port 5005
	// The second parameter is nil, which means the default multiplexer is used
	http.ListenAndServe(":5005", nil)
}
