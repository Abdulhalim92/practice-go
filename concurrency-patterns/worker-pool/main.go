package main

import (
	"fmt"
	"sync"
)

// Шаблон пула рабочих предполагает создание группы рабочих программ
// для одновременной обработки задач, ограничивая количество одновременных
// операций. Этот шаблон полезен, когда вам нужно выполнить большое
// количество задач.
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	// Enqueue jobs
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
