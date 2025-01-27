package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

func ViewTemplate(collectionName string) (bson.M, error) {
	var result bson.M
	if exists, _ := CheckCollectionExists(collectionName); !exists {
		return nil, fmt.Errorf("Collection %s does not exist", collectionName)
	}
	collection := DB.Collection(collectionName)
	err := collection.FindOne(context.TODO(), bson.D{{"docType", "template"}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateTemplate(collectionName string) error {
	indexModel := mongo.IndexModel{
        Keys: bson.D{{"docType", 1}}, 
        Options: options.Index().SetName("template").SetPartialFilterExpression(bson.D{
            {"docType", "template"}, 
        }).SetUnique(true),
    }

	if exists, _ := CheckCollectionExists(collectionName); !exists {
		return fmt.Errorf("Collection %s does not exist", collectionName)
	}
	collection := DB.Collection(collectionName)

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel) 
	if err != nil {
		return err
	}

	var result bson.M
	err = collection.FindOne(context.TODO(), bson.D{{"docType", "draft"}}).Decode(&result)
	if err != nil {
		return err
	}

	result["docType"] = "template"
	delete(result, "_id")

	if _, err = collection.InsertOne(context.TODO(), result); err != nil {
		return err
	}

	return nil
}