package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Age  float64
}

func main() {
	//tip 1->make value with struct
	p1 := Person{Id: 1, Name: "Alireza", Age: 22}
	fmt.Printf("[P1]%v \n",p1)
	//tip 2
	p2 := &Person{Id: 2, Name: "Naser", Age: 23}
	fmt.Printf("[P2]%v \n",p2)
	//tip 3
	p3 := new(Person)
	p3.Age = 18
	p3.Id = 3
	p3.Name = "Amir"
	fmt.Printf("[P3]%v \n",p3)
	//tip 4
	Result := CreatePerson(Person{Id: 4, Name: "Mohammad", Age: 22})
	fmt.Printf("[P4]%v \n", Result)
	//tip 5
	p5:=CreateNewPerson(5, "Sadra", 34)
	fmt.Printf("[P5]%v \n",p5)
	ResultsAge:=p5.DoubleAge()
	fmt.Printf("[P5 Method]%v \n", ResultsAge)
	//last method test
	Response:=p2.NewPerson2()
	fmt.Printf("[P2 Method]%v \n", Response)
}
func CreatePerson(p1 Person) *Person {
	return &Person{Id: p1.Id, Name: p1.Name, Age: p1.Age}
}
func CreateNewPerson(Id int, Name string, Age float64) *Person {
	return &Person{Id: Id, Name: Name, Age: Age}
}

// method
func (p *Person) DoubleAge() float64 {
	return p.Age * 2
}
func (p *Person) NewPerson2() *Person{
	return &Person{Id: p.Id,Name: p.Name,Age: p.Age}
}
