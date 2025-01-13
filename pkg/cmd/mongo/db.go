package mongodb

import (
	"byteserver/pkg/utils"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client mongo.Client

func ConnectDB() {
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
		log.Fatal("Error connecting to MongoDB")
		panic(err)
	}

	// Stops connection to MongoDB once script finishes
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	database := client.Database("byte")
	database.CreateCollection(context.Background(), "spring-2025")
}