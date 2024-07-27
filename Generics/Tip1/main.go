package main

import "fmt"
type Iinputes interface{
	int|float64|bool|struct{}|[]interface{}|string
}
func main(){
 //test func SumI
 ResultI:=SumI(2,4)
 fmt.Println("SumI=>",ResultI)
 //test func Sumf
 ResultF:=SumF(2.4356,99.09)
 fmt.Println("SumF=>",ResultF)
 //test Generic Functions
 Inputes(true,false,true,true,true,true)
 Inputes(2,4,1,44.334,11.345,33)
 Inputes("Alireza","Alizadeh")
}
//we write 2 different functions for sum of two different types of value
func SumI(a,b int)int{
	return a+b
}
func SumF(a,b float64)float64{
	return a+b
}
//sometimes we need some functions with supporting many types of values
//in this problems we use the generics 
func SumerGeneric[T int|float64](a,b T)T{
	return a+b

}
//in Generics We Can Use With Interface
func Inputes[T Iinputes](inputsList ...T){
	for index,item:=range inputsList{
		fmt.Println("The Index is:",index,"The Item is:",item)
	}
}
