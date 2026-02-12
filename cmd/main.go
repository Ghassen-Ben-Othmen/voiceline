package main

import (
	"log"

	"github.com/G-b-o/voice-line/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.POST("/upload", handlers.UploadAudio)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
