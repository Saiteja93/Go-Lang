package main

import (
	"fmt"
	"errors"
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
	cargo string

}

func (t *Normaltruck) LoadCargo() error {
	return nil
}

func (t *Normaltruck) Unloadcargo() error {
	return nil
}

type Electrictruck struct{
	id string
	caro int
	battery float64
}

func Processtruck (truck Truck) error{

	err := truck.LoadCargo()
	if err != nil{
		return fmt.Errorf("error loading cargo: %w",err)
	}

	err = truck.Unloadcargo()
	if err!= nil{
		return fmt.Errorf("error unloading cargo: %w", err)

	}
	return nil

}

func main(){
	trucks:= []Normaltruck{
		{id : "sai-123"},
		{id : "teja-456"},
		{id : "hima-899"},
	}

	etrucks:= []Electrictruck {
		{id : "Ele-hareesh-456"},
		{id : "Ele-jaisurya-899"},
	}

	for _, truck := trucks range {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		err := Processtruck(truck)
		if err != nil {
			log.Fatalf("Error processing truck: %s", errerr)

		}
		fmt.Printf("Truck %s departed./n,", truck.id)
		
	}

}


