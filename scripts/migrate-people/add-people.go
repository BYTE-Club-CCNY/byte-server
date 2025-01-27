package main

import (
	"os"
	"fmt"
	"encoding/csv"
	schema "byteserver/pkg/schemas"
	"byteserver/pkg/database"
	"byteserver/pkg/utils"
)

func main () {
	utils.InitEnv()
	database.InitDB()
	fmt.Println("Database connection:", database.DB)
	file, err := os.Open("./scripts/migrate-people/data.csv")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	for _, record := range records {
		user := new(schema.User)
		user.PersonalEmail = record[0]
		user.FirstName = record[1]
		if record[2] != "" {
			user.MiddleName = record[2]
		}
		user.LastName = record[3]
		user.CunyEmail = record[4]
		user.Emplid = record[5]
		user.Discord = record[6]
		fmt.Println(user)

		res := database.DB.Create(&user)

		if res.Error != nil{
			fmt.Println("adding user error for: ", record)
		}
    }
	fmt.Println("Successfully added all people")
}