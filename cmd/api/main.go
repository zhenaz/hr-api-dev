package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"codeid.hr-api/pkg/database"
)

func main() {
	// initialize database connection
	_, err := database.SetupDB()
	if err != nil {
		log.Fatal("Failed to Connect %w", err)
	}

	// setup router
	router := gin.Default()

	// call handlers
	router.GET("/", helloworldHandler)

	//run server
	router.Run(":8080")
}

func helloworldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
		"status":  "running",
	})
}