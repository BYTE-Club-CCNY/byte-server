package database

import (
	"byteserver/pkg/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PageSize int = 10;

func InitDB() error {
	var err error
	utils.InitEnv()

	host, _ := utils.GetEnv("POSTGRESQL_DB_HOST")
	user, _ := utils.GetEnv("POSTGRESQL_DB_USER")
	password, _ := utils.GetEnv("POSTGRESQL_DB_PASSWORD")
	dbname, _ := utils.GetEnv("POSTGRESQL_DB")
	port, _ := utils.GetEnv("POSTGRESQL_DB_PORT")

	conn_string := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(conn_string))

	if err != nil {
		return err
	}

	fmt.Println("Database connected...")
	return nil
}

func Paginate(page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * PageSize
		return db.Offset(offset).Limit(PageSize)
	}
}