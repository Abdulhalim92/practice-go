package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu           sync.Mutex
	requestCount int
)

func main() {
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	requestCount++
	mu.Unlock()
	fmt.Fprintf(w, "OK - count : %d \n", requestCount)
}
