package main
import(
"fmt"
)
func main(){
//1-
sayHello()
//2-
Sumer(2,55)
//3-
fmt.Println(GetIp())
//4-
FullName:=PrintName("Alireza Alizadeh Aghdam")
fmt.Println(FullName)
//5-
MultiInput(1,2,3,4,5,"ALIREZA","ALIZADEH",'2','c',struct{Id int;Name string;Age float64}{Id: 1,Name: "Naser",Age: 22},[]float64{22.3,11.33,66,534.33},true,false)
//6-
NewMultiInputs(1,2,3,4,5,"ALIREZA","ALIZADEH",'2','c',struct{Id int;Name string;Age float64}{Id: 1,Name: "Naser",Age: 22},[]float64{22.3,11.33,66,534.33},true,false)
//7-
salary,tax,payment:=CalculateHotelPrice(5,220)
fmt.Println("Salary: ",salary,"Tax: ",tax,"Payment: ",payment)
}
//1-function without any inputs and outputs
func sayHello(){
    fmt.Println("Hello World")
}

//2-function with inputs and outputs
func Sumer(a,b int){
	fmt.Println(a+b)
}

//3-function without input with outputs
func GetIp()string{
	return "127.8000.1.1"
}

//4-function with input and no outputs
func PrintName(name string)string{
    return name
}
//5-function with many inputs
func MultiInput(inputs ...any){
	for index,item:=range inputs{
		fmt.Println("The Index is: ",index,"The Item is: ",item)
	}
}
//6-function with many inputs
func NewMultiInputs(inputs ...interface{}){
	for index,item:=range inputs{
        fmt.Println("The Index is: ",index,"The Item is: ",item)
    }
}
//7-multi input and multi outputs
func CalculateHotelPrice(Nights float64,PriceForNight float64)(salary , tax,payment float64){
	salary=Nights*PriceForNight
	tax=salary*0.9
	payment=salary+tax
	return salary,tax,payment
}

