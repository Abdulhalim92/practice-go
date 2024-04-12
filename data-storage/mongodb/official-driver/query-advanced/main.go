package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"practice_go/data-storage/mongodb/official-driver/model"
)

func queryBySelector(dbName string, collectionName string, client *mongo.Client) error {
	// Получение коллекции
	collection := client.Database(dbName).Collection(collectionName)

	// Фильтр
	q1 := bson.M{"firstname": bson.M{"$eq": "John"}}
	q2 := bson.M{"lastname": bson.M{"$eq": "Doe"}}
	clauses := bson.A{q1, q2}

	filter := bson.M{"$and": clauses}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Printf("failed to find: %v\n", err)
		return err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var result model.Teacher
		err = cur.Decode(&result)
		if err != nil {
			fmt.Printf("failed to decode: %v\n", err)
			return err
		}
		fmt.Printf("result: %+v\n", result)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("failed to iterate cursor: %v\n", err)
		return err
	}

	return nil

}

// получить всех учителей, у которых имя Джон и фамилия Доу
/*
{
    "$and": [{
            "firstname": {
                "$eq": "John"
            }
        },
        {
            "lastname": {
                "$eq": "Doe"
            }
        }
    ]
}
*/
