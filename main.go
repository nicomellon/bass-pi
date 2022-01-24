package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/nicomellon/bass-pi/handlers"
)

var db *sql.DB

func main() {

	// Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "database-1.chdxg6xj6r67.eu-central-1.rds.amazonaws.com:3306",
        DBName: "bass_pi",
		AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")



	r := gin.Default()

	r.GET("/api/manufacturers", handlers.GetManufacturers(db))

	r.Run()
}
