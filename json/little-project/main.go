package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BookingLocation struct {
	Country   string `json:"country"`
	State     string
	StateName string
	Zipcode   string
	Timezone  string
	Latitude  string
	Longitude string
	City      string
	Continent string
}

func main() {
	chanel:=make(chan string)
	for i := 0; i < 100; i++ {
       go GetRequest(chanel)
	}
	b:=new(BookingLocation)
	time.Sleep(time.Second*1)
	for i := 0; i < 100;i++{
		Result:=<-chanel
		fmt.Println(Result)
		json.Unmarshal([]byte(Result),b )
		fmt.Println(b.City)
		fmt.Println(b.StateName)
		fmt.Println(b.Timezone)
		fmt.Println(b.Continent)
	}
}
func GetRequest(chanel chan string) {
	Client := http.Client{}
	Request, err := http.NewRequest("GET", "https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location", nil)
	if err != nil {
		fmt.Println(err)
	}
	Request.Header = http.Header{}
	Request.Header.Add("accept:", "application/json")
	ClientResult, err := Client.Do(Request)
	if err != nil {
		panic(err)
	}
	ResponseBody, err := ioutil.ReadAll(ClientResult.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ResponseBody))
	chanel <- string(ResponseBody)

}