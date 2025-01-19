package mongodb

import (
	"byteserver/pkg/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

var DB *mongo.Database

func Connect() error {
	err := utils.InitEnv()

	if err != nil {
		panic(".env file missing!")
	}
	
	uri, err := utils.GetEnv("MONGO_URI")
	if err != nil {
		panic("MONGO_URI value missing")
	}

	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	fmt.Println("MongoDB successfully connected...")

	// Initializes the database if it does not exist
	DB = client.Database("byte-apps")

	// Need to create a collection with at least one document, so that the 
	// database persists. Once deployed, we can remove this part of the code
	// to save storage space. 
	collection := DB.Collection("sample")
	validationDoc := bson.D{{Key: "validation", Value: "test"}}
	_, err = collection.InsertOne(context.Background(), validationDoc)
	if err != nil {
		fmt.Printf("Failed to validate collection 'sample': %v", err)
	}

	fmt.Printf("%s initialized", DB.Name())

	return nil
}