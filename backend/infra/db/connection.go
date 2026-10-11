package db

import (
	"fmt"
	"log"
	"scalvid/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(dbConfig *config.DBConfig) string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	if !dbConfig.EnableSSLMode {
		connectionString += " sslmode=disable"
	}

	return connectionString
}

func NewConnection(dbConfig *config.DBConfig) *sqlx.DB {
	connectionString := GetConnectionString(dbConfig)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Println("Error connecting to database")
		log.Fatal(err)
	}
	log.Printf("Database connected (%s:%d/%s)", dbConfig.Host, dbConfig.Port, dbConfig.Name)
	return db
}
