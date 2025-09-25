package main
import (
	"fmt"
	"time"
)
func worker(id int, jobs <-chan string, done chan bool) {
    defer func() { done <- true }()  // signal completion
    for job := range jobs {
        fmt.Printf("Worker %d sending %s\n", id, job)
        time.Sleep(time.Second * 5)
    }
}

func main() {
    jobs := make(chan string, 5)
    done := make(chan bool)
    numWorkers := 3

    for i := 1; i <= numWorkers; i++ {
        go worker(i, jobs, done)
    }

    for i := 0; i < 5; i++ {
        jobs <- fmt.Sprintf("%d@gmail.com", i)
    }
    close(jobs)

    for i := 0; i < numWorkers; i++ {
        <-done
    }
    fmt.Println("All emails sent")
}
