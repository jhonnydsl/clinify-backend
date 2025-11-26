package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jhonnydsl/clinify-backend/src/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading enviroment variables")
		return
	}
	
	err = repository.Connect()
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	} else {
		log.Println("connection estabilished")
	}
	defer repository.DB.Close()

	app := gin.Default()

	app.Run(":8080")
}