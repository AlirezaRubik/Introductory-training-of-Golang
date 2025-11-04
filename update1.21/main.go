package main

import "fmt"

func main() {
	//max number
	fmt.Println(max(2, 12, 534, 11, 634))
	//min number
	fmt.Println(min(2, 12, 534, 11, 634))

	//clear
	var salaries= map[string]float64{}
	salaries["alireza"]=2055
	fmt.Println(salaries["alireza"])
	//clear func
	clear(salaries)
	fmt.Println(salaries)
}
