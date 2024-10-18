package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
}

func Config(key string, defaultValue string) string {

	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}
