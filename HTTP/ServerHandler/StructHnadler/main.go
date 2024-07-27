package main

import (
	"fmt"
	"net/http"
	"time"
)

type RequestHandler struct {
}

func main() {
	Server := http.Server{
		Addr:         ":5008",
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		Handler:      &RequestHandler{},
	}
	Server.ListenAndServe()
}
func (req *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set the response status code to 200 OK
	w.WriteHeader(http.StatusOK)

	// Send a response with formatted text (tip 1)
	fmt.Fprintf(w, "Hello This is Test [Test  1 xd]")

	// Send another response using the Write method (tip 2)
	w.Write([]byte("This is Test Tip2 xd"))
}
