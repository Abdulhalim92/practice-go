package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://username:password@localhost:27094/db"))
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("failed to ping: %v\n", err)
		return
	}
}
