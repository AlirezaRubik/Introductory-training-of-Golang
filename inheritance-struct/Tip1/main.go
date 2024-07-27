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

func main() {
	//making labtop struct
	l1 := new(Labtop)
	l1.Id = 1
	l1.Color = "black"
	l1.Cpu = "core i7 7500u"
	l1.Name = "asus x541 u"
	l1.Hdd = 1024
	l1.MadeBy = "china"
	l1.KeyboardType = "razer"
	l1.Power = 320
	l1.Ram = 16
	l1.Os = "windows"
	l1.SSD = true
	l1.Webcam = true
	l1.Weight=1200
	l1.ScreenSize=16
	l1.Price=1000
	fmt.Printf("%v \n", l1)

	//making Phone struct
	p1 := new(Phone)
	p1.Id = 1
	p1.Color = "gray"
	p1.Cpu = "snapp dragon GN4"
	p1.Name = "Samsung S25Ultra"
	p1.Hdd = 1024
	p1.MadeBy = "china"
	p1.Power = 67
	p1.Ram = 16
	p1.Os = "android 15"
	p1.Camera=200
	p1.NetWorkTypeSupport="5G"
	p1.ScreenSize=6.9
	p1.Weight=290
	p1.Price=1000
	p1.SimCardCount=2
	fmt.Printf("%v \n", p1)
}
