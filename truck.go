package main

import "fmt"
import "errors"

type Truck struct{
	id string
	department string
}

type TruckError struct {
	TruckID string
	Reason string
}

func (e *TruckError) Error()string {
	return fmt.Sprintf("truck error: %s",e.TruckID, e.Reason)
}


func ProcessDetails(truck Truck) error {

	if truck.id == ""{
		return &TruckError{TruckID : "UNKNOWN" , Reason : "missing truck id"}
	}

	validDepartents := map[string] bool {
		"Vegetables" : true,
		"Clothing" : true,
		"cosmetics" : true,
	}

	if !validDepartents[truck.department]{
		return &TruckError{TruckID : truck.id, Reason : "invalid department" + truck.department}
	}



	fmt.Print("\nProcess trucks: ",truck.id,"\n")
	return nil
}

func main(){

	trucks := []Truck {
		{id : "sai-123", department : "Vegetables"},
		{id : "Teja-456", department:   "Clothing"},
		{id : "hima-789", department : "cosmetics"},
		{id : "Moni-1224", department:   "Furniture"},
		{id : "", department : "cosmetics"},
	}
	for _,truck := range trucks{
		err := ProcessDetails(truck)
		if err != nil {
			var tErr *TruckError
			if errors.As(err, &tErr) {
				fmt.Println("Custom error: ", tErr)
			}else {
				fmt.Println("General error: ", err)
				}
				continue
			}

			
		
		fmt.Println("Truck ", truck.id, " arrived. Department is  ", truck.department, "\n")
	}
}
