package main

import (
	"errors"
	"fmt"
	"time"
)

type Room struct {
	Id           int
	BedCount     int
	Status       bool
	PricePerHour float64
	Tip          string
}

var RoomsList []Room = addDefaultRoom()

func main() {
	for {
		Res := setMenu()
		switch Res {
		case 0:
			return
		case 1:
			showDefaultRooms()
		case 2:
			addNewRoom()
		case 3:
			reserveRoom()
		case 4:
			unreserveRoom()
		}
	}

}
func showDefaultRooms() {
	for _, room := range RoomsList {
		fmt.Printf("ID: %d, Bed Count: %d, Status: %t, Price Per Hour: %.2f, Tip: %s \n", room.Id, room.BedCount, room.Status, room.PricePerHour, room.Tip)
	}
}
func addDefaultRoom() []Room {
	var DefaultRoomsList []Room
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 1, BedCount: 1, Status: false, PricePerHour: 200, Tip: "Standard"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 2, BedCount: 2, Status: false, PricePerHour: 250, Tip: "Sweat"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 3, BedCount: 3, Status: false, PricePerHour: 300, Tip: "Double"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 4, BedCount: 4, Status: false, PricePerHour: 200, Tip: "Standard"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 5, BedCount: 5, Status: false, PricePerHour: 250, Tip: "Sweat"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 6, BedCount: 4, Status: false, PricePerHour: 300, Tip: "Double"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 7, BedCount: 3, Status: false, PricePerHour: 200, Tip: "Standard"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 8, BedCount: 2, Status: false, PricePerHour: 250, Tip: "Sweat"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 9, BedCount: 1, Status: false, PricePerHour: 300, Tip: "Double"})
	DefaultRoomsList = append(DefaultRoomsList, Room{Id: 10, BedCount: 4, Status: false, PricePerHour: 200, Tip: "Standard"})
	return DefaultRoomsList
}
func setMenu() int {
	fmt.Println("1-Rooms List")
	fmt.Println("2-Add New Room")
	fmt.Println("3-Reserve Room")
	fmt.Println("4-UnReserve Room")
	fmt.Println("0-Exit")
	fmt.Println("Please Enter Input:")
	var input int
	fmt.Scanln(&input)
	return input
}
func GetId() (int, error) {
	fmt.Println("Please Enter The Room Id:")
	var id int
	fmt.Scanln(&id)
	if id <= -1 {
		return 0, errors.New("Error In Count of Room")
	}
	return id, nil
}

func GetBedCount() (int, error) {
	fmt.Println("Please Enter The Room Beds Count:")
	var beds int
	fmt.Scanln(&beds)
	if beds <= -1 {
		return 0, errors.New("Error In Count Beds")
	}
	return beds, nil
}
func GetNightCount() (int, error) {
	fmt.Println("Please Enter The Room Nights Count:")
	var Nights int
	fmt.Scanln(&Nights)
	if Nights <= -1 {
		return 0, errors.New("Error In Count Nights")
	}
	return Nights, nil
}
func tipMenue() string {
	fmt.Println("******* PLEASE ENTER ROOM TIP *******")
	fmt.Println("1-Standard")
	fmt.Println("2-Sweat")
	fmt.Println("3-Double")
	fmt.Println("Please Enter Your Choice:")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 1:
		return "Standard"
	case 2:
		return "Sweat"
	case 3:
		return "Double"
	}
	return "Standard"
}

func reserveRoom() {
	id, err := GetId()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(RoomsList); i++ {
		if RoomsList[i].Id == id {
			if RoomsList[i].Status == true {
				fmt.Println("This Room Is Already Reserved")
				time.Sleep(time.Second * 3)
				return
			} else {
				RoomsList[i].Status = true
				fmt.Println("Room Is Reserved Successfully")
				Nights,err:=GetNightCount()
				if err!= nil {
                    fmt.Println(err)
                    return
                }
				pay,tax,payment:=PriceCalculate(Nights,RoomsList[i].PricePerHour)
				if pay==0 ||tax==0 ||payment==0 {
					fmt.Println("Error In Reserve Room")
					RoomsList[i].Status = false
					time.Sleep(time.Second * 3)
					return
				}
				fmt.Println("The Pay Price is:",pay,"The Tax is:",tax,"The Payment is:",payment)
				time.Sleep(time.Second * 3)
				return
			}
		}
	}
}
func unreserveRoom() {
	id, err := GetId()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(RoomsList); i++ {
		if RoomsList[i].Id == id {
			if RoomsList[i].Status == false {
				fmt.Println("This Room Is Already UnReserved")
				time.Sleep(time.Second * 3)
				return
			} else {
				RoomsList[i].Status = false
				fmt.Println("Room Is UnReserved Successfully")
				time.Sleep(time.Second * 3)
				return
			}
		}
	}
}
func addNewRoom() {
	id, err := GetId()
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 3)
		return
	}
	for i := 0; i < len(RoomsList); i++ {
		if RoomsList[i].Id == id {
			fmt.Println("Room With This Id Already Exists")
			time.Sleep(time.Second * 3)
			return
		}
	}
	beds, err := GetBedCount()
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 3)
		return
	}
	roomTip := tipMenue()

	Price := 200.00
	switch roomTip {
	case "Standard":
		Price = 200.00
	case "Sweat":
		Price = 250.00
	case "Double":
		Price = 300.00
	}
	RoomsList = append(RoomsList, Room{Id: id, BedCount: beds, Status: false, Tip: roomTip, PricePerHour: Price})
}
func PriceCalculate(nights int, price float64) (pay, tax, payment float64) {
	if price <=0 ||nights <= 0 {
		return 0,0,0
	}
	pay = float64(nights) * price
	tax = pay * 0.9
	payment = pay + tax
	return pay, tax, payment
}
