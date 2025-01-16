package mongodb

import (
	"testing"
	"log"
)

func TestInitDB(t *testing.T) {
	err := Connect()

	if err != nil {
		log.Fatal("Error connecting to MongoDB")
		panic(err)
	}
}