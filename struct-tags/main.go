package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)
type Persons struct {
	Id int `json:"id,omitempty" xml:"ID,omitempty"`
	Name string `json:"NAME,omitempty" xml:"Name,omitempty"`
	Age float64 `json:"-" xml:"Age,omitempty"`
	High float64 `json:"HIGH ,omitempty" xml:"-"`
	Weight float64 `json:"WEIGHT,omitempty" xml:"weight"`
}
func main(){
//struct tags can help us in show jsons 
//struct tags can help us in show xml
//struct tags can help us in validate with binding
//struct tags can help us in gorm validations
p1:=Persons{Id:1,Name: "Alireza",Age:22,High:183,Weight:112}
fmt.Printf("%v\n", p1)

p2:=Persons{Id:2,Name: "Naser",Age:23,High:180}
fmt.Printf("%v\n", p2)

p3:=Persons{Id:3,Age:24,Weight:115}
fmt.Printf("%v \n",p3)


//convert to json
jsonP1,err:=json.Marshal(p1)
if err!=nil{
    fmt.Println(err)
}
fmt.Printf("[JSON PERSON 1]%s \n", jsonP1)


jsonP2,err:=json.Marshal(p2)
if err!=nil{
    fmt.Println(err)
}
fmt.Printf("[JSON PERSON 1]%s \n", jsonP2)

jsonP3,err:=json.Marshal(p3)
if err!=nil{
    fmt.Println(err)
}
fmt.Printf("[JSON PERSON 1]%s \n", jsonP3)

//convert to xml
xmlP1,err:=xml.Marshal(p1)
if err!=nil{
    fmt.Println(err)
}
fmt.Printf("[XML PERSON 1]%s \n", xmlP1)

xmlP2,err:=xml.Marshal(p2)
if err!=nil{
    fmt.Println(err)
}
fmt.Printf("[XML PERSON 1]%s \n", xmlP2)

xmlP3,err:=xml.Marshal(p3)
if err!=nil{
    fmt.Println(err)
}
fmt.Printf("[XML PERSON 1]%s \n", xmlP3)

}