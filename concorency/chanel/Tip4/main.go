package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			SetChanel(chanel, i)
		}()
		
	}
	time.Sleep(time.Second * 1)
	for i := 0; i<10;i++{
		Result:=<-chanel
		fmt.Println(Result)
	}
}
func SetChanel(chanel chan int, index int) *chan int {
	chanel <- index
	return &chanel
}