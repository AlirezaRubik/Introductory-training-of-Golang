package main

import (
	requestsystem "information/Request"
	"net/http"
	"time"
)
func main(){
	Server:=http.Server{
        Addr: ":5005",
		WriteTimeout: time.Second*10,
		ReadTimeout: time.Second*10,
		Handler: &requestsystem.Request{},
	}
	Server.ListenAndServe()
}