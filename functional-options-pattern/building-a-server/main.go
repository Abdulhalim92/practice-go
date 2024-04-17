package main

import (
	"log"
	"practice_go/functional-options-pattern/building-a-server/server"
)

func main() {
	srv := server.NewServer("localhost", 8080)
	if err := srv.Start; err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
