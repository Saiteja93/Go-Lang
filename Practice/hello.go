package main

import (
	"fmt"
	"time"
)


func emailsender(emailchan chan string, done chan bool){

	defer func(){ done <-true}()

	

	for email := range emailchan{
		fmt.Println("Sending email to ",email)
		time.Sleep(time.Second *2)
	}


}
func main(){
	emailchan:= make(chan string, 5)
	done := make(chan bool)

	go emailsender(emailchan,done)

	for i:=0; i<5; i++{

		emailchan <- fmt.Sprintf("%d@gmail.com",i)
	}
	close(emailchan)
	fmt.Println("Done sending emails")
	<- done

	//emailchan <- "1@gmail.com"
	//emailchan <- "2@fmail.com"



	//fmt.Println(<-emailchan)
	//fmt.Println(<-emailchan)
}