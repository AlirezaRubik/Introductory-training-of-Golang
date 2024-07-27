package main
import(
	"encoding/json"
	"fmt"
)
type Person struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Age    float64 `json:"age"`
	Weight float64 `json:"weight"`
}

func main() {
	// we want convert the person result to json for show too users
	//make person
	p1 := Person{Id: 1, Name: "Alireza", Age: 22, Weight: 112}
    
	//convert to json
	JsonRes,err:=json.Marshal(p1)
	//handle errors
	if err!=nil{
        fmt.Println(err)
    }
	//show json with %s=string
    fmt.Printf("%s \n",JsonRes)

}
