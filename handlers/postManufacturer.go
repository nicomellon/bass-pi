package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicomellon/bass-pi/models"
)

func PostManufacturer(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var mfct models.Manufacturer
		c.Bind(&mfct)

		result, err := db.Exec("INSERT INTO manufacturers (name, founded_year, nationality, logo) VALUES (?, ?, ?, ?)", mfct.Name, mfct.FoundedYear, mfct.Nationality, mfct.Logo)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)

		c.JSON(http.StatusCreated, mfct)
		
		/* row := db.QueryRow("SELECT * FROM manufacturers WHERE id = ?;", id)
		if err := row.Scan(&mfct.ID, &mfct.Name, &mfct.FoundedYear, &mfct.Nationality, &mfct.Logo); err != nil {
			if err == sql.ErrNoRows {
				// return alb, fmt.Errorf("albumsById %d: no such album", id)
				panic(err)
			}
			// return alb, fmt.Errorf("albumsById %d: %v", id, err)
		}
		// return alb, nil

		c.JSON(http.StatusOK, mfct)
 */
	}
    /* result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil */
}