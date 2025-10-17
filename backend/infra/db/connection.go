package db

import (
	"fmt"
	"log"
	"os"
	"scalvid/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	cfg := config.GetConfig()

	// Use "scalvid-db" as host when running in Docker, "localhost" otherwise
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
}

func NewConnection() *sqlx.DB {
	log.Println("GetConnectionString", GetConnectionString())
	db, err := sqlx.Connect("postgres", GetConnectionString())
	if err != nil {
		log.Println("Error connecting to database")
		log.Fatal(err)
	}
	return db
}
