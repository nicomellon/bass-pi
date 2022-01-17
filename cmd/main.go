package main

import (
	"github.com/gin-gonic/gin"

	"github.com/nicomellon/bass-pi/pkg/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/basses", handlers.GetAllBasses)
	// router.GET("/basses/random", handlers.GetRandomBass)
	router.GET("/basses/:id", handlers.GetBassByID)
	router.POST("/basses", handlers.PostBasses)
	// router.PUT("/basses/:id", handlers.UpdateBassByID)
	// router.DELETE("/basses/:id", handlers.DeleteBassByID)

	// router.GET("/manufacturers", handlers.GetAllManufacturers)
	// router.GET("/manufacturers/:id/basses", handlers.GetBassesFromManufacturer)
	// router.GET("/manufacturers/:id", handlers.GetManufacturerByID)
	// router.POST("/manufacturers", handlers.PostManufacturers)
	// router.PUT("/manufacturers/:id", handlers.UpdateManufacturerByID)
	// router.DELETE("/manufacturers/:id", handlers.DeleteManufacturerByID)
	
	router.Run("localhost:8080")
}