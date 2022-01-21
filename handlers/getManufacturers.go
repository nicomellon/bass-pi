package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicomellon/bass-pi/models"
	"github.com/nicomellon/bass-pi/sqldb"

	_ "github.com/go-sql-driver/mysql"
)

func GetManufacturers(c *gin.Context) {

	rows, err := sqldb.DB.Query("SELECT * FROM manufacturers;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var manufacturers []models.Manufacturer
	for rows.Next() {
		var manufacturer models.Manufacturer
		err = rows.Scan(&manufacturer.ID, &manufacturer.Name, &manufacturer.FoundedYear, &manufacturer.Nationality, &manufacturer.Logo)
		if err != nil {
			panic(err)
		}
		manufacturers = append(manufacturers, manufacturer)
	}
	c.IndentedJSON(http.StatusOK, manufacturers)
}
