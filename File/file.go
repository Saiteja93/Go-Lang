package main

import (
	"fmt"
	"os"
)

func main(){
	f,err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil{
		panic(err)
	}

	fmt.Println("File name: ",fileInfo.Name())
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("Is directory:", fileInfo.IsDir())

	
}
