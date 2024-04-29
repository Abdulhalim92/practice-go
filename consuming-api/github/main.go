package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// https://developer.github.com/v3 --- Документация API Github

func main() {
	client := http.Client{
		Timeout: time.Duration(5) * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com", nil)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("TOKEN")))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf("Body: %s\n", body)
	fmt.Printf("Response status: %s\n", resp.Status)
}

// GITHUB_TOKEN=aabbcc go run main.go
