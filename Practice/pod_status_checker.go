package main

import (
	"fmt"
	"time"
)

type Pod struct{
	name string
	status string
}
func Getpodstatus(pods []Pod){
	for _, pod := range pods{
	fmt.Printf("Processing pod: %s |  status: %s\n", pod.name, pod.status)
	}
}
func main() {
	pods := []Pod{
		{"Nginx", "Running"},
		{"Api", "Failed"},
		{"DB-3", "CrashLoopBackOff"},
		{"Payment-4", "ImagePullBackOff"},

	}
	println("Fetching pod status")

	for i:= 0; i<4; i++{
		Getpodstatus(pods)
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Fetching status of pods completed")


}





