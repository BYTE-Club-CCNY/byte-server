package mongodb

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

var client mongo.Client

func ConnectDB() {
	// Load environment variables so we can retrieve them
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("MONGO_URI") // retrieve mongodb URI

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