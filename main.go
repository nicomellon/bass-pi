package main

import (
	"github.com/gin-gonic/gin"

	"github.com/nicomellon/bass-pi/handlers"
	"github.com/nicomellon/bass-pi/sqldb"
)

func main() {
	sqldb.ConnectDB()
	
	router := gin.Default()
	router.GET("/api/manufacturers", handlers.GetManufacturers)

	router.Run("localhost:8080")
}
