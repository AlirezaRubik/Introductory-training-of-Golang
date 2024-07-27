package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan string)
	go func() {
		SetChanels(chanel)
	}()
    fmt.Println(<-chanel)
	fmt.Println(<-chanel)
	fmt.Println(<-chanel)
	fmt.Println(<-chanel)
	time.Sleep(time.Second*1)
}
func SetChanels(input chan<- string) {
	input <- "test"
	input <- "test1"
	input <- "test2"
	input <- "test3"
}
