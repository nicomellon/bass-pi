package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicomellon/bass-pi/models"
)

func GetBasses(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM basses;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var basses []models.Bass
		for rows.Next() {
			var b models.Bass
			err = rows.Scan(&b.ID, &b.ManufacturerID, &b.Name, &b.Strings, &b.LaunchYear, &b.Image)
			if err != nil {
				panic(err)
			}
			basses = append(basses, b)
		}
		c.JSON(http.StatusOK, basses)
	}
} 