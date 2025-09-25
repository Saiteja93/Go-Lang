package main

import "fmt"

func main (){

	var name string
	var age int
	var percentage float64

	fmt.Print("enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("enter your age: ")
	fmt.Scanln(&age)

    fmt.Print("enter your percantage: ")
	fmt.Scanln(&percentage)

	fmt.Println("\nDetails entered by user")
	fmt.Println("name:",name)
	fmt.Println("age:", age)
	fmt.Println("percentage:", percentage)
}


	


