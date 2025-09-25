package main
import "fmt"

type Truck struct{
	id string
	department string
}

func ProcessTruck(truck Truck){
	fmt.Printf("\nProcessing truck: %s\n",truck.id)
}

func main(){
	trucks := []Truck{
		{id : "sai-123", department : "vegetable"},
		{id : "Teja-456", department : "clothin"},
		{id : "guvvala-789", department: "cosmetics"},


	}
	for _,truck :=range trucks{
		ProcessTruck(truck)
		fmt.Print("Truck ",truck.id, " arrived. Department: ",truck.department,"\n")
		
	}
}