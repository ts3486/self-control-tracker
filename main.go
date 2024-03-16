package main

import (
	"aimeechat_api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	defer database.DB.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on
}