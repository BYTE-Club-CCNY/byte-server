package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNewCollection(season *string) error {
	err := db.CreateCollection(context.TODO(), *season) 
	if err != nil {
		return err
	}
	return nil
}

func GetAllApps(season string) []bson.D {
	collection := db.Collection(season)
	var result []bson.D
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.Decode(&result); err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())
	return result
}