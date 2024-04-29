package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	// GET запрос
	getRequest(&client)

	// POST запрос
	postRequest(&client)
}

func getRequest(client *http.Client) {
	resp, err := client.Get("https://www.goolge.com")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error during reading response body: %s", err)
		return
	}

	fmt.Printf("Body: %s\n", body)
}

func postRequest(client *http.Client) {
	myJson := bytes.NewBuffer([]byte(`"name":"Abduhalim"`))

	resp, err := client.Post("https://www.goolge.com", "application/json", myJson)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(resp.Header)
}
