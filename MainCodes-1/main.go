package main

import (
	"encoding/json"
	"fmt"
	//"strings"
	//"unsafe"
	//"time"
)

type Order struct {
	Id     int
	Name   string
	Price  float64
	Status SetStatus
}

type SetStatus int

const (
	WaitingPayment SetStatus = 0
	GoodPayment    SetStatus = 1
	BadPayment     SetStatus = -1
	RefoundPayment SetStatus = -2
)

func main() {
	//1-printing show infors
	//fmt.Println("Hello world! This is Println")
	//fmt.Print("Hello world! This is Print \n")
	// fmt.Printf("Hello world! This is Printf %d",2) //help: %v:values|%d:int|%s:string|%c:character|%f:float|%T:Type

	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	//2-size of different types of values
	// var Num int
	// var Num2 int8
	// var Num3 int16
	// var Num4 int32
	// var Num5 int64
	// var Float float32
	// var Float64 float64
	// var Str string
	// fmt.Printf("Size of int: %d \n", unsafe.Sizeof(Num))
	// fmt.Printf("Size of int8: %d \n", unsafe.Sizeof(Num2))
	// fmt.Printf("Size of int16: %d \n", unsafe.Sizeof(Num3))
	// fmt.Printf("Size of int32: %d \n", unsafe.Sizeof(Num4))
	// fmt.Printf("Size of int64: %d \n", unsafe.Sizeof(Num5))
	// fmt.Printf("Size of float32: %d \n", unsafe.Sizeof(Float))
	// fmt.Printf("Size of float64: %d \n", unsafe.Sizeof(Float64))
	// fmt.Printf("Size of string: %d \n", unsafe.Sizeof(Str))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//3-different loops
	//////////////
	//3.1 infinite loops
	// count:=0
	// for{
	// 	 fmt.Println("Hello",count)
	//      count++
	// 	 time.Sleep(time.Second*1)
	// }
	//////////////
	//3.2 for in arrays
	// arr:=[5]int{1,2,3,4,5}
	// for i := 0; i < len(arr); i++{
	// 	fmt.Println(arr[i])
	// }
	//or
	// for index,item:=range arr{
	// 	fmt.Printf("Index: %d, Item: %d\n",index,item)
	// }
	//////////////
	//3.3 loop with if
	// for i:=0; i<100; i++{
	//    fmt.Println("The Index is:",i)
	// }
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//4-conditions
	//switch case
	//4.1
	// fmt.Println("Please Enter The Number of the week:")
	// var input int
	// fmt.Scanln(&input)
	// switch input{
	// case 0:
	// 	fmt.Println("Monday")
	// case 1:
	// 	fmt.Println("Tuesday")
	// case 2:
	// 	fmt.Println("Wednesday")
	// case 3:
	// 	fmt.Println("Thursday")
	// case 4:
	// 	fmt.Println("Friday")
	// case 5:
	// 	fmt.Println("Saturday ")
	// case 6:
	// 	fmt.Println("Sunday ")
	// }
	//////////////
	//4.2
	// fmt.Println("Please Enter The Score:")
	// var score float64
	// fmt.Scanln(&score)
	// switch{
	// case score<=20 &&score>=17:
	// 	fmt.Println("Group A")
	// case score<=16&&score>=12:
	// 	fmt.Println("Group B")
	// case score<=11 &&score>=0:
	// 	fmt.Println("Group C")
	// default:
	// 	fmt.Println("Error: Invalid Score")
	// }
	//////////////
	//4.3
	//fallthrough=>when we want use fallthrough we start from end
	//month 1-12 ===> in fallthrough= 12-1
	//if the conditions are true then next condition are work ==>example 12 is true => 11 is working 10 is working until 1
	// fmt.Println("Please Enter The Number of The Month:")
	// var month int
	// var days=0
	// fmt.Scanf("%d",&month)
	// switch month{
	// case 12:
	// 	days+=29
	// 	fallthrough
	// case 11:days+=30
	// 	fallthrough
	// case 10:
	// 	days+=30
	// 	fallthrough
	// case 9:
	// 	days+=30
	// 	fallthrough
	// case 8:
	// 	days+=30
	// 	fallthrough
	// case 7:
	// 	days+=30
	// 	fallthrough
	// case 6:
	// 	days+=31
	// 	fallthrough
	// case 5:
	// 	days+=31
	// 	fallthrough
	// case 4:
	// 	days+=31
	// 	fallthrough
	// case 3:
	// 	days+=31
	// 	fallthrough
	// case 2:
	// 	days+=31
	// 	fallthrough
	// case 1:
	// 	days+=31
	// }
	// fmt.Println("The Total Days are:",days)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//5-if and else
	// fmt.Println("Please Enter Your Name:")
	// var name string
	// var age float64
	// var gender string
	// fmt.Scanln(&name)
	// if len(name)<=2{
	//     fmt.Println("Error In Len Of The Name Please Enter Your Name Correctly")
	// 	return
	// }else{
	// 	fmt.Println("Please Enter Your Age:")
	//     fmt.Scanf("%f",&age)
	// 	if age<=17{
	// 		fmt.Println("Error In Age Count Please Enter Your Age Correctly or You Are Under Age")
	// 		return
	// 	}else{
	//        fmt.Println("Please Enter Your gender (F/M):")
	// 	   fmt.Scanln(&gender)
	// 	   if gender=="f"||gender=="F"||gender=="Female"||gender=="female"||gender=="woman"||gender=="Woman"||gender=="WOMAN"||gender=="FEMALE"{
	//            fmt.Println("Hello Ms",name,"Welcome To Club")
	// 	   }else{
	// 		fmt.Println("Hello Mr",name,"Welcome To Club")
	// 	   }
	// 	}
	// }
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//6-Break and Continue in loops
	// for i:=0;i<100;i++{
	// 	fmt.Println("We Have",i)
	// 	if i==69{
	// 		fmt.Println("I Found 69")
	// 		break
	// 	}
	// }
	//////////////
	// for i:=0;i<100;i++{
	// 	if i%2==0{
	// 		fmt.Println("Continue:",i)
	// 		continue
	// 	}
	// }
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//7-access data values with {block}
	//var name="alireza"
	//{
	//you can access the name value inside this block
	//	fmt.Println(name)
	//	var age =22
	//	fmt.Println(age)
	//in this block we access the age
	//}
	//access name out of this block
	//fmt.Println(name)
	//you cant access the age in here
	//fmt.Println(age)//error
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//8-working with rune
	// ch1:="😂"
	// ch2:="😉"
	// fmt.Printf("The Ch1 Len is:%d The Character of Ch1 is: %c and The Type of The Ch1 is %T and The Size of The Ch1 is: %d and The Ch1 is: %v \n",len(ch1),ch1,ch1,unsafe.Sizeof(ch1),ch1)
	// fmt.Printf("The ch2 Len is:%d The Character of ch2 is: %c and The Type of The ch2 is %T and The Size of The ch2 is: %d and The ch2 is: %v \n",len(ch2),ch2,ch2,unsafe.Sizeof(ch2),ch2)
	//rune is working with string
	//len is working with int
	//sometimes when we want len of the something we get wrong length
	//in rune always we get the right length because rune is working with string
	// Str1:="Hello World 😂😉"
	// fmt.Println("The Len of Str1 is:",len(Str1)) //true len is 14 //but len give me the 20
	// MyRune:=[]rune("Hello World 😂😉")
	// fmt.Println("The Len of Str1 is:",len(MyRune)) //true len is 14
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//9-const values
	//you can never change the const values
	//you can change the normal values of different type
	//but you cant change the const value only in first time you can give value

	//const Url="https://www.rubikcomputer.com"
	//fmt.Println("The Url is:",Url)

	//if you want change the const url you see the error

	//Url="https://"
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//10-pointers
	//pointers working with address of ram=>random memory access
	// var i = 33
	// var j = 100
	// fmt.Printf("The Value of i is:%d and the Type of i is %T and The Size of The i is : %d and The Character of The i is %c \n", i, i, unsafe.Sizeof(i), i)
	// fmt.Printf("The Value of j is:%d and the Type of j is %T and The Size of The j is : %d and The Character of The j is %c \n", j, j, unsafe.Sizeof(j), j)
	// var ip *int //in here we say ip like pointer but pointer to what?
	// var jp *int
	// ip = &i      //ip is pointer to i address &=address, *=pointer
	// jp = &j
	// fmt.Printf("The Value of ip is:%d and the Type of ip is %T and The Size of The ip is : %d and The Character of The ip is %c \n", *ip, *ip, unsafe.Sizeof(*ip), *ip)
	// fmt.Printf("The Value of jp is:%d and the Type of jp is %T and The Size of The jp is : %d and The Character of The jp is %c \n", *jp, *jp, unsafe.Sizeof(*jp), *jp)
	// *[ip = 40000] //remove []
	// *[jp = 90000]//remove []
	// fmt.Printf("[NEW]The Value of ip is:%d and the Type of ip is %T and The Size of The ip is : %d and The Character of The ip is %c \n", *ip, &ip, unsafe.Sizeof(*ip), *ip)
	// fmt.Printf("[NEW]The Value of jp is:%d and the Type of jp is %T and The Size of The jp is : %d and The Character of The jp is %c \n", *jp, &jp, unsafe.Sizeof(*jp), *jp)
	// fmt.Printf("[NEW]The Value of i is:%d and the Type of i is %T and The Size of The i is : %d and The Character of The i is %c \n", i, i, unsafe.Sizeof(i), i)
	// fmt.Printf("[NEW]The Value of j is:%d and the Type of j is %T and The Size of The j is : %d and The Character of The j is %c \n", j, j, unsafe.Sizeof(j), j)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//11-strings
	// const Str="Hello Golang"
	// const Str1="golang is the best programming language"
	// fmt.Println(strings.Title(Str))//show tittle
	// fmt.Println(strings.ToUpper(Str))//get bigger all letters
	// fmt.Println(strings.ToLower(Str))//get lower all letters
	// fmt.Println(strings.Index(Str,"G"))//where is the G in Str? give index
	// fmt.Println(strings.Cut(Str,"Golang"))//cutting
	// fmt.Println(strings.Repeat(Str,10))//repeat something
	// fmt.Println(strings.Replace(Str,"Golang","World",1))//Replace ,string,old string,new string,rnd num
	// fmt.Println(strings.Contains(Str,"Str1"))//is 2 string same?
	// fmt.Println(strings.HasPrefix(Str,"Hello"))//start with?[true]
	// fmt.Println(strings.HasPrefix(Str,"Golang"))//start with?[false]
	// fmt.Println(strings.HasSuffix(Str,"Hello"))//end with? false
	// fmt.Println(strings.HasSuffix(Str,"Golang"))//end with? true
	// fmt.Println(strings.Count(Str1,"g")) //how many we have same argument ?
	// strings.Join([]string{"hello", "world"}, " ")
	// fmt.Println(strings.Trim(Str,"Hello"))//like cut but you are terminating
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//12-how can i give value to different types
	//12.1
	// var num1 int
	// num1=22
	//12.2
	// var num2 =23
	//12.3
	// var Str ="alireza"
	//12.4
	// Num3:=111
	//12.5
	// var f1,Num4,Str1=22.34,55,"Alireza"

	// fmt.Println("[Num]",num1)
	// fmt.Println("[Num]",num2)
	// fmt.Println("[Str]",Str)
	// fmt.Println("[Num3]",Num3)
	// fmt.Println("Float:",f1,"Num4",Num4,"String:",Str1)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//13- call by value and call by reference
	// i := 20
	// fmt.Println("The I Value Count Before Do CallByValue:", i)
	// CallByValue(i)
	// fmt.Println("The I Value Count After Do CallByValue:", i)
	// fmt.Println("The I Value Count Before Do CallByReference:", i)
	// CallByReference(&i)
	// fmt.Println("The I Value Count After Do CallByReference:", i)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//14-working with multi const -struct- type values
	//make some informations from Struct
	FirstOrder:=Order{Id: 1,Name: "Mr Robot Book",Price: 10,Status: GoodPayment}
	SecOrder:=Order{Id: 2,Name: "Labtop",Price: 1230,Status: RefoundPayment}
	LastOrder:=Order{Id: 3,Name: "SteamGiftCard",Price: 100,Status: WaitingPayment}
     
	//convert to json
	JsonResOrder1,err:=json.Marshal(FirstOrder)
	//error provider
	if err!=nil{
		fmt.Println(err)
	}
	JsonResOrder2,err:=json.Marshal(SecOrder)
	if err!=nil{
		fmt.Println(err)
	}
	JsonResOrder3,err:=json.Marshal(LastOrder)
	if err!=nil{
		fmt.Println(err)
	}
	//show to users
	fmt.Printf("%s \n",JsonResOrder1)
	fmt.Printf("%s \n",JsonResOrder2)
	fmt.Printf("%s \n",JsonResOrder3)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
}
func CallByValue(input int) {
	fmt.Println("[Start CallByValue]", input)
	input = 20000
	fmt.Println("[End CallByValue]", input)
}
func CallByReference(input *int) {
	fmt.Println("[Start CallByReference]", *input)
	*input = 20000
	fmt.Println("[End CallByReference]", *input)
}
