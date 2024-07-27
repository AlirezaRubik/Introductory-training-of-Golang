package main

import (
	"fmt"
)

type IAnimal interface {
	Sleep()
	Eat()
	Speak()
}
type IHuman interface {
	IAnimal
	Think()
}
type Cat struct {
	Id int
}
type Human struct {
	Name string
}

func main() {
	var Animals IAnimal
	Cats := new(Cat)
	Cats.Id = 101
	Animals = Cats
    Animals.Eat()
	Animals.Sleep()
	Animals.Speak()

	var Humans IHuman
	Humans = &Human{Name: "Alireza"}
	Humans.Eat()
	Humans.Sleep()
	Humans.Speak()
	Humans.Think()

}
func (c *Cat) Sleep() {
	fmt.Println("Cat With This Id:", c.Id, "is sleeping")
}
func (c *Cat) Eat() {
	fmt.Println("Cat With This Id:", c.Id, "is eating")
}
func (c *Cat) Speak() {
	fmt.Println("Cat With This Id:", c.Id, "is meowing")
}

func (H *Human) Sleep() {
	fmt.Println(H.Name + " sleeping")
}
func (H *Human) Eat() {
	fmt.Println(H.Name + " Eating")
}
func (H *Human) Speak() {
	fmt.Println(H.Name + " Speaking")
}
func (H *Human) Think() {
	fmt.Println(H.Name + " Thinking")
}

