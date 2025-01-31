package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckCollectionExists(searchCollection string) (bool, error) {
	filter := bson.D{{"name", searchCollection}}
	collectionNames, err := MongoDB.ListCollectionNames(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return len(collectionNames) > 0, nil
}