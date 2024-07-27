package main

import (
	"fmt"
	"sync"
)
type Config struct{

}
var(
	Configs *Config
	Once sync.Once
)
func Set(){
	Once.Do(func() {
		if Configs == nil{
			Configs = &Config{}
			fmt.Println("The NewConfig Has Been Set")
		}else{
			fmt.Println("The Old Config")
		}
	})
}
func main(){
    for i:=0;i<100;i++{
        fmt.Println("Try",i,"We Have:",)
		Set()
	}
}