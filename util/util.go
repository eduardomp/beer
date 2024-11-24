package util

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func SetEnv(key string, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		log.Fatal("Error setting env variable")
	}
}

func Config(key string, defaultValue string) string {
	LoadDotEnv()
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	configPathFromContext := filepath.Join(exPath, defaultValue)

	if _, err := os.Stat(configPathFromContext); err == nil {
		return configPathFromContext
	}

	return defaultValue
}
