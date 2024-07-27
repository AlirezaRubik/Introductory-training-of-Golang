package main

import "fmt"

// Person represents a person with various attributes.
type Person struct {
	Id        int
	Name      string
	Age       float64
	Weight    float64
	HairColor string
	High      float64
}

func main() {
	// Test case 1: Invalid Id (-1)
	PT1 := RegisterNewPerson(-1, "Alireza", 22, 112, "Black", 183)
	fmt.Printf("%v \n", PT1)
	// Test case 2: Invalid Name (less than 3 characters)
	PT2 := RegisterNewPerson(2, "Na", 22, 112, "Yellow", 120)
	fmt.Printf("%v \n", PT2)
	// Test case 3: Invalid Age (0)
	PT3 := RegisterNewPerson(3, "Mohsen", 0, 87, "Orange", 167)
	fmt.Printf("%v \n", PT3)
	// Test case 4: Invalid Weight (0)
	PT4 := RegisterNewPerson(3, "Mohsen", 16, 0, "Brown", 167)
	fmt.Printf("%v \n", PT4)
	// Test case 5: Invalid HairColor (less than 3 characters)
	PT5 := RegisterNewPerson(3, "Mohsen", 21, 100, "GE", 189)
	fmt.Printf("%v \n", PT5)
	// Test case 6: Invalid High (-165)
	PT6 := RegisterNewPerson(3, "Mohsen", 21, 100, "GE", -165)
	fmt.Printf("%v \n", PT6)
}

// RegisterNewPerson creates a new Person object with the given parameters.
// It validates the input parameters and panics if any of them are invalid.
func RegisterNewPerson(Id int, Name string, Age float64, Weight float64, HairColor string, High float64) *Person {
	defer func() {
		Resutl := recover()
		fmt.Println(Resutl)
	}()
	if Id <= 0 {
		panic("Invalid Id <=0")
	} else {
		if len(Name) <= 2 {
			panic("Invalid Name length <=2")
		} else {
			if Age <= 0 {
				panic("Invalid Age <=0")
			} else {
				if Weight <= 0 {
					panic("Invalid Weight <=0")
				} else {
					if len(HairColor) <= 2 {
						panic("Invalid HairColor length <=2")
					} else {
						if High <= 0 {
							panic("Invalid High <=0")
						} else {
							return &Person{Id: Id, Name: Name, Age: Age, Weight: Weight, HairColor: HairColor, High: High}
						}
					}
				}
			}
		}
	}
}
