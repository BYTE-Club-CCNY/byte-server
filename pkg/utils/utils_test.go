package utils

import (
	"fmt"
	"testing"
)

func TestInitEnv(t *testing.T) {
	err := InitEnv()

	if err != nil {
		panic(err)
	}
}

func TestEnvironmentValuesExist(t *testing.T) {
	_ = InitEnv()

	keys := [5]string{"POSTGRESQL_DB_HOST", "POSTGRESQL_DB_USER", 
	"POSTGRESQL_DB_PASSWORD", "POSTGRESQL_DB", "POSTGRESQL_DB_PORT"};

	for _, key := range keys {
		_, error := GetEnv(key)

		if error != nil {
			panic(fmt.Sprintf("Missing %s", key));
		}
	}
}