package data

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECTION_STRING")))

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	return db
}
