package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	Written, err := File("/home/alireza/Desktop/Learn-Golang-p1/defer/logs.log", "/home/alireza/Desktop/Learn-Golang-p1/defer/BackUplogs.log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The Result is: %v \n", Written)
}
func File(filesource string, fileDestination string) (int64, error) {
	//defer is running after all codes
	//you can handle panic error with defer
	//when you have panic only defer can be run
	if len(filesource) <= 3 {
		return 0, errors.New("Errors in file source")
	}
	if len(fileDestination) <= 3 {
		return 0, errors.New("Errors in file destinations")
	}
	OpeningFileSource, err := os.Open(filesource)
	if err != nil {
		return 0, err
	}
	defer OpeningFileSource.Close()
	DestinationFileSource, err := os.Create(fileDestination)
	if err != nil {
		return 0, err
	}
	defer DestinationFileSource.Close()
	return io.Copy(DestinationFileSource, OpeningFileSource)
}
