package main

import (
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Panicf("failed to create request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicf("failed to do request: %v", err)
	}
	log.Println("status code:", resp.StatusCode)
}
