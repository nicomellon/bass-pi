package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nicomellon/bass-pi/pkg/mocks"

	"github.com/nicomellon/bass-pi/pkg/models"
)

// postBasses adds a bass from JSON received in the request body.
func PostBasses(ctx *gin.Context) {
	var newBass models.Bass

	// Call BindJSON to bind the received JSON to newBass.
	if err := ctx.BindJSON(&newBass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new bass to the slice.
	mocks.Basses = append(mocks.Basses, newBass)
	ctx.IndentedJSON(http.StatusCreated, newBass)
}
