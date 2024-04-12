package main

import (
	"fmt"
	"log"
	"net/http"
)

var requestCount int

func main() {
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
		return
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	requestCount++
	fmt.Fprintf(w, "request count: %d", requestCount)
}
