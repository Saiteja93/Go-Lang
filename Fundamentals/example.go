package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Task represents a unit of work.
type Task struct {
	ID       int
	Duration time.Duration
}

// Global metrics for observability.
var (
	tasksProcessed uint64
	tasksFailed    uint64
)

// worker is a goroutine that processes tasks from a channel.
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		log.Printf("Worker %d: Starting task %d...", id, task.ID)

		// Simulate a random failure.
		if rand.Intn(100) < 10 { // 10% chance of failure
			log.Printf("Worker %d: Task %d FAILED!", id, task.ID)
			atomic.AddUint64(&tasksFailed, 1)
			continue // Skip to the next task
		}

		// Simulate work by sleeping for the task's duration.
		time.Sleep(task.Duration)

		log.Printf("Worker %d: Finished task %d.", id, task.ID)
		atomic.AddUint64(&tasksProcessed, 1)
	}
}

// main is the entry point of the program.
func main() {
	const (
		numWorkers = 5
		numTasks   = 50
	)

	// A buffered channel to hold tasks.
	tasks := make(chan Task, numTasks)
	var wg sync.WaitGroup

	// Start the worker pool.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Produce tasks and send them to the channel.
	for i := 1; i <= numTasks; i++ {
		duration := time.Duration(rand.Intn(500)+50) * time.Millisecond
		tasks <- Task{ID: i, Duration: duration}
	}
	close(tasks) // Signal that no more tasks will be sent.

	// Wait for all workers to finish.
	wg.Wait()

	// Print final metrics.
	fmt.Println("\n--- Summary ---")
	fmt.Printf("Total Tasks Processed: %d\n", atomic.LoadUint64(&tasksProcessed))
	fmt.Printf("Total Tasks Failed: %d\n", atomic.LoadUint64(&tasksFailed))
	fmt.Println("-----------------")
}