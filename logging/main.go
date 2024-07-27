package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
)

func init() {
	flages := log.LUTC | log.Ltime | log.Ldate
	file, err := os.OpenFile("//home//alireza//Desktop//Learn-Golang-p1//logging//logs.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetFlags(flages)
	log.SetOutput(file)
	Debug = log.New(file, "Debug", flages)
	Info = log.New(file, "Info", flages)
	Warning = log.New(file, "Warning", flages)
	Error = log.New(file, "Error", flages)
	Fatal = log.New(file, "Fatal", flages)
}

func main() {
	for i := 1; i < 20; i++ {
		// Send request to the server using UrlRequest function
		UrlRequest("https://jsonplaceholder.typicode.com/todos/", i)
	}
	// Loop to send requests to the server Test 2
	for i := 0; i > -20; i-- {
		// Send request to the server using UrlRequest function
		UrlRequest("https://jsonplaceholder.typicode.com/todos/", i)
	}
	// Loop to send requests to the server Test 3
	for i := 0; i < 20; i++ {
		// Send request to the server using UrlRequest function
		UrlRequest("Nothing", i)
	}
}

// UrlRequest function to send HTTP request to the server
func UrlRequest(Url string, Index int) {
	defer func() {
        Res:=recover()
		if Res!=nil{
            Error.Println(Res)
            Error.Println("Recover With Recover")
        }
	}()
	if Index <= 0 {
		// Error 1: Index out of range
		Error.Println("Index out of range")
		return 
	}
	if len(Url) <= 2 {
		// Error 2: Invalid URL
		Error.Println("Invalid URL")
		return
	}
	// Form the complete URL using the Index
	Urls := fmt.Sprintf("%s%d", Url, Index)
	fmt.Println(Urls)

	// Send HTTP GET request to the server
	Request, err := http.Get(Urls)
	if err != nil {
		Error.Printf("Request: %v | Request failed \n", err.Error())
		return
	}

	// Read the response body
	RequestReadr, err := ioutil.ReadAll(Request.Body)
	if err != nil {
		Error.Printf("ReadBody: %v | ioutil.ReadAll Cant ReadBody\n", err.Error())
		return
	}
	

	// Log the response received
	Info.Printf("Request: %s | Response Received", string(RequestReadr))

	// Open or create a file for writing response bodies
	RequsetHandler, err := os.OpenFile("//home//alireza//Desktop//Learn-Golang-p1//logging//RequestHandles.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Error.Printf("Error opening FilePath: %v \n", err)
		return
	}

	// Write the response body to the specified file
	Res := bytes.NewBuffer(RequestReadr)
	RequsetHandler.Write(Res.Bytes())
	Info.Printf("Task %d completed",Index)

}
