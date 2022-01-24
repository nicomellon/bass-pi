package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"), // export DBUSER=admin
		Passwd:               os.Getenv("DBPASS"), // export DBPASS=12341234
		Net:                  "tcp",
		Addr:                 "database-1.chdxg6xj6r67.eu-central-1.rds.amazonaws.com:3306",
		DBName:               "bass_pi",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database!")

	return db
}