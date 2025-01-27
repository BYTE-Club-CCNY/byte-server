package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertJSON(collection *mongo.Collection, params bson.M) error {
	if _, err := collection.InsertOne(context.TODO(), params); err != nil {
		return err
	}
	return nil
}

func UpdateJSON(collection *mongo.Collection, params bson.M, filter bson.D) error {
	update := bson.D{{"$set", params}}
	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}
	return nil
}

func UpdateOrInsertJSON (params bson.M, isSubmit bool) error {
	cohort_id := params["cohort_id"].(string)
	collectionName := "cohort-" + cohort_id
	if exists, _ := CheckCollectionExists(collectionName); !exists {
		return fmt.Errorf("Collection %s does not exist", collectionName)
	}
	collection := MongoDB.Collection(collectionName)
	delete(params, "cohort_id")

	if isSubmit {
		params["submitted"] = true
	}

	filter := bson.D{{"user_id", params["user_id"]}}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return UpdateJSON(collection, params, filter)
	} else {
		return InsertJSON(collection, params)
	}
}

func GetApps(collectionName string, pages, limit int) ([]bson.M, error) {
	if exists, _ := CheckCollectionExists(collectionName); !exists {
		return nil, fmt.Errorf("Collection %s does not exist", collectionName)
	}
	collection := MongoDB.Collection(collectionName)
	var result []bson.M
	
	filter := bson.D{{"submitted", true}}
	skip := int64((pages - 1) * limit)
	options := options.Find().SetLimit(int64(limit)).SetSkip(skip)
	docs, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}

	if err = docs.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}