package mongodb

import (
	"testing"
)

func TestInitDB(t *testing.T) {

	err := Connect()

	if err != nil {
		panic(err)
	}
}