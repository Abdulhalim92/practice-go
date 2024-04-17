package main

import (
	"log"
	"practice_go/functional-options-pattern/option-3/server"
	"time"
)

func main() {
	srv := server.NewServer(
		server.WithHost("localhost"),
		server.WithPort(8080),
		server.WithMaxConn(100),
		server.WithTimeout(5*time.Second),
	) // Use functional options to configure the server.
	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
