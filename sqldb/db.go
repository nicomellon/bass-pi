package sqldb

import "database/sql"

// DB is a global variable to hold db connection
var DB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() {
	db, err := sql.Open("mysql", "admin:12341234@tcp(database-1.chdxg6xj6r67.eu-central-1.rds.amazonaws.com:3306)/bass_pi")
	if err != nil {
		panic(err)
	}

	DB = db
}