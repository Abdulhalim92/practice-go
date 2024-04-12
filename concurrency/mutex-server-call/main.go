package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go caller(&wg)
	}
	wg.Wait()
}

func caller(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		res, err := http.Get("http://localhost:8080/status")
		if err != nil {
			log.Fatalf("failed to get: %v", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("failed to read: %v", err)
			return
		}

		log.Println(string(data))
	}

	wg.Done()
}
