package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
)

func CreateNewCohort(ctx context.Context, cohort string) error {
	if cohort == "" {
		return errors.New("collection name cannot be empty")
	}

	if err := MongoDB.CreateCollection(ctx, cohort); err != nil {
		return fmt.Errorf("failed to create collection '%s': %w", cohort, err)
	}

	return nil
}

func GetAllData(ctx context.Context, collectionName string) ([]bson.M, error) {
	if exists, _ := CheckCollectionExists(collectionName); !exists {
		return nil, fmt.Errorf("Collection %s does not exist", collectionName)
	}
	collection := MongoDB.Collection(collectionName)
	var result []bson.M
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}