package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nicomellon/bass-pi/pkg/models"

	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllBasses responds with the list of all basses as JSON.
func GetAllBasses(ctx *gin.Context) {
	fmt.Println("Attempting to connect to database...")

	db, err := sql.Open("mysql", "admin:12341234@tcp(database-1.chdxg6xj6r67.eu-central-1.rds.amazonaws.com:3306)/bass_pi")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")
	defer db.Close()

	q := `
		SELECT id, name 
		FROM basses;
	`

	rows, err := db.Query(q)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var basses []string

	for rows.Next() {
		var bass models.Bass
        // for each row, scan the result into our bass composite object
        err = rows.Scan(&bass.ID, &bass.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
		basses = append(basses, bass.Name)		
	}
	
	fmt.Println("Database connection closed")
	ctx.IndentedJSON(http.StatusOK, basses)
}
