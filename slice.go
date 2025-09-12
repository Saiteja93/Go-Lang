package main

import "fmt"

func main(){
	var sai = make([]float64, 5,10)
	fmt.Println("sai:", sai, "length:", len(sai), "capacity:", cap(sai))
	sai = append(sai, 1,2,3,4,5,6,7,8,9,10)
	fmt.Println("sai:", sai, "length:", len(sai), "capacity:", cap(sai))
}