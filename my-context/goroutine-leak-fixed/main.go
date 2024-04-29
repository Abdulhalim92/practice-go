package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

// Утечка горутины -- исправление
func main() {
	log.Println("begin program")
	go launch()
	time.Sleep(time.Millisecond)
	log.Printf("Goroutine count: %d\n", runtime.NumGoroutine())
	for {

	}
}

func doSth2(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("second goroutine return")
		return
	}
}

func doSth(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("first goroutine return")
		return
	}
}

func launch() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // отменяет операцию и удаляет контекст с родительского контекста
	log.Println("launch first goroutine")
	go doSth(ctx)
	log.Println("launch second goroutine")
	go doSth2(ctx)
}
