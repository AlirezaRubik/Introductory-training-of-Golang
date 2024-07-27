package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	// Set the global log level to Info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// Set the time format for logs to Unix format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Loop to send requests to the server
	for i := 1; i < 20; i++ {
		// Send request to the server using UrlRequest function
		logger := UrlRequest("https://jsonplaceholder.typicode.com/todos/", i)
		logger.Info().Msg(fmt.Sprintf("Task With %d ID Done", i))
	}
}

// UrlRequest function to send HTTP request to the server
func UrlRequest(Url string, Index int) *zerolog.Logger {
	// Open or create a file for writing logs
	FilePath, err := os.OpenFile("//home//alireza//Desktop//Learn-Golang-p1//zerolog//logs.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	logger := zerolog.New(FilePath).With().Timestamp().Logger()

	// Form the complete URL using the Index
	Urls := fmt.Sprintf("%s%d", Url, Index)
	fmt.Println(Urls)

	// Send HTTP GET request to the server
	Request, err := http.Get(Urls)
	if err != nil {
		logger.Error().Str("Request:", err.Error()).Msg("Request failed")
	}

	// Read the response body
	RequestReadr, err := ioutil.ReadAll(Request.Body)
	if err != nil {
		logger.Error().Str("ReadBody:", err.Error()).Msg("ioutil.ReadAll Cant ReadBody")
	}
	defer Request.Body.Close()

	// Log the response received
	logger.Info().Str("Request:", string(RequestReadr)).Msg("Response Received")

	// Open or create a file for writing response bodies
	RequsetHandler, err := os.OpenFile("//home//alireza//Desktop//Learn-Golang-p1//zerolog//RequestHandles.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.Error().Str("File:", err.Error()).Msg("Handler File Error")
	}

	// Write the response body to the specified file
	Res := bytes.NewBuffer(RequestReadr)
	RequsetHandler.Write(Res.Bytes())

	// Return the logger for use in the main function
	return &logger
}
