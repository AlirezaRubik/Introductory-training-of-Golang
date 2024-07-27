package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Age    float64 `json:"age"`
	Weight float64 `json:"weight"`
}

func main() {
	//we save the json codes in value with [``]
	JsonEncoder := `{"id":1,"name":"Alireza","age":22,"weight":112}`
    //make empty value of struct
	p1 := new(Person)
	//unmarshal json to struct with Unmarshal function  from encoding/json package
	json.Unmarshal([]byte(JsonEncoder), p1)
	//show the values of struct with fmt.Printf function
	fmt.Printf("The Id is: %v\n", p1.Id)
	fmt.Printf("The Name is: %v\n", p1.Name)
	fmt.Printf("The Age is: %v\n", p1.Age)
}
