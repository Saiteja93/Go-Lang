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
	t := Normaltruck{cargo: "0"}
	log.Printf("Address of t: %p",&t)
	Filltruckcargo(t)

}
func Filltruckcargo(t Normaltruck){
	t.cargo = "100"
	log.Printf("Address of fill: %p", &t)
}
	

