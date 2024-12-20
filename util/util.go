package util

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadDotEnv(filepath string) {
	err := godotenv.Load(filepath)
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

func GetEnv(key string, defaultValue ...string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return ""
}

func GetBeerConfig() string {
	result := GetEnv("BEER_CONFIG_PATH")

	if result == "" {
		exPath := GetExecutablePath()

		result = filepath.Join(exPath, "config.beer.yml")
	}

	return result
}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
