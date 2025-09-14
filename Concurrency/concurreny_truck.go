package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	
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


func Processtruck (ctx context.Context, truck Truck) error{



	fmt.Printf("\nStarted proceesing truck: %+v \n", truck)

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	err = truck.Unloadcargo()
	if err!= nil{
		return fmt.Errorf("error unloading cargo: %w", err)

	}

	fmt.Printf("Finished processing truck: %+v \n", truck)
	return nil

}
	
func ProcessFleet(ctx context.Context, trucks []Truck) error{
	var wg sync.WaitGroup
	errorsChan := make(chan error, len(trucks))
	defer close(errorsChan)
	

	for _,t := range trucks{
		wg.Add(1)
	go func(t Truck){
		
		
		if err := Processtruck(ctx, t); err != nil{
			log.Println(err)
			errorsChan <- err
		}
		defer wg.Done()	
		
		}(t)
		
	}
	wg.Wait()
	
	select{
	case err := <- errorsChan:
	return err

	default:
		return nil

	}
	
	}




func main(){
	ctx := context.Background()

	fleet := []Truck {
		&Normaltruck{id : "NT1", cargo : 0},
		&Electrictruck{id : "ET1", cargo: 0, battery : 100},
		&Normaltruck{id : "NT2", cargo : 0},
		&Electrictruck{id : "ET2", cargo: 0, battery : 100},
		&Normaltruck{id : "NT3", cargo : 0},
		&Electrictruck{id : "ET3", cargo: 0, battery : 100},
		

	}
	if err := ProcessFleet(ctx,fleet); err != nil{
		fmt.Printf("Error processing fleet: %v", err)
		return

	}

	fmt.Println("All trucks are processed")
	
	

	}

	

	
