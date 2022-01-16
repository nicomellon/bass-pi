package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Bass struct {
	Name string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Launch_year string `json:"launch_year"`
}

var basses = []Bass{
	{ Name: "Precision Bass", Manufacturer: "Fender", Launch_year: "1951" },
	{ Name: "Jazz Bass", Manufacturer: "Fender", Launch_year: "1960" },
	{ Name: "StingRay", Manufacturer: "Ernie Ball Music Man", Launch_year: "1976" },
}

func main() {
	server := gin.Default()
	server.GET("/basses", getBasses)

	server.Run("localhost:8080")
}

// getBasses responds with the list of all basses as JSON.
func getBasses(ctx *gin.Context) {
    ctx.IndentedJSON(http.StatusOK, basses)
}