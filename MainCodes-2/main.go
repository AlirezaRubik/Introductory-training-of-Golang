package main

import "fmt"

func main() {
	//1-
	//making array:
	//*type1:
	// var array1 = [4]int{1, 2, 3, 4}
	// fmt.Printf("%v \n", array1)
	//*type2:
	// array2 := [5]string{"Alireza", "Mohammad", "Naser", "Rostam", "Hossein"}
	// fmt.Printf("%v \n", array2)
	//*type3:
	// array3 := [...]int{22, 366, 4, 787}
	// fmt.Printf("%v \n", array3)
	//*type4:
	// var array4 = [5]int{}
	// array4[0]=2
	// array4[1]=23
	// array4[2]=45
	// array4[3]=14
	// array4[4]=245
	// fmt.Printf("%v \n", array4)
	//*type5:
	// array5 := [5]string{}
	// array5[0]="alireza"
	// array5[1]="Mohsen"
	// array5[2]="Naser"
	// array5[3]="Milad"
	// array5[4]="Mohammad"
	// fmt.Printf("%v \n", array5)
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//2-Slices
	//type1:
	// var Slice1 = []int{1, 2, 3, 4, 5}
	// fmt.Printf("%v \n", Slice1)
	//type2:
	// Slice2 := []string{"Alireza", "Naser", "Milad", "Mohsen"}
	// fmt.Printf("%v \n", Slice2)
	//type3:
	// Slice3 := []int{1, 2, 3, 4, 5}
	// fmt.Printf("%v \n", Slice3)
	//type4:
	// Slice4 := []int{}
	// Slice4 = append(Slice4, 36)
	// Slice4 = append(Slice4, 234)
	// Slice4 = append(Slice4, 3456)
	// Slice4 = append(Slice4, 6578)
	// fmt.Printf("%v \n", Slice4)
	//type 5:
	// var Slice5 =[]string{}
	// Slice5=append(Slice5, "Alireza")
	// Slice5=append(Slice5, "Naser")
	// Slice5=append(Slice5, "Milad")
	// Slice5=append(Slice5, "Mohsen")
	// fmt.Printf("%v \n", Slice5)
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//3-terminate slice with nil
	// Slice1:=[]int{12,24,23,52352}
	// fmt.Printf("Before Terminate %v \n",Slice1)
	// Slice1=nil
	// fmt.Printf("After Terminate %v \n",Slice1)
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//4-terminate slice with equal
	// Slice1:=[]int{12,24,23,52352}
	// fmt.Printf("Before Terminate %v \n",Slice1)
	// Slice1=Slice1[:0]
	// fmt.Printf("After Terminate %v \n",Slice1)
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//5-len and cap in array:
	// array1:=[4]int{12,24,23,}
	// fmt.Printf("The Len Of The Array: %v and The Cap of The Array is : %v  and The Values are: %v \n",len(array1),cap(array1),array1)
	// array1[3]=22
	// fmt.Printf("The Len Of The Array: %v and The Cap of The Array is : %v  and The Values are: %v \n",len(array1),cap(array1),array1)
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//6-len and cap in slice
	// Slice := []int{12, 24, 23}
	// fmt.Printf("The Len Of The Slice: %v and The Cap of The Slice is : %v  and The Values are: %v \n", len(Slice), cap(Slice), Slice)
	// Slice=append(Slice, 22321)
	// Slice=append(Slice, 3456)
	// Slice=append(Slice, 42356)
	// Slice=append(Slice, 5678)
	// fmt.Printf("The Len Of The Slice: %v and The Cap of The Slice is : %v  and The Values are: %v \n", len(Slice), cap(Slice), Slice)
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//7-Copy in Slice:
	// Slice1:=[]int{12,312,41,41,312}
	// Slice2:=[]int{2,3456,252,421}
	// fmt.Printf("[BEFORE]The Values of The Slice 1 are %v \n",Slice1)
	// fmt.Printf("[BEFORE]The Values of The Slice 2 are %v \n",Slice2)
	// copy(Slice1,Slice2)
	// fmt.Printf("[AFTER]The Values of The Slice 1 are %v \n",Slice1)
	// fmt.Printf("[AFTER]The Values of The Slice 2 are %v \n",Slice2)
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//8-Loop In Slice and Array:
	// Slice1 := []int{12, 312, 41, 41, 312}
	// Slice2 := []string{"Alireza", "Naser", "Milad", "Mohsen"}
	// for i := 0; i < len(Slice1); i++ {
	// 	fmt.Printf("The Value is :%v \n", Slice1[i])
	// }
	// for index,item:=range Slice2{
	// 	fmt.Printf("The Value is :%v and The Index is :%v \n", item,index)
	// }
	/////////////////////
	// var array1 = [4]int{1, 2, 3, 4}
	// array2 := [5]string{"Alireza", "Mohammad", "Naser", "Rostam", "Hossein"}
	// for i := 0; i < len(array1); i++ {
	// 	fmt.Printf("The Value is :%v \n", array1[i])
	// }
	// for index, item := range array2 {
	// 	fmt.Printf("The Value is :%v and The Index is :%v \n", item, index)
	// }
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//9-Merching 2 Slices
	// Slice1 := []int{12, 312, 41, 41, 312}
	// Slice2 := []int{464,456,4,5,52,3243}
	// fmt.Printf("[Before append]The Values are: %v and Len of Slice1 is: %d and capacity %d\n", Slice1,len(Slice1),cap(Slice1))
	// fmt.Printf("[Before append]The Values are: %v and Len of Slice2 is: %d and capacity %d\n", Slice2,len(Slice2),cap(Slice2))
	// Slice1=append(Slice1,Slice2...)
	// fmt.Printf("[AFTER append]The Values are: %v and Len of Slice1 is: %d and capacity %d\n", Slice1,len(Slice1),cap(Slice1))
	// fmt.Printf("[AFTER append]The Values are: %v and Len of Slice2 is: %d and capacity %d\n", Slice2,len(Slice2),cap(Slice2))
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//call by value and call by reference in slice
	// Slice:=[]int{60, 70, 80, 90,}
	// fmt.Printf("[Before Go To Call By Value]Slice Values are: %v \n", Slice)
	// CallByValueSlice(Slice)
	// fmt.Printf("[After Go To Call By Value]Slice Values are: %v \n", Slice)
	// fmt.Printf("[Before Go To Call By Reference]Slice Values are: %v \n", Slice)
	// CallByReferenceSlice(&Slice)
	// fmt.Printf("[After Go To Call By Value]Slice Values are: %v \n", Slice)
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//call by value and call by reference in Array
	// Array:=[4]int{60, 70, 80, 90}
	// fmt.Printf("[Before Go To Call By Value]Slice Values are: %v \n", Array)
	// CallByValueArray(Array)
	// fmt.Printf("[After Go To Call By Value]Slice Values are: %v \n", Array)
	// fmt.Printf("[Before Go To Call By Reference]Slice Values are: %v \n", Array)
	// CallByReferenceArray(&Array)
	// fmt.Printf("[After Go To Call By Value]Slice Values are: %v \n", Array)

}
func CallByValueSlice(slice []int) {
	fmt.Printf("[Before Do Call By Value]Slice Values are: %v \n", slice)
	slice = append(slice, 23)
	slice = append(slice, 45)
	slice = append(slice, 24563)
	slice = append(slice, 789)
	slice = append(slice, 245363)
	fmt.Printf("[After Do Call By Value]Slice Values are: %v \n", slice)
}
func CallByReferenceSlice(slice *[]int) {
	fmt.Printf("[Before Do Call By Reference]Slice Values are: %v \n", slice)
	*(slice) = append(*(slice), 23)
	*(slice) = append(*(slice), 45)
	*(slice) = append(*(slice), 24563)
	*(slice) = append(*(slice), 789)
	*(slice) = append(*(slice), 245363)
	fmt.Printf("[After Do Call By Reference]Slice Values are: %v \n", slice)
}

func CallByValueArray(Array [4]int) {
	fmt.Printf("[Before Do Call By Value]Array Values are: %v \n", Array)
	Array[0] = 23
	Array[1] = 42356
	Array[2] = 567890
	Array[3] = 12345
	fmt.Printf("[After Do Call By Value]Array Values are: %v \n", Array)
}
func CallByReferenceArray(Array *[4]int) {
	fmt.Printf("[Before Do Call By Reference]Array Values are: %v \n", Array)
	Array[0] = 23
	Array[1] = 42356
	Array[2] = 567890
	Array[3] = 12345
	fmt.Printf("[After Do Call By Reference]Array Values are: %v \n", Array)
}