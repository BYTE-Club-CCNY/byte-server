package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
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
	collection := DB.Collection("cohort-" + cohort_id)
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