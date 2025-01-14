package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

func InitEnv() error {
	err := godotenv.Load()
	if err != nil {
	  panic("Error loading .env file")
	}

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