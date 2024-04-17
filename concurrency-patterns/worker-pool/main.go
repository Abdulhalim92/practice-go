package main

import (
	"context"
	"fmt"
	"sync"
)

/*Implement worker pool example-3 and some job*/

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
		fmt.Printf("worker: %d\n", i)
		wg.Add(1)
		//fan-out worker goroutines
		//reading from job channel and
		//pushing calculates into results channel
		go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.Done)
	close(wp.results)
}

//Jobs Batch

type Job struct {
	Descriptor JobDescriptor
	ExecFn     ExecutionFn
	Args       interface{}
}

type Result struct {
	Value int //int for example useCase
	Err   error
}

type JobDescriptor string

type ExecutionFn func(context.Context, interface{}) (int, error)

func (j Job) execute(ctx context.Context) Result {
	val, err := j.ExecFn(ctx, j.Args)
	if err != nil {
		return Result{
			Err:   err,
			Value: 0,
		}
	}

	fmt.Printf("Result: %d\n", val)

	return Result{
		Value: val,
		Err:   nil,
	}
}

// GenerateFrom ------------------------------
// Job Producer
func (wp WorkerPool) GenerateFrom(jobsBulk []Job) {
	for i := range jobsBulk {
		wp.jobs <- jobsBulk[i]
	}

	close(wp.jobs)
}

// Worker(consumer)
func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("worker: %v\n", job)
			//fan-in job execution multiplexing results into the results channel
			results <- job.execute(ctx)
		case <-ctx.Done():
			fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}

const wgCount = 4

func main() {

	wp := New(wgCount)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//empty slice -- non-nil
	go wp.GenerateFrom([]Job{})

	go wp.Run(ctx)

	for {
		select {
		case r, ok := <-wp.results:
			if !ok {
				continue
			}
			fmt.Printf("result: %v\n", r)
		case <-wp.Done:
			return
		default:
		}
	}
}
