package main

import "fmt"

//multi const
const (
	FullTimeWorkerSalary    = 3500
	PartTimeWorkerSalary    = 1500
	ShiftedTimeWorkerSalary = 1080
	ExtraTimeWorkerSalary   = 500
)
//makeing structs for different types of workers
type FullTimeWorkerStruct struct {
	Id          int
	NationalCode string
	Name        string
	Age         float64
	Hours       float64
}
type ShiftTimeWorkerStruct struct {
	Id          int
	NationalCode string
	Name        string
	Age         float64
	Hours       float64
}
type PartTimeWorkerStruct struct {
	Id           int
	NationalCode string
	Name         string
	Age          float64
	Hours        float64
}
//In Real Abstract We Use The Interface
//In Real Abstract We Use Structs for SetEmployee Type
type IEmployee interface {
	Workers(Hours float64) (Salary, Tax, Payment float64)
}

func main() {
	//making fulltime worker like array
	fullTimeEmployees := []FullTimeWorkerStruct{
		{Id: 1, NationalCode: "1234567890", Name: "Pejman Rezaee", Hours: 40,Age: 22},
		{Id: 2, NationalCode: "4836524125", Name: "Maryam Hosseini", Hours: 120,Age: 21},
	}
    //making shifttime worker like array
	partTimeEmployees := []ShiftTimeWorkerStruct{
		{Id: 3, NationalCode: "6563453455", Name: "Milad Hassani", Hours: 100,Age:33},
		{Id: 4, NationalCode: "5435435435", Name: "Maryam Rezaee", Hours: 87,Age: 21},
	}
    //making parttime worker like array
	shiftEmployees := []PartTimeWorkerStruct{
		{Id: 5, NationalCode: "3123123213", Name: "Shahin", Hours: 150,Age: 55},
		{Id: 6, NationalCode: "6363454355", Name: "Masoud", Hours: 168,Age: 49},
	}
	//send the arrays of informations to function
	for _, item := range fullTimeEmployees {
		Salary, Tax, Payment := item.Workers(item.Hours)
		fmt.Printf("[FulltimeWorker]The Salary: %f, Tax: %f, Payment: %f\n", Salary, Tax, Payment)
	}
	for _, item := range shiftEmployees {
		Salary, Tax, Payment := item.Workers(item.Hours)
		fmt.Printf("[ShiftTimeWorker]The Salary: %f, Tax: %f, Payment: %f\n", Salary, Tax, Payment)
	}
	for _, item := range partTimeEmployees {
		Salary, Tax, Payment := item.Workers(item.Hours)
		fmt.Printf("[PartTimeWorker]The Salary: %f, Tax: %f, Payment: %f\n", Salary, Tax, Payment)
	}
}
//function with interface and employee methods
func (e *FullTimeWorkerStruct) Workers(Hours float64) (Salary, Tax, Payment float64) {
	Salary = FullTimeWorkerSalary + (Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}
func (e *ShiftTimeWorkerStruct) Workers(Hours float64) (Salary, Tax, Payment float64) {
	Salary = ShiftedTimeWorkerSalary + (Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}
func (e *PartTimeWorkerStruct) Workers(Hours float64) (Salary, Tax, Payment float64) {
	Salary = PartTimeWorkerSalary + (Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}
//now we dont need selector for check the employee type