package main
import "fmt"

func main() {

	var age int
	var occupation string

	fmt.Print("Enter your age:")
	fmt.Scanln(&age)
	fmt.Print("Enter your occupatio:")
	
	fmt.Scanln(&occupation)

	if age <=18 {
		fmt.Println("\nYour age is under age to receive vote. please come when you are 18 years old")
		

	}else if age>=60{
		fmt.Println("\nYour age is not valid for applying for new vote")

	}else{
		if occupation == "Government"{
			fmt.Println("\nVote will be issued in thasildar office. Please visit thasildhar office.")

		}else{
			fmt.Println("\nVote is issued")
			
		}
		
		

		
	}
	fmt.Println("Your age is:", age)
	fmt.Println("Your occupation is: ", occupation)
	fmt.Println("\nThank you for visiting vote portal!")
}