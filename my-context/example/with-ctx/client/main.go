package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithTimeout(rootCtx, 10*time.Millisecond)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Panicf("failed to create request: %v", err)
		return
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to do request: %v", err)
		return
	}
	log.Println("status code:", resp.StatusCode)
}
