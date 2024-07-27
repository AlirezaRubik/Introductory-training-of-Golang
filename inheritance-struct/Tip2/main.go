package main

import "fmt"

type Thing struct {
	//parent of all struct
	Id     int
	Price  float64
	Name   string
	MadeBy string
	Color  string
	Weight float64
}
type Electronic struct {
	Thing //inheritance from Thing
	Power float64
	Hdd   float64
	Os    string
	Ram   float64
	Cpu   string
	ScreenSize  float64
}
type Phone struct {
	Electronic         // inheritance from Electronic
	SimCardCount       int
	NetWorkTypeSupport string
	Camera             float64
}
type Labtop struct {
	Electronic   //inheritance from Electronic
	SSD          bool
	KeyboardType string
	Webcam       bool
}
func main(){
	l1:=Labtop{SSD: true,KeyboardType: "razer",Webcam: true,Electronic: Electronic{Power: 350,Hdd: 1024,Os: "windows",Ram: 16,Cpu: "core i7 7500",ScreenSize: 16,Thing:Thing{Id: 1,Name: "Asus X541u",Weight: 1200,Color: "black",MadeBy: "china",Price: 1000}}}
	fmt.Printf("[LAPTOP] %v \n",l1)

	p1:=Phone{SimCardCount: 2,NetWorkTypeSupport: "5G",Camera: 200,Electronic:Electronic{Power: 67,Hdd: 1024,Ram: 16,Os: "Android",ScreenSize: 6.9,Thing:Thing{Id: 1,Name: "Samsung S25 Ultra",Weight: 200,Color: "Gray",MadeBy: "China",Price: 1000}}}
	fmt.Printf("[Phone] %v \n",p1)
}