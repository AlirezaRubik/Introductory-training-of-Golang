package main
import(
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

func main() {
	// we can send some informations with different methods like builder
	//making person
	p1:=new(Person)
	p1.SetId(1)
	p1.SetName("Alireza")
	p1.SetAge(22)
	p1.SetHairColor("Black")
	p1.SetWeight(112)
	p1.SetHigh(183)
	fmt.Printf("%v \n", p1)
}

// make builder function
func (p *Person) SetId(Id int) *Person {
	p.Id = Id
	return p
}
func (p *Person) SetName(Name string) *Person {
	p.Name = Name
    return p
}
func (p *Person) SetAge(Age int) *Person {
    p.Age = Age
    return p
}
func (p *Person) SetWeight(Weight float64) *Person {
    p.Weight = Weight
    return p
}
func (p *Person) SetHigh(High float64) *Person {
    p.High = High
    return p
}
func (p *Person) SetHairColor(SetHairColor string) *Person {
    p.HairColor = SetHairColor
    return p
}

