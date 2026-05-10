package main

import (
	"fmt"
	"reflect"
)
func main(){
	var name string 
	res:=reflect.TypeOf(name).String()
    fmt.Println("type of name:",res)
}