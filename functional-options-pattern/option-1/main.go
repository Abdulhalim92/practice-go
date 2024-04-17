package main

import (
	"log"
	"practice_go/functional-options-pattern/option-1/server"
	"time"
)

func main() {
	srv := server.NewWithTimeoutAndMaxConn("localhost", 8080, 5*time.Second, 100)
	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
