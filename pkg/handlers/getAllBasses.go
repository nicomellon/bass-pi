package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nicomellon/bass-pi/pkg/models"

	"github.com/nicomellon/bass-pi/pkg/sqldb"
)

// GetAllBasses responds with the list of all basses as JSON.
func GetAllBasses(w http.ResponseWriter, r *http.Request) {
	
	q := `
	SELECT id, name 
	FROM basses;
	`
	
	rows, err := sqldb.DB.Query(q)
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
	
	json.NewEncoder(w).Encode(basses)
}
