package main

import (
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-ybaspinar/pkg/handlers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	handlers.ApiRoutes()
}
