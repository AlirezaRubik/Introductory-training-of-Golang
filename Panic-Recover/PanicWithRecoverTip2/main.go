package main

import (
	"fmt"
	"log"
)
func main(){
	//1-
	//arrays:=[]int{3,4,6,13,5}
	//Array(arrays,12)
	//2-
	//PanicMaker()
    //3-
	//PanicLoger()
	//4-
	//Divide(10,2)
	//Divide(10,0)//panic
	//5-
	//DivideDefer(10,2)
	//DivideDefer(10,0)//panic and ending with defer
	//6-
	DivideDeferCover(10,2)
	DivideDeferCover(10,0)//Balanced by defer and recover
}
//panic when we want index out of range array
//1-
func Array(arr[]int,index int){
fmt.Println(arr[index])
}
//2-
//Making Panic in Functions
func PanicMaker(){
	panic("Hello Im Panic")
}
//3-
//Making Panic in Log
func PanicLoger(){
	log.Fatal("Hello Im Panic")
}
//4-
func Divide(a,b int){
	fmt.Println("Starting Divide")
	fmt.Println(a/b)
	fmt.Println("Ending Divide")
}
//5-
func DivideDefer(a,b int){
	defer fmt.Println("Ending Divide With Defer")
	fmt.Println("Starting Divide")
	fmt.Println(a/b)
	fmt.Println("Ending Divide")
}
//6-
func DivideDeferCover(a,b int){
	defer func(){
		fmt.Println("Ending Divide With Defer")
		err:=recover()
		if err!=nil{
            fmt.Println(err)
			fmt.Println("Recover With Recover")
        }
	}()
	fmt.Println("Starting Divide")
	fmt.Println(a/b)
	fmt.Println("Ending Divide")
}