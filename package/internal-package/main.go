package main

import (
	"fmt"

	jalaali "github.com/jalaali/go-jalaali"
)

func main() {
	//in here we use the go jalaali package with internal system 
	//we change the go mod setting 
	//from 
	//require github.com/jalaali/go-jalaali v0.0.0-20210801064154-80525e88d958
	//to
    //replace github.com/jalaali/go-jalaali => ./package/internal-package/go-jalaali
	year, month, day, err := jalaali.ToGregorian(1383, 222, 222)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The Year is %v \n", year)
	fmt.Printf("The Month is %v \n", month)
	fmt.Printf("The Day is %v \n", day)
}
