package config

import (
	"log"
	Defaults "micrach-gateway/defaults"
	"os"
)

type AppConfig struct {
	// Port is the port to listen on
	Port int
	// Env is the environment to run in
	Env string
	// API key is the secret key to use for authentication
	ApiKey string
}

type DbConfig struct {
	// Url is the database url
	Url string
}

// Gets the application configuration
func getAppConfig() AppConfig {
	port := Defaults.GetInt(os.Getenv("PORT"), 3001)
	env := Defaults.GetString(os.Getenv("ENV"), "development")
	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Panicln("API_KEY is not set")
	}

	return AppConfig{
		Port:   port,
		Env:    env,
		ApiKey: apiKey,
	}
}

// Gets the database configuration
func getDbConfig() DbConfig {
	url := Defaults.GetString(os.Getenv("MONGO_URL"), "localhost/micrach-gateway")

	return DbConfig{
		Url: url,
	}
}

var App AppConfig
var Db DbConfig

// Initializes the configuration. Should be called before any other functions.
func Init() {
	App = getAppConfig()
	Db = getDbConfig()
}
