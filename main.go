package main

import (
	"os"

	api "github.com/caiocalmeida/go-webservice-ref/internal/api/gin"
	"github.com/caiocalmeida/go-webservice-ref/internal/data"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		godotenv.Load(".env.local")
	} else {
		godotenv.Load(".env")
	}

	r := api.NewRouter(api.NewUserController(data.NewUserRepository(data.NewDB())))
	r.Start()
}
