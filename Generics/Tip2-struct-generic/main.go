package main

import "fmt"

type Item[T int] struct {
	ItemsList []T
}

func main() {
Items:=new(Item[int])
Items.InsertData(22)
Items.InsertData(34522)
Items.InsertData(67)
Items.InsertData(25672)
Items.InsertData(567)
Items.InsertData(345)
Items.InsertData(657)
Items.InsertData(2345342)

Items.ShowAllInfos()

Items.InsertWithIndex(8585,3)
Items.InsertWithIndex(6969,6)
Items.InsertWithIndex(8569,0)

Items.ShowAllInfos()

Items.DeleteData(6969)
Items.DeleteData(25672)
Items.DeleteData(345)
Items.DeleteData(22)
Items.DeleteData(657)

Items.ShowAllInfos()

}
func (Itm *Item[T]) ShowAllInfos() {
	fmt.Printf("\n \n")
	for index, item := range Itm.ItemsList {
		fmt.Println("The Index is:", index, "The Item is:", item)
	}
}
func (Itm *Item[T]) InsertData(input T) {
	Itm.ItemsList = append(Itm.ItemsList, input)
}
func (Itm *Item[T])InsertWithIndex(input T, index int){
	Itm.ItemsList=append(Itm.ItemsList, input)
	copy(Itm.ItemsList[index+1:], Itm.ItemsList[index:])
	Itm.ItemsList[index]=input
}
func (Itm *Item[T])DeleteData(input T) {
	for i:=0;i<len(Itm.ItemsList);i++{
		if Itm.ItemsList[i]==input{
			Itm.ItemsList=append(Itm.ItemsList[:i], Itm.ItemsList[i+1:]...)
		}
	}
}