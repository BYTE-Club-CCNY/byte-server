package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InitEnv() error {
    file, err := os.Open(".env")
    if err != nil {
        return errors.New(".env file not found in root of application directory")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	
	var line []string

    for scanner.Scan() {
        line = strings.Split(scanner.Text(), "=")

		if len(line) > 2 {
			return errors.New(".env file is formated incorrectly")
		}
		os.Setenv(line[0], line[1])
    }
	return nil
}

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("ENV value %s missing", key)
	}

	return value, nil;
}

func PrintParams(c *fiber.Ctx) {
	for key, value := range c.AllParams() {
		fmt.Printf("%s: %s\t", key, value)
	}
	fmt.Print("\n")
}

func PrintQueries(c *fiber.Ctx) {
	for key, value := range c.Queries() {
		fmt.Printf("%s: %s\t", key, value)
	}
	fmt.Print("\n")
}

func Validate(c *fiber.Ctx, dest interface{}) error {
	var validate = validator.New()

	fmt.Println(dest)

	if err := c.BodyParser(dest); err != nil {
        return fmt.Errorf("body parsing failed: %w", err)
	}

	if err := validate.Struct(dest); err != nil {
        return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}