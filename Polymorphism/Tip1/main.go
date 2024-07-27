package main

import "fmt"

const (
	FullTimeWorkerSalary    = 3500
	PartTimeWorkerSalary    = 1500
	ShiftedTimeWorkerSalary = 1080
	ExtraTimeWorkerSalary   = 500
)

type FullTimeWorkerStruct struct {
	Id           int
	NationalCode string
	Name         string
	Age          float64
	Hours        float64
}
type ShiftTimeWorkerStruct struct {
	Id           int
	NationalCode string
	Name         string
	Age          float64
	Hours        float64
}
type PartTimeWorkerStruct struct {
	Id           int
	NationalCode string
	Name         string
	Age          float64
	Hours        float64
}

type IEmployee interface {
	Workers(Hours float64) (Salary, Tax, Payment float64)
}

func main() {
	fullTimeEmployees := []FullTimeWorkerStruct{
		{Id: 1, NationalCode: "1234567890", Name: "Pejman Rezaee", Hours: 40, Age: 22},
		{Id: 2, NationalCode: "4836524125", Name: "Maryam Hosseini", Hours: 120, Age: 21},
	}

	partTimeEmployees := []ShiftTimeWorkerStruct{
		{Id: 3, NationalCode: "6563453455", Name: "Milad Hassani", Hours: 100, Age: 33},
		{Id: 4, NationalCode: "5435435435", Name: "Maryam Rezaee", Hours: 87, Age: 21},
	}

	shiftEmployees := []PartTimeWorkerStruct{
		{Id: 5, NationalCode: "3123123213", Name: "Shahin", Hours: 150, Age: 55},
		{Id: 6, NationalCode: "6363454355", Name: "Masoud", Hours: 168, Age: 49},
	}
	for _, item := range fullTimeEmployees {
		Salary, Tax, Payment := item.Workers()
		fmt.Printf("[FulltimeWorker]The Salary: %f, Tax: %f, Payment: %f\n", Salary, Tax, Payment)
	}
	for _, item := range shiftEmployees {
		Salary, Tax, Payment := item.Workers()
		fmt.Printf("[ShiftTimeWorker]The Salary: %f, Tax: %f, Payment: %f\n", Salary, Tax, Payment)
	}
	for _, item := range partTimeEmployees {
		Salary, Tax, Payment := item.Workers()
		fmt.Printf("[PartTimeWorker]The Salary: %f, Tax: %f, Payment: %f\n", Salary, Tax, Payment)
	}
}
func (e *FullTimeWorkerStruct) Workers() (Salary, Tax, Payment float64) {
	Salary = FullTimeWorkerSalary + (e.Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}
func (e *ShiftTimeWorkerStruct) Workers() (Salary, Tax, Payment float64) {
	Salary = ShiftedTimeWorkerSalary + (e.Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}
func (e *PartTimeWorkerStruct) Workers() (Salary, Tax, Payment float64) {
	Salary = PartTimeWorkerSalary + (e.Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}
