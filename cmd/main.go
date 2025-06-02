package main

import (
	"log"

	"github.com/a5512167086/url-shortener/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	_ = db
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080")
}
