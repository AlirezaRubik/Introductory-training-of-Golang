package main

import (
	"fmt"
	"sort"
)

func main() {
	//func in func =>anonymous function

	//tip 1
	func() {
		fmt.Println("Hello world!")
	}()

	//2-tip2
	Res := func(inputs ...any) {
		for index, item := range inputs {
			fmt.Println("The Index is;", index, "The Item is: ", item)
		}
	}
	Res(1, 2, 3, 4, 5, "ALIREZA", "ALIZADEH", '2', 'c', struct {
		Id   int
		Name string
		Age  float64
	}{Id: 1, Name: "Naser", Age: 22}, []float64{22.3, 11.33, 66, 534.33}, true, false)

	//3-tip3
	Res1 := func(inputs ...interface{}) {
		for index, item := range inputs {
			fmt.Println("The Index is;", index, "The Item is: ", item)
		}
	}

	Res1(1, 2, 3, 4, 5, "ALIREZA", "ALIZADEH", '2', 'c', struct {
		Id   int
		Name string
		Age  float64
	}{Id: 1, Name: "Naser", Age: 22}, []float64{22.3, 11.33, 66, 534.33}, true, false)

	//4-tip4
    fmt.Println(func ()string{
		return "127.8000.1.1"
	}())
    //5-tip5
    Slice:=[]int{4,235,23,4123,41,323,13,252345}
	sort.Slice(Slice, func(i, j int) bool {
		return Slice[i]> Slice[j]
	})
	fmt.Printf("%v \n",Slice)
	//6-Tip6
	func(input int){
        fmt.Println(input)
	}(232)
	//6-Tip6
	func (inputs ...int){
       for _,item:=range inputs{
           fmt.Println(item)
       }
	}(22,2,31,31,3,12,1245,23,565,745,34,524,321)
}
