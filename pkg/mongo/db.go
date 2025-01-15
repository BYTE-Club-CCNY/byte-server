package mongodb

import (
	"byteserver/pkg/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

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
	db = client.Database("byte-apps")
	fmt.Printf("%s initialized", db.Name())

	return nil
}