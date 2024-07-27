package main

import (
	"fmt"
)

type Person struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Weight    float64 `json:"weight"`
	High      float64 `json:"high"`
	HairColor string  `json:"hairColor"`
}
type Fperson func(*Person)

func main() {
	Response:=Constractor(SetId(1),SetName("Alireza"),SetAge(22),SetHairColor("Black"),SetWeight(112),SetHigh(183))
	fmt.Printf("%v \n",Response)
}
func Constractor(Options ...Fperson)*Person{
	p1:=new(Person)
	for _,item:=range Options{
        item(p1)
	}
	return p1
}
func SetId(Id int) Fperson{
	return func (p *Person)  {
		p.Id = Id
	}
}
func SetName(Name string) Fperson {
	return func(p *Person) {
		p.Name = Name
	}
}
func SetAge(Age int) Fperson {
	return func(p *Person) {
		p.Age = Age
	}
}
func SetWeight(Weight float64) Fperson {
	return func(p *Person) {
		p.Weight = Weight
	}
}
func SetHigh(hight float64) Fperson {
	return func(p *Person) {
		p.High = hight
	}
}
func SetHairColor(HairColor string) Fperson {
	return func(p *Person) {
		p.HairColor = HairColor
	}
}
