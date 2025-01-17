package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"fmt"
)

func CreateNewCollection(ctx context.Context, collection string) error {
	if collection == "" {
		return errors.New("collection name cannot be empty")
	}

	if err := db.CreateCollection(ctx, collection); err != nil {
		log.Printf("failed to create collection '%s': %w", collection, err)
		return fmt.Errorf("failed to create collection '%s': %w", collection, err)
	}
	
	return nil
}

func GetAllData(ctx context.Context, season string) ([]bson.M, error) {
	collection := db.Collection(season)
	var result []bson.M
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Failed to retrieve applications from collection '%s': %v", season, err)
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &result); err != nil {
		log.Printf("Failed to decode documents from collection '%s': %v", season, err)
		return nil, err
	}

	return result, nil
}