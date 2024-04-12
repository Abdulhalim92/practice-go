package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"practice_go/data-storage/mongodb/official-driver/model"
)

func readById(dbName string, collectionName string, client *mongo.Client) error {
	// Получение коллекции
	collection := client.Database(dbName).Collection(collectionName)

	// Идентификатор
	objectID, err := primitive.ObjectIDFromHex("5c091d3b734016209db89f76")
	if err != nil {
		fmt.Printf("failed to convert: %v\n", err)
		return err
	}

	// Фильтр
	filter := bson.M{"_id": objectID}

	// Запрос
	result := model.Teacher{}

	// Поиск
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("failed to find: %v\n", err)
		return err
	}

	// Вывод
	fmt.Printf("result: %+v\n", result)

	return nil
}
