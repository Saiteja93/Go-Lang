package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup){

	defer w.Done()
	fmt.Println("Task id:", id)
}
func main(){
	var wg sync.WaitGroup
	fmt.Println("All tasks are started")
	for i := 0; i<=10; i++ {
		wg.Add(1)
		go task(i, &wg)
		

	}
	wg.Wait()

	fmt.Println("All tasks are completed")
	
}