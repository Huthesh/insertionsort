package algo

import (

)


func InsertionSort(array []int){
	length:=len(array)
	for i:=0;i<length;i++{
		key:=array[i]
		j:=i
		for j>0 && array[j-1]>key {
				array[j]=array[j-1]
				j--
		}
		
		array[j]=key
		
	}
}
