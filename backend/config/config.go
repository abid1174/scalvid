package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version  string
	HttpPort int
	GoEnv    string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get VERSION from .env file
	version := os.Getenv("VERSION")
	if version == "" {
		version = "1.0.0"
	}

	// Get HTTP_PORT from .env file
	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		log.Fatal("HTTP_PORT is not set")
	}

	// Convert HTTP_PORT to int
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		log.Fatal("Error converting HTTP_PORT to int")
	}

	// Get GO_ENV from .env file
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		log.Fatal("GO_ENV is not set")
	}

	configurations = Config{
		Version:  version,
		HttpPort: httpPort,
		GoEnv:    goEnv,
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}
