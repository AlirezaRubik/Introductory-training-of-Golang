package main

import (
	"fmt"
   Set "informations/concorency/sync.Once/BadExample/Set"
)
func main(){
	for i:=0;i<1000;i++{
		fmt.Println("Try",i,"We Have:",)
		Set.CheckStatus()
	}
}