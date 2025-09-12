package main

import "fmt"

func main() {
	var day int
	fmt.Print("Enter the day: ")
	fmt.Scanln(&day)
	switch day{
	case 1,3,5:
		fmt.Println("it is odd weekday")
	case 2,4:
		fmt.Println("it is even weekday")
	case 6,7:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it is not a weekday, give valid input")


	}
}