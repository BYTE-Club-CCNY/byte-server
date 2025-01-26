package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func InsertJSON(params bson.M) error {
	cohort_id := params["cohort_id"].(string) 
	collection := DB.Collection("cohort-" + cohort_id)
	delete(params, "cohort_id")
	if _, err := collection.InsertOne(context.TODO(), params); err != nil {
		return err
	}
	return nil
}