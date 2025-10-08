package main

import (
	"fmt"
	"runtime"
)

func main() {
	//cpu core count
	fmt.Println(runtime.NumCPU())
	//go routine count
	fmt.Println(runtime.NumGoroutine())
	//os,arch,version
	fmt.Println(runtime.GOOS)
    //go version
	fmt.Println(runtime.Version())
}
