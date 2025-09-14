package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type TruckManager struct {
	mu     sync.Mutex
	trucks map[string]*Truck
}

func NewTruckManager() *TruckManager {
	return &TruckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *TruckManager) AddTruck(id string, cargo int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.trucks[id] = &Truck{
		ID:    id,
		Cargo: cargo,
	}
	return nil
}

func (tm *TruckManager) GetTruck(id string) (*Truck, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	truck, ok := tm.trucks[id]
	if !ok {
		return nil, ErrTruckNotFound
	}
	return truck, nil
}

func (tm *TruckManager) RemoveTruck(id string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	delete(tm.trucks, id)
	return nil
}

func (tm *TruckManager) UpdateTruckCargo(id string, cargo int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	truck, ok := tm.trucks[id]
	if !ok {
		return ErrTruckNotFound
	}

	truck.Cargo = cargo
	return nil
}

// Example usage
func main() {
	tm := NewTruckManager()

	// Add trucks concurrently
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			tm.AddTruck(fmt.Sprintf("T%d", id), id*10)
		}(i)
	}
	wg.Wait()

	// Read trucks concurrently
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			truck, err := tm.GetTruck(fmt.Sprintf("T%d", id))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Truck: %+v\n", truck)
		}(i)
	}
	wg.Wait()
}
