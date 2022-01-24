package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicomellon/bass-pi/models"
)

func GetManufacturerByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var mfct models.Manufacturer
		id := c.Param("id")
		
		row := db.QueryRow("SELECT * FROM manufacturers WHERE id = ?;", id)
		if err := row.Scan(&mfct.ID, &mfct.Name, &mfct.FoundedYear, &mfct.Nationality, &mfct.Logo); err != nil {
			if err == sql.ErrNoRows {
				// return alb, fmt.Errorf("albumsById %d: no such album", id)
				panic(err)
			}
			// return alb, fmt.Errorf("albumsById %d: %v", id, err)
		}
		// return alb, nil

		c.JSON(http.StatusOK, mfct)

	}

}