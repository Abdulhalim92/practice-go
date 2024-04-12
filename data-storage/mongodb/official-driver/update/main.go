package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"practice_go/data-storage/mongodb/official-driver/model"
)

func updateData(dbName string, collectionName string, client *mongo.Client) error {
	// Получение коллекции
	collection := client.Database(dbName).Collection(collectionName)

	// Фильтр
	filter := bson.M{"firstname": "John"}
	update := bson.M{"$set": bson.M{"lastname": "CoffeeBean"}}

	// Обновление
	res := collection.FindOneAndUpdate(context.Background(), filter, update)

	// Вывод
	result := model.Teacher{}
	err := res.Decode(&result)
	if err != nil {
		fmt.Printf("failed to decode: %v\n", err)
		return err
	}

	fmt.Printf("result: %+v\n", result)

	return nil
}

// Output:
// result: Teacher{Id:ObjectID("5c091d3b734016209db89f76") Firstname:John Lastname:CoffeeBean CreateTime:2018-12-06 23:59:39.338 +0100 CET}

/*
Фильтр:
	{
	    "firstname": {
	        "$eq": "John"
	}
Или:
	{"firstname":"John"}
Применяемая модификация:
	{
	    "$set": {
	        "lastname": "CoffeeBean"
	}
*/
