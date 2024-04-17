package main

import (
	"log"
	"practice_go/functional-options-pattern/option-2/server"
	"time"
)

func main() {
	srv := server.NewServer(server.Config{
		Host:    "localhost",
		Port:    8080,
		Timeout: 5 * time.Second,
		MaxConn: 100,
	})
	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
