package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// добавление заголовка
	req.Header.Add("Accept", `application/json`)

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
	fmt.Printf("Body: %s", body)
}
