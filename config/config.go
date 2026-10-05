package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version      string
	HttpPort     int
	GoEnv        string
	JwtSecretKey string
	DB           *DBConfig
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

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Fatal("JWT Secret Key is required!")
	}

	// DB CONFIG
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB Host is required!")
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		log.Fatal("DB Port is required!")
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatal("DB Port must be integer!")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB User is required!")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB Password is required!")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB Name is required!")
	}

	enableSSLMode := os.Getenv("ENABLE_SSL_MODE")
	isSSLModeEnabled, err := strconv.ParseBool(enableSSLMode)
	if err != nil {
		log.Fatal("Invalid enableSSLMode format!", err)
	}

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          dbPort,
		Name:          dbName,
		User:          dbUser,
		Password:      dbPassword,
		EnableSSLMode: isSSLModeEnabled,
	}

	configurations = &Config{
		Version:      version,
		HttpPort:     httpPort,
		GoEnv:        goEnv,
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
