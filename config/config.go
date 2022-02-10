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

func getAppConfig() AppConfig {
	port := Defaults.GetInt(os.Getenv("PORT"), 3000)
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

var App AppConfig

func Init() {
	App = getAppConfig()
}
