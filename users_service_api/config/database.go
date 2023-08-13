package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseInterface interface {
	LoadEnvFile()
	ConnectToDB()
	MigrateTables()
}

func NewDatabase() *Database {
	return &Database{}
}

type Database struct {
	DB *gorm.DB
}

func (db *Database) LoadEnvFile() {
	enverror := godotenv.Load()
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
}

func (db *Database) ConnectToDB() {
	var err error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	database_string := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", host, user, password, db_name, port)
	db.DB, err = gorm.Open(postgres.Open(database_string), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}

}
