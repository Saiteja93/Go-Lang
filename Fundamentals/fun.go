package main

import "fmt"

func math(x int, y int) (add int, sub int){

	add = x+y
	sub = x-y
	return

	
	

	
}

func main(){

	a:= 10
	b := 20

	add, sub := math(a,b)
	fmt.Println(add, sub)

	
	
	
}
