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

func GetBeerConfigFile() string {
	// create a default config file if it does not exist
	defaultConfig := []byte("apps:")
	configFile := filepath.Join(GetBeerConfigDir(), "config.beer.yml")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err = os.WriteFile(configFile, defaultConfig, 0644)
		if err != nil {
			log.Fatal("Error creating default config file")
		}
	}
	return GetEnv("BEER_CONFIG_PATH", configFile)
}

func GetBeerConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting user home directory")
	}
	// Check if the .beer directory exists
	beerDir := filepath.Join(homeDir, ".beer")
	if _, err := os.Stat(beerDir); os.IsNotExist(err) {
		// Create the .beer directory if it does not exist
		err = os.MkdirAll(beerDir, 0755)
		if err != nil {
			log.Fatal("Error creating .beer directory")
		}
		err = os.MkdirAll(filepath.Join(beerDir, "projects"), 0755)
		if err != nil {
			log.Fatal("Error creating .beer/projects directory")
		}
	}
	return beerDir
}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
