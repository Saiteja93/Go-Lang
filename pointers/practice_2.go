package main

import "fmt"

func sum(arr *[]int) int{
	total := 0
	for _,v := range *arr{
		total+=v
	}
	return total
}

func main(){
	numbers :=[]int {1,2,3,4,5}
	fmt.Println("sum:",sum(&numbers))
}