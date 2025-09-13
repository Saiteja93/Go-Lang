package main

import (
	"errors"
	"fmt"
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
	// Check empty name
	if p.name == "" {
		return &PodError{name: "Invalid Pod", reason: "pod name must be provided"}
	}

	// Define valid statuses
	validStatus := map[string]bool{
		"Started": true,
		"Failed":  true,
		"Pending": true,
	}

	// Check if status is valid
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

	for _, pod := range pods {
		err := PodsStatus(pod)

		if err != nil {
			var tErr *PodError
			if errors.As(err, &tErr) {
				fmt.Println("⚠️ Custom error:", tErr)
			} else {
				fmt.Println("General error:", err)
			}
			continue
		}

		// Handle pod status
		if pod.status == "Failed" {
			fmt.Printf("Pod %s has failed. Restart required.\n\n", pod.name)
		} else {
			fmt.Printf(" Pod %s is %s\n\n", pod.name, pod.status)
		}
	}
}
