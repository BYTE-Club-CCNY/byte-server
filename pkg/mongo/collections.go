package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"fmt"
)

func CreateNewCohort(ctx context.Context, cohort string) error {
	if cohort == "" {
		return errors.New("collection name cannot be empty")
	}

	if err := DB.CreateCollection(ctx, cohort); err != nil {
		log.Printf("failed to create collection '%s': %s", cohort, err.Error())
		return fmt.Errorf("failed to create collection '%s': %w", cohort, err)
	}

	return nil
}

func GetAllData(ctx context.Context, collectionName string) ([]bson.M, error) {
	if exists, _ := CheckCollectionExists(collectionName); !exists {
		return nil, fmt.Errorf("Collection %s does not exist", collectionName)
	}
	collection := DB.Collection(collectionName)
	var result []bson.M
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Failed to retrieve applications from collection '%s': %v", collectionName, err)
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &result); err != nil {
		log.Printf("Failed to decode documents from collection '%s': %v", collectionName, err)
		return nil, err
	}

	return result, nil
}