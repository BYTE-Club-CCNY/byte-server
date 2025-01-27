package mongodb

import (
	"byteserver/pkg/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func Connect() error {
	utils.InitEnv()

	uri, err := utils.GetEnv("MONGO_URI")
	if err != nil {
		return err
	}

	fmt.Println(uri)

	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	fmt.Println("MongoDB successfully connected...")

	// Initializes the database if it does not exist
	MongoDB = client.Database("byte-apps")

	// Need to create a collection with at least one document, so that the 
	// database persists. Once deployed, we can remove this part of the code
	// to save storage space. 
	collection := MongoDB.Collection("sample")
	validationDoc := bson.D{{Key: "validation", Value: "test"}}
	_, err = collection.InsertOne(context.Background(), validationDoc)
	if err != nil {
		return err
	}

	fmt.Printf("%s initialized", MongoDB.Name())

	return nil
}