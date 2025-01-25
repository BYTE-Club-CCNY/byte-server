package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"byteserver/pkg/utils"
	"context"
)

func EditDraft(data utils.EditDraft) error {
	cohort := "cohort-" + data.Cohort_id
	collection := DB.Collection(cohort)

	filter := bson.D{{"docType", "draft"}}
	addData := bson.D{{"deadline", data.Deadline}, {"questions", data.Questions}}
	update := bson.D{{"$set", addData}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func CreateDraft(Cohort_id string) error {
	indexModel := mongo.IndexModel{
        Keys: bson.D{{"docType", 1}}, 
        Options: options.Index().SetPartialFilterExpression(bson.D{
            {"docType", "draft"}, 
        }).SetUnique(true),
    }

	cohort := "cohort-" + Cohort_id
	collection := DB.Collection(cohort)

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

func ViewDraft(Cohort_id string) (bson.M, error) {
	var result bson.M
	collection := DB.Collection(Cohort_id)
	err := collection.FindOne(context.TODO(), bson.D{{"docType", "draft"}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}