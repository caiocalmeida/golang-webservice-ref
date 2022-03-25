package data

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Config() {
	newDb, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECTION_STRING")))

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db = newDb
}
