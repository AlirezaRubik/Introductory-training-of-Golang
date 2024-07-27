package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Test copying a file
	err := File("/home/alireza/Desktop/Learn-Golang-p1/Errors/WrappingErrors/Tip2/logs.log", "/home/alireza/Desktop/Learn-Golang-p1/Errors/WrappingErrors/Tip2/NewLog.log")
	if err != nil {
		fmt.Println(err)
	}
}

// File copies the content of the source file to the destination file
func File(source string, destination string) error {
	// Check if the source path is at least 3 characters long
	if len(source) < 3 {
		return fmt.Errorf("error in source file length")
	}
	// Check if the destination path is at least 3 characters long
	if len(destination) < 3 {
		return fmt.Errorf("error in destination file length")
	}

	// Open the source file
	Source, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer Source.Close()

	// Create the destination file
	Destination, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer Destination.Close()

	// Copy the content from source to destination
	_, err = io.Copy(Destination, Source)
	if err != nil {
		return fmt.Errorf("error copying content: %v", err)
	}

	return nil
}
