package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nicomellon/bass-pi/database"
	"github.com/nicomellon/bass-pi/handlers"
)

func main() {
	db := database.Connection()
	r := gin.Default()

	// manufacturers routes
	m := r.Group("/api/manufacturers")
	{
		m.GET("/", handlers.GetManufacturers(db))
		m.POST("/", handlers.PostManufacturer(db))
		m.GET("/:id", handlers.GetManufacturerByID(db))
	}

	// basses routes
	b := r.Group("/api/basses")
	{
		b.GET("/", handlers.GetBasses(db))
	}

	r.Run()
}
