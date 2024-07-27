package main

import (
	"fmt"
)

type Person struct {
	int
	string
	float64
}

// int=id
// string=name
// float64=age
func main() {
	//sometimes we need structs without name for values in struct
	p1 := Person{1, "Alireza", 22.5}
	fmt.Println(p1)
	fmt.Println(p1.int)
	fmt.Println(p1.string)
	fmt.Println(p1.float64)
	//or
	fmt.Printf("%v \n", p1)
}
