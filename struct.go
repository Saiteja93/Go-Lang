package main
import "fmt"

type model struct{
	name string
	year float32
	availabilty string

}
func main(){
	var vw model
		vw.name = "Tiguan";
		vw.year = 2024;
		vw.availabilty = "Next month";
		car(vw)
		
}	
func car(v model){
	fmt.Println("model:",v.name)
	fmt.Println("year:",v.year)
	fmt.Println("availability:",v.availabilty)
}

