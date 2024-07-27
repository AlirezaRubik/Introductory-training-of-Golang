package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan int)
	for i := 0; i < 100; i++ {
		go AddNum(i, chanel)
	}
	time.Sleep(time.Second * 1)
	for i := 0; i < 100; i++ {
		Response := <-chanel
		fmt.Println(Response)
	}
}
func AddNum(index int, chanel chan int) {
	chanel <- index
}
