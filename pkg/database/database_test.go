package database

import "testing"

func TestInitDB(t *testing.T) {
	err := InitDB()

	if err != nil {
		panic("Failed to connect to database")
	}
}