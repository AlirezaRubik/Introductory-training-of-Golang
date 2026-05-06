
package main

import (
	"fmt"
)

type ItemsList[T comparable] struct {
	Items []T
}

func main() {
	var itm = new(ItemsList[int])
	itm.AddNewItem(11)
	itm.AddNewItem(22)
	itm.AddNewItem(33)
	itm.AddNewItem(44)
	itm.AddNewItem(55)
	itm.AddNewItem(66)

	itm.ShowAllItems()

	itm.AddItemWithIndex(2111, 2)
	itm.AddItemWithIndex(8585, 5)

	itm.ShowAllItems()

	itm.RemoveWithIndex(2111)
	itm.RemoveWithIndex(11)

	itm.ShowAllItems()
}
func (itm *ItemsList[T]) AddNewItem(input T) {
	itm.Items = append(itm.Items, input)
}
func (itm *ItemsList[T]) ShowAllItems() {
	fmt.Printf("\n")
	for index, item := range itm.Items {
		fmt.Println("index", index, "item:", item)
	}
}
func (itm *ItemsList[T]) AddItemWithIndex(inputs T, index int) {
	itm.Items = append(itm.Items, inputs)
	copy(itm.Items[index+1:], itm.Items[index:])
	itm.Items[index] = inputs
}

func (itm *ItemsList[T]) RemoveWithIndex(inputs T) {
	for index, item := range itm.Items {
		if item == inputs {
			itm.Items = append(itm.Items[:index], itm.Items[index+1:]...)
		}
	}
}