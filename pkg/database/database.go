package database

import (
	"byteserver/pkg/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// errors ignored cause tests already covered it
	host, _ := utils.GetEnv("POSTGRESQL_DB_HOST")
	user, _ := utils.GetEnv("POSTGRESQL_DB_USER")
	password, _ := utils.GetEnv("POSTGRESQL_DB_PASSWORD")
	dbname, _ := utils.GetEnv("POSTGRESQL_DB")
	port, _ := utils.GetEnv("POSTGRESQL_DB_PORT")

	conn_string := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(conn_string))

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected...")
}
