package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicomellon/bass-pi/pkg/mocks"
)

// GetAllBasses responds with the list of all basses as JSON.
func GetAllBasses(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, mocks.Basses)
}
