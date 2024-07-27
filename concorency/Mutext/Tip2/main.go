package main

import (
	"fmt"
	"sync"
)

var Salary int = 11000000

func main() {
	wg := sync.WaitGroup{}
	mutext := sync.Mutex{}
	for i := 0; i < 10; i++ {
       wg.Add(1)
        go func() {
          mutext.Lock()
		  GiveSalarys(&Salary)
		  mutext.Unlock()
		  wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("After Pay For 10 Person We Have:",Salary)
}
func GiveSalarys(Salary *int) {
	PerPerson := 1000000
	*Salary -= PerPerson
	fmt.Println("After Pay We Have:", *Salary)
}
