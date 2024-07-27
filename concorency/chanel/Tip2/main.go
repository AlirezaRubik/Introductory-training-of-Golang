package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan int)
	go SetChanel(chanel)
	time.Sleep(time.Second * 1)
	Response := <-chanel
	fmt.Println(Response)
	Response = <-chanel
	fmt.Println(Response)
	Response = <-chanel
	fmt.Println(Response)
	Response = <-chanel
	fmt.Println(Response)
}
func SetChanel(chanel chan int) {
	chanel <- 22
	chanel <- 344
	chanel <- 4532
	chanel <- 234
}
