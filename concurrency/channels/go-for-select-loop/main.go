package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	// go forSelectLoop()
	go doWork(done)
	time.Sleep(3 * time.Second)
	close(done)
}

func forSelectLoop() {
	for {
		select {
		default:
			fmt.Println("DOING WORK")
		}
	}

}

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}
