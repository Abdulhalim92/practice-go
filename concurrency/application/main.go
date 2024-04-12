package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"practice_go/concurrency/application/visit"
	"sync"
)

type DailyStat struct {
	Date   string         `json:"date"`
	ByPage map[string]int `json:"byPage"`
}

type Task struct {
	Date   string
	Visits []visit.Visit
}

func main() {
	dayStats, err := loadJSON()
	if err != nil {
		fmt.Printf("failed to load input data: %v", err)
		return
	}

	// print the results
	for s, visits := range dayStats {
		fmt.Printf("%s: %v\n", s, visits)
	}

	var w8 sync.WaitGroup
	w8.Add(len(dayStats))

	inputCh := make(chan Task, 10)
	outputCh := make(chan DailyStat, len(dayStats))

	// create the workers
	for k := 0; k < len(dayStats); k++ {
		go worker(inputCh, k, outputCh, &w8)
	}

	// send the tasks
	for date, visits := range dayStats {
		inputCh <- Task{
			Date:   date,
			Visits: visits,
		}
	}

	// we say that we will not send any new data on the input channel
	close(inputCh)

	// wait for all tasks to be completed
	w8.Wait()

	// сбор результатов
	// when all treatment is finished we close the output channel
	close(outputCh)

	// collect the result
	done := make([]DailyStat, 0, len(dayStats))

	for out := range outputCh {
		done = append(done, out)
	}

	res, err := json.MarshalIndent(done, "", "  ")
	if err != nil {
		log.Printf("failed to marshal the data: %v", err)
		return
	}

	err = os.WriteFile("./concurrency/application/data/results.json", res, 0644)
	if err != nil {
		log.Printf("failed to write results: %v", err)
		return
	}

	log.Println("done")
}

func loadJSON() (map[string][]visit.Visit, error) {
	data, err := os.ReadFile("./concurrency/application/data/input.json")
	if err != nil {
		log.Printf("failed to load input data: %v", err)
		return nil, err
	}

	dayStats := make(map[string][]visit.Visit)

	err = json.Unmarshal(data, &dayStats)
	if err != nil {
		log.Printf("failed to unmarshal the data from file: %v", err)
		return nil, err
	}

	return dayStats, nil
}

func worker(in chan Task, workerID int, out chan DailyStat, w8 *sync.WaitGroup) {
	for received := range in {
		m := make(map[string]int)
		for _, v := range received.Visits {
			m[v.Page]++
		}
		out <- DailyStat{
			Date:   received.Date,
			ByPage: m,
		}
		fmt.Printf("[worker %d] finished task for %s\n", workerID, received.Date)
	}
	// when the channel is closed the for loop is exited
	log.Println("worker quit")
	w8.Done()
}
