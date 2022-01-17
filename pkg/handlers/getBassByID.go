package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nicomellon/bass-pi/pkg/mocks"
)

// getBassByID locates the bass whose ID value matches the id
// parameter sent by the client, then returns that bass as a response.
func GetBassByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Loop through the list of basses, looking for
	// an album whose ID value matches the parameter.
	for _, a := range mocks.Basses {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "bass not found"})
}
