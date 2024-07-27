package main

import (
	"fmt"
)

func main() {
	//sometimes we need unique value structers in little time
	p1 := struct {
		Id   int
		Name string
		Age  float64
	}{
		Id:   1,
		Name: "Alireza",
		Age:  22.5,
	}
	fmt.Printf("%v \n", p1)
}
