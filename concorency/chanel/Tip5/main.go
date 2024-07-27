package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {
	Chanel := make(chan string)

	for i := 0; i < 100; i++ {

		go func() {
			SendRequest(Chanel, i)

		}()
	}
	time.Sleep(time.Second * 1)
	for i := 0; i < 100; i++ {
		Response := <-Chanel
		fmt.Println(Response)
	}
}
func SendRequest(input chan<- string, index int) chan<- string {
	Url := "https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(index)
	Request, err := http.Get(Url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	RequestBody, err := ioutil.ReadAll(Request.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer Request.Body.Close()
	input <- string(RequestBody)
	return input
}
