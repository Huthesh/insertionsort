package main

import (
	"insertionsort/algo"
	"fmt"
)


func main(){
	array:=[]int{100, 13, 50, 17, 111, 103}
	PrintArray(array)
	algo.InsertionSort(array)
	PrintArray(array)
}

func PrintArray(array []int) {
	fmt.Println()
	for num:= range array {
		fmt.Print(array[num])
		fmt.Print(" ")
	} 
}