package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"practice_go/data-storage/mongodb/official-driver/model"
)

func selectData(dbName string, collectionName string, client *mongo.Client) error {
	// Получение коллекции
	collection := client.Database(dbName).Collection(collectionName)

	// Фильтр
	filter := bson.D{{"firstname", "John"}}

	// Запрос
	result := model.Teacher{}

	// Поиск
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("failed to find: %v\n", err)
		return err
	}

	// Вывод
	fmt.Printf("result: %+v\n", result)

	return nil
}
