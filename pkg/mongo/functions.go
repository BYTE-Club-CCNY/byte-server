package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func CreateNewCollection(season *string) error {
	err := db.CreateCollection(context.TODO(), *season) 
	if err != nil {
		return err
	}
	return nil
}

func GetAllApps(season string) []bson.M {
	collection := db.Collection(season)
	var result []bson.M
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Failed to retrieve applications from collection")
		panic(err)
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Println("Failed to decode all documents")
		panic(err)
	}


	return result
}