package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicomellon/bass-pi/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetManufacturers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		rows, err := db.Query("SELECT * FROM manufacturers;")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
	
		var manufacturers []models.Manufacturer
		for rows.Next() {
			var manufacturer models.Manufacturer
			err = rows.Scan(&manufacturer.ID, &manufacturer.Name, &manufacturer.FoundedYear, &manufacturer.Nationality, &manufacturer.Logo) // alguna manera mejor de hacer esto?
			if err != nil {
				panic(err)
			}
			manufacturers = append(manufacturers, manufacturer)
		}
		c.JSON(http.StatusOK, manufacturers)
	}

}
