package main

import (
	"chatty/backend/database"
	"chatty/backend/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	server.Init()
	database.Init()
}
