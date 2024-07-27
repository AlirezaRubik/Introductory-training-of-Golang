package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// ErrorMaker struct represents a custom error with relevant information.
type ErrorMaker struct {
	Message    string // Message describing the error
	Errors     error  // Original error returned
	ClientCode int    // Client-specific error code
	Result     bool   // Result indicating success or failure
}

func main() {
	// TEST 1: Attempt to copy a file without specifying the destination
	Test1 := ScanFile("//home//alireza//Desktop//Learn-Golang-p1//Errors//WrappingErrors//Tip1//logs.log", "")
	fmt.Printf("%v \n", Test1)

	// TEST 2: Attempt to copy a file specifying both source and destination
	Test2 := ScanFile("//home//alireza//Desktop//Learn-Golang-p1//Errors//WrappingErrors//Tip1//logs.log", "//home//alireza//Desktop//Learn-Golang-p1//Errors//WrappingErrors//Tip1//BackUplogs.log")
	fmt.Printf("%v \n", Test2)

	// TEST 3: Repeat of TEST 1 to demonstrate handling the same error scenario
	Test3 := ScanFile("//home//alireza//Desktop//Learn-Golang-p1//Errors//WrappingErrors//Tip1//logs.log", "")
	fmt.Printf("%v \n", Test3)
}

// GenerateError creates a new ErrorMaker instance with provided details.
func GenerateError(Message string, Errors error, ClientCode int, Result bool) *ErrorMaker {
	err := ErrorMaker{
		Message:    Message,
		Errors:     Errors,
		ClientCode: ClientCode,
		Result:     Result,
	}
	return &err
}

// ScanFile attempts to copy a file from Source to Destination.
// It returns an ErrorMaker instance based on the success or failure of the operation.
func ScanFile(Source, Destination string) interface{} {
	// Check if the length of Source path is valid
	if len(Source) <= 2 {
		return GenerateError("Invalid source string Len", errors.New("Invalid Source Path"), -1, false)
	}
	// Check if the length of Destination path is valid
	if len(Destination) <= 2 {
		return GenerateError("Invalid Destination string Len", errors.New("Invalid Destination Path"), -1, false)
	}

	// Open the source file
	OpenSourceFile, err := os.Open(Source)
	if err != nil {
		return GenerateError("Error Opening Source File", err, -1, false)
	}
	defer OpenSourceFile.Close()

	// Create the destination file
	OpenDestinationFile, err := os.Create(Destination)
	if err != nil {
		return GenerateError("Error Create Destination File", err, -1, false)
	}
	defer OpenDestinationFile.Close()

	// Copy the content from source to destination
	_, err = io.Copy(OpenDestinationFile, OpenSourceFile)
	if err != nil {
		return GenerateError("Error Copy Source File", errors.New("Error copying"), -1, false)
	}

	// Return success message
	return GenerateError("File copied successfully", nil, 200, true)
}
