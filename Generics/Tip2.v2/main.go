
package main

import (
	"encoding/json"
	"fmt"
)

type Items[T any] struct {
	Items []T
}

func main() {
	var itm = new(Items[int])
	itm.AddItem(11)
	itm.AddItem(22)
	itm.AddItem(33)
	itm.AddItem(44)
	itm.AddItem(55)
	itm.AddItem(66)

	itm.ShowAll()

	itm.AddIndex(2111, 2)
	itm.AddIndex(8585, 5)

	itm.ShowAll()

	itm.DeleteValue(2111)
	itm.DeleteValue(11)

	itm.ShowAll()
}
func (itm *Items[T]) ShowAll() {
	fmt.Printf("\n\n")
	for index, item := range itm.Items {
		fmt.Println("The Index is:", index, "Item is:", item)
	}
}
func (itm *Items[T]) AddItem(input T) {
	itm.Items = append(itm.Items, input)
}

func (itm *Items[T]) AddIndex(input T, index int) {
	itm.Items = append(itm.Items, input)
	copy(itm.Items[index+1:], itm.Items[index:])
	itm.Items[index] = input
}

func (itm *Items[T]) DeleteValue(value T) {
	for i:=0;i<len(itm.Items);i++{
		if convertor(itm.Items[i]) == convertor(value) {
			itm.Items = append(itm.Items[:i], itm.Items[i+1:]...)
			break
		}
	}
}
//after 1.21 cmp.Equal deleted 
func convertor[T any](value T) string {
	j1, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error In Convert:", err)
		return ""
	}
	return string(j1)
}

