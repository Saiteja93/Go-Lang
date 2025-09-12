package main
import "fmt"

func main(){
	j := 21
	fmt.Println("printing even numbers below")

	for i:=0; i<j; i++{
		if i%2 !=0{
		continue
		}

		fmt.Println(i)

	}

	fmt.Printf("\nPrinting mutiplication table of 2\n")

	for x:=1; x<j; x++{

		fmt.Println("2*", x, "=", x*2)
	}
}