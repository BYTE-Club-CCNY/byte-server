package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"byteserver/pkg/utils"
	"context"
	"fmt"
)

func EditDraft(data utils.EditDraft) error {
	collectionName := "cohort-" + data.Cohort_id
	if exists, _ := CheckCollectionExists(collectionName); !exists {
		return fmt.Errorf("Collection %s does not exist", collectionName)
	}
	collection := DB.Collection(collectionName)

	filter := bson.D{{"docType", "draft"}}
	addData := bson.D{{"deadline", data.Deadline}, {"questions", data.Questions}}
	update := bson.D{{"$set", addData}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func CreateDraft(collectionName string) error {
	indexModel := mongo.IndexModel{
        Keys: bson.D{{"docType", 1}}, 
        Options: options.Index().SetName("draft").SetPartialFilterExpression(bson.D{
            {"docType", "draft"}, 
        }).SetUnique(true),
    }

	collection := DB.Collection(collectionName)

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel) 
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.TODO(), bson.D{
		{"docType", "draft"},
	})	

	if err != nil {
		return err
	}

	return nil
}

func ViewDraft(collectionName string) (bson.M, error) {
	var result bson.M
	
	if exists, _ := CheckCollectionExists(collectionName); !exists {
		return nil, fmt.Errorf("Collection %s does not exist", collectionName)
	}
	
	collection := DB.Collection(collectionName)
	err := collection.FindOne(context.TODO(), bson.D{{"docType", "draft"}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}