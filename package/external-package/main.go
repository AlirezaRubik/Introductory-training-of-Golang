package main

import (
	"fmt"

	jalaali "github.com/jalaali/go-jalaali"
)
func main(){
    year,month,day,err:=jalaali.ToGregorian(1383,04,18)
	if err!=nil{
        panic(err)
    }
	fmt.Printf("The Year is %v \n",year)
	fmt.Printf("The Month is %v \n",month)
	fmt.Printf("The Day is %v \n",day)
}