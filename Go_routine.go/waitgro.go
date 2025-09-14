package main

import (
	"errors"
	"fmt"
	"sync"
)

// Custom Error Type
type PodError struct {
	name   string
	reason string
}

func (e *PodError) Error() string {
	return fmt.Sprintf("Pod Error - %s: %s", e.name, e.reason)
}

// Pod struct
type Pod struct {
	name   string
	status string
}

// Validate Pod status
func PodsStatus(p Pod) error {
	if p.name == "" {
		return &PodError{name: "Invalid Pod", reason: "pod name must be provided"}
	}

	validStatus := map[string]bool{
		"Started": true,
		"Failed":  true,
		"Pending": true,
		"Running" : true,
	}

	if !validStatus[p.status] {
		return &PodError{name: p.name, reason: "Invalid pod status: " + p.status}
	}

	fmt.Printf("Processing pod: %s\n", p.name)
	return nil
}

func main() {
	pods := []Pod{
		{name: "pod-1", status: "Running"},
		{name: "pod-2", status: "Pending"},
		{name: "pod-3", status: "Failed"},
		{name: "pod-4", status: "not"},
		{name: "", status: "Pending"},
	}

	var wg sync.WaitGroup

	for _, pod := range pods {
		wg.Add(1)

		// Launch goroutine per pod
		go func(p Pod) {
			defer wg.Done()

			err := PodsStatus(p)

			if err != nil {
				var tErr *PodError
				if errors.As(err, &tErr) {
					fmt.Println("⚠️ Custom error:", tErr)
				} else {
					fmt.Println("General error:", err)
				}
				return
			}

			if p.status == "Failed" {
				fmt.Printf("🚨 Pod %s has failed. Restart required.\n\n", p.name)
			} else {
				fmt.Printf("✅ Pod %s is %s\n\n", p.name, p.status)
			}
		}(pod) // pass pod as argument (avoid closure bug!)
	}

	wg.Wait()
	fmt.Println("✅ All pod checks completed (concurrently).")
}
