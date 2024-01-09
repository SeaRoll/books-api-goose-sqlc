package main

import (
	"books-api/api"
	"books-api/db"
	"log"

	"github.com/joho/godotenv"
)

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initConfig()
	db.Connect()
	defer db.Disconnect()
	apio := api.New()
	api.Run(apio)
}
