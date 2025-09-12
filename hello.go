package main

import (
	"fmt"
	
)
func main(){

	const (
		lastname = "guvvala"
		firstname = "saiteja"
	)
	var roll = 25
	var name string 
	var perc float32 
	name = "saiteja"
	perc = 93.5
	fmt.Println("hi hello my percentage:",perc)
	fmt.Println(" my name is", name)
	fmt.Println("my roll number:",roll)
	fmt.Print(perc, "\n",roll)
	fmt.Printf("my lastname: %s and my firstname: %s",lastname,firstname)
}


