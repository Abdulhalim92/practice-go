package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(2 * time.Second) // simulate work
		fmt.Println("worker", id, "finished job", j)
		results <- j + " job done"
	}
}

func main() {
	numJobs := 5

	jobList := []string{"job1", "job2", "job3", "job4", "job5"}

	// Create channels
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)

	// Create workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for _, job := range jobList {
		jobs <- job
	}
	close(jobs)

	// Get results
	for i := 1; i <= 3; i++ {
		fmt.Printf("Worker %d: %s\n", i, <-results)
	}
}
