package main
import(
  "fmt"
)
func main(){
	//map is like array 
	//map has key and value system 
	//map is faster than array
	//in map you cant change the key but you can change the value 

	//making map
	//we want write salary system
	//name=>slary
	Salary:=make(map[string]float64)

	//insert data
	Salary["Alireza"]=5000
	Salary["Naser"]=2500
	Salary["Amir"]=1100
	fmt.Printf("[INSERT]%v \n",Salary)

	//updating 
	Salary["Alireza"]=5990
	Salary["Naser"]=2000
    Salary["Amir"]=1000
	fmt.Printf("[UPDATE]%v \n",Salary)

	//Delete
	delete(Salary,"Amir")
	fmt.Printf("[DELETE]%v \n",Salary)
	//search

	
	Result,ok:=Salary["Alireza"]
	if ok{
		fmt.Printf("I Found The %v \n",Result)
	}else{
		fmt.Printf("NotFound")
	}


	Result,ok=Salary["Amir"]
	if ok{
		fmt.Printf("I Found The %v \n",Result)
	}else{
		fmt.Printf("NotFound")
	}

}