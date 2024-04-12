package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func create(client *mongo.Client) error {
	// Получение коллекции
	collection := client.Database("db").Collection("school")

	res, err := collection.InsertOne(context.Background(), bson.M{"firstname": "John", "lastname": "Doe", "create_time": time.Now()})
	if err != nil {
		fmt.Printf("failed to insert: %v\n", err)
		return err
	}

	id := res.InsertedID
	fmt.Printf("inserted id: %v\n", id)

	return nil
}

// bson.M{} --- map[string]interface{}
/*
{
    "_id": {
        "$oid": "5c0919b285d8ae1a8afe6c80"
    },
    "lastname": "Doe",
    "create_time": {
        "$date": "2018-12-06T12:44:34.168Z"
    },
    "firstname": "John"
}
*/
