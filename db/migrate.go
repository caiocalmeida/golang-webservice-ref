package main

import (
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		godotenv.Load(".env.local")
	} else {
		godotenv.Load(".env")
	}

	m, err := migrate.New(
		"file://db/migrations",
		os.Getenv("POSTGRES_CONNECTION_STRING"),
	)

	if err != nil {
		log.Fatal("Error initiating migration: ", err)
	}

	target, err := strconv.ParseUint(os.Getenv("MIGRATION_TARGET_VER"), 10, 0)

	if err != nil {
		log.Fatal("Error finding migration target version: ", err)
	}

	if err = m.Migrate(uint(target)); err != nil {
		log.Fatal("Error during migration: ", err)
	}
}
