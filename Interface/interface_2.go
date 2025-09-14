package main

import (
	"errors"
	"fmt"
	"log"

	
)

var (
	ErrNotImplemented = errors.New("Not implemented")
	ErrNotFound = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	Unloadcargo() error
}


type Normaltruck struct{
	id string
	cargo int

}

func (t *Normaltruck) LoadCargo() error {
	t.cargo +=1
	return nil
}

func (t *Normaltruck) Unloadcargo() error {
	return nil
}

type Electrictruck struct{
	id string
	cargo int
	battery float64
}

func (e *Electrictruck) LoadCargo() error{
	e.cargo += 1
	e.battery += -1

	return nil
	
}

func (e *Electrictruck) Unloadcargo() error{
	e.cargo = 0
	e.battery += -1

	return nil
}


func Processtruck (truck Truck) error{



	fmt.Printf("\nProceesing truck: %+v \n", truck)

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	err = truck.Unloadcargo()
	if err!= nil{
		return fmt.Errorf("error unloading cargo: %w", err)

	}
	return nil

}

func main(){
	/* trucks:= []Normaltruck{
		{id : "sai-123"},
		{id : "teja-456"},
		{id : "hima-899"},
	} */

	/* etrucks:= []Electrictruck {
		{id : "Ele-hareesh-456"},
		{id : "Ele-jaisurya-899"},
	}  */
	  nt := &Normaltruck{id: "1"}
	  et := &Electrictruck{id : "2"}
	err := Processtruck(nt)
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)

	}
		
	err = Processtruck(et)
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)

	}
	log.Println("normaltruck cargo", nt.cargo)
	log.Println("electric truck battery",et.battery)

		
}



	
		


