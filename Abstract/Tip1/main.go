package main

import "fmt"

//this is the bad exmaple for Abstract
//We Have MultiConst For SetSalarys
//Tip is for Set Employee Type

const (
	FullTimeWorkerSalary    = 3500
	PartTimeWorkerSalary    = 1500
	ShiftedTimeWorkerSalary = 1080
	ExtraTimeWorkerSalary   = 500
)

type Employee struct {
	Id           int
	NationalCode string
	Name         string
	Age          float64
	Hours        float64
	Tip          string
}

func main() {
	//making Array With Employers List
	Employees := []Employee{
		{Id: 1, NationalCode: "23VFE34", Name: "Alireza Alizadeh Aghdam", Age: 22, Tip: "FullTime", Hours: 40},
		{Id: 2, NationalCode: "VDETR324", Name: "Naser Mohammadi", Age: 24, Tip: "ShiftTime", Hours: 22},
		{Id: 3, NationalCode: "YHTYU44", Name: "Mohsen", Age: 18, Tip: "PartTime", Hours: 33.2},
	}
	// Send Informations With For To SetEmployee
	for _, item := range Employees {
		Salary, Tax, Payment := item.SetEmployee(item.Hours)
		fmt.Printf("The Salary: %f, Tax: %f, Payment: %f\n", Salary, Tax, Payment)
	}
}

// FullTimeWorker Function Methods
func (e *Employee) FullTimeWorker(Hours float64) (Salary, Tax, Payment float64) {
	Salary = FullTimeWorkerSalary + (Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}

// ShiftTimeWorker Function Methods
func (e *Employee) ShiftTimeWorker(Hours float64) (Salary, Tax, Payment float64) {
	Salary = ShiftedTimeWorkerSalary + (Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}

// PartTimeWorker FunctionMethods
func (e *Employee) PartTimeWorker(Hours float64) (Salary, Tax, Payment float64) {
	Salary = PartTimeWorkerSalary + (Hours * ExtraTimeWorkerSalary)
	Tax = Salary * 0.9
	Payment = Salary - Tax
	return Salary, Tax, Payment
}

// set Employee type With Methods
func (e *Employee) SetEmployee(Hours float64) (Salary, Tax, Payment float64) {
	if e.Tip == "FullTime" {
		return e.FullTimeWorker(Hours)
	} else if e.Tip == "ShiftTime" {
		return e.ShiftTimeWorker(Hours)
	} else if e.Tip == "PartTime" {
		return e.PartTimeWorker(Hours)
	}
	return 0, 0, 0
}
