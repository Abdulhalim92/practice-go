package main

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case jobs, ok := <-jobs:
			if !ok {
				fmt.Printf("worker: jobs channel closed\n")
				return
			}
			// fan-in job execution multiplexing returns into the results channel
			results <- jobs.execute(ctx)
		case <-ctx.Done():
			fmt.Printf("worker: context cancelled. Error details: %v\n", ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
		}
	}
}

type WorkerPool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
	Done         chan struct{}
}

func New(count int) WorkerPool {
	return WorkerPool{
		workersCount: count,
		jobs:         make(chan Job, count),
		results:      make(chan Result, count),
		Done:         make(chan struct{}),
	}
}

func (wp WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		// fan-out worker goroutines
		// reading from jobs channel and
		// pushing calls into results channel
		go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.Done)
	close(wp.results)
}

func (wp WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp WorkerPool) GenerateFrom(jobBulk []Job) {
	for i := range jobBulk {
		wp.jobs <- jobBulk[i]
	}
	close(wp.jobs)
}
