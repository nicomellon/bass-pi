package main

import (
	"github.com/gin-gonic/gin"

	"github.com/nicomellon/bass-pi/pkg/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/basses", handlers.GetAllBasses)
	router.GET("/basses/:id", handlers.GetBassByID)
	router.POST("/basses", handlers.PostBasses)

	router.Run("localhost:8080")
}