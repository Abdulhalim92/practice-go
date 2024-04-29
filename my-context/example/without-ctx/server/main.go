package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request path:", r.URL.Path)
		time.Sleep(3 * time.Second)
		//time.Sleep(60 * time.Second)

		_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
		if err != nil {
			log.Printf("failed to write response: %v", err)
		}
		log.Println("response sent")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicf("failed to listen and serve: %v", err)
	}
}
