package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func TestInitEnv(t *testing.T) {
	file, err := os.Create(".env")
	if err != nil {
		panic("Failed to create .env file for testing")
	}
    defer file.Close()

	var testInput string = "KEY=VALUE";
    buf := []byte(testInput)

	_, err = file.Write(buf)
	if err != nil {
		panic("Could not write to .env file")
	}

	err = InitEnv()

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

func TestValidate(t *testing.T) {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

    bodyMap := map[string]string{
        "name":				"John Doe",
        "cuny_email":		"john.doe@cuny.edu",
        "personal_email":	"john.doe@gmail.com",
        "discord":			"johndoe#1234",
        "emplid":			"12345678",
    }	
	newBody, _ := json.Marshal(bodyMap)
	c.Request().SetBody(newBody)
	c.Request().Header.SetContentType("application/json")
	
	var users AddUsersBody
	err := Validate(c, &users)

	if err != nil {
		fmt.Println("ruh roh raggy")
		panic(err.Error())
	}

	var passed bool = false
	passed = users.Name				==	"John Doe"
	passed = users.CunyEmail		==	"john.doe@cuny.edu"
	passed = users.PersonalEmail	==	"john.doe@gmail.com"
    passed = users.Discord			==	"johndoe#1234"
	passed = users.Emplid			==	"12345678"

	if passed != true {
		panic("variables did not save correctly")
	}
}