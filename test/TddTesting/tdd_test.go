package TestTdd

import (
	"fmt"
	"testing"
)

func TestSetSalary(t *testing.T) {
	//arrange
	BaseSalary := 10000000.0
	ExtraHourse := 12.0
	want := 800000.0
	//act
	got := SetSalary(BaseSalary, ExtraHourse)
	//assert
	if got != want {
		t.Errorf("[Salary]got %v, want %v", got, want)
	}
}
func TestSetInsurance(t *testing.T) {
	//arrange
	BaseSalary := 10000000.0
	ExtraHourse := 800000.0
	want := 324000.0
	//act
	got := SetInsurance(BaseSalary, ExtraHourse)
	//assert
	if got != want {
		t.Errorf("[Insurance]got %v, want %v", got, want)
	} else {
		fmt.Printf("t: %v\n", got)
	}

}
func TestSetTax(t *testing.T) {
	//arrange
	BaseSalary := 10000000.0
	ExtraHourse := 800000.0
	want := 972000.0
	//act
	got := SetTax(BaseSalary, ExtraHourse)
	//assert
	if got != want {
		t.Errorf("[Insurance]got %v, want %v", got, want)
	}
}
func TestFinalSalary(t *testing.T) {
	BaseSalary := 10000000.0
	ExtraHouse := 800000.0
	Karaneh := 5000000.0
	want := 14504000.0
	got := FinalSalary(BaseSalary, ExtraHouse, Karaneh)
	if got != want {
		t.Errorf("[TestFinalSalary]got %v, want %v", got, want)
	}
}
func TestFinalSalary2(t *testing.T) {
	BaseSalary := 10000000.0
	ExtraHouse := 600000.0
	Karaneh := 5000000.0
	want := 14504000.0
	got := FinalSalary(BaseSalary, ExtraHouse, Karaneh)
	if got != want {
		t.Errorf("[TestFinalSalary]got %v, want %v", got, want)
	}
}