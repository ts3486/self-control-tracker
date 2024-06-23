package main

import (
	"sct-api/database"
	"sct-api/routes"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
        log.Fatal("Error loading .env file")
    }
	r := gin.Default()
	database.Connect()
	defer database.DB.Close()

	routes.RegisterRoutes(r, database.DB)
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on
}