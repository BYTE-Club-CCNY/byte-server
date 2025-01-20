package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

func CreateDraft(Cohort_id string) error {
	indexModel := mongo.IndexModel{
        Keys: bson.D{{"docType", 1}}, 
        Options: options.Index().SetPartialFilterExpression(bson.D{
            {"docType", "template"}, 
        }).SetUnique(true),
    }

	cohort := "cohort-" + Cohort_id
	collection := DB.Collection(cohort)

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel) 
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.TODO(), bson.D{
		{"docType", "template"},
	})	

	if err != nil {
		return err
	}

	return nil
}

func ViewDraft(Cohort_id string) (bson.M, error) {
	var result bson.M
	collection := DB.Collection(Cohort_id)
	err := collection.FindOne(context.TODO(), bson.D{{"docType", "template"}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}