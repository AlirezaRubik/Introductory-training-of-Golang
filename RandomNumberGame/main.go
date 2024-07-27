package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//set random number
	rnd := rand.Intn(1000)
	//set times try 
	count := 0
	//infinite loop until done!
	for {
		fmt.Println("Please Enter The Random Number:")
		var num int
			//get random number from user
		fmt.Scanln(&num)
	    //check if user guess correct number
        count++
		if num == rnd {
			fmt.Println("Congratulations! You've guessed the correct number.")
			fmt.Printf("After %d Times Try \n",count)
			break
		}
		//check if user guess higher or lower than random number
		if num > rnd {
			fmt.Println("Its High Number Enter Little")
			count++
			continue
		}
		//if user guess lower than random number
		if num < rnd {
			fmt.Println("Its Low Number Enter Higher")
			count++
			continue
		}
	}
	//set random number to 0 again
	count=0
}
