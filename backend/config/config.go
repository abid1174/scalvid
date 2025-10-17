package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version    string
	HttpPort   int
	GoEnv      string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func loadConfig() {
	// Load .env file if it exists (for local development)
	// In Docker, environment variables are provided by docker-compose
	_ = godotenv.Load()

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

	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	configurations = Config{
		Version:    version,
		HttpPort:   httpPort,
		GoEnv:      goEnv,
		DbPort:     dbPort,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbName:     dbName,
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}
