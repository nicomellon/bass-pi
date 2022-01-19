package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nicomellon/bass-pi/pkg/handlers"
	"github.com/nicomellon/bass-pi/pkg/sqldb"
)

func main() {
	sqldb.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/basses", handlers.GetAllBasses)
	// router.GET("/basses/random", handlers.GetRandomBass)
	// TODO router.GET("/basses/:id", handlers.GetBassByID)
	// TODO router.POST("/basses", handlers.PostBasses)
	// router.PUT("/basses/:id", handlers.UpdateBassByID)
	// router.DELETE("/basses/:id", handlers.DeleteBassByID)

	// router.GET("/manufacturers", handlers.GetAllManufacturers)
	// router.GET("/manufacturers/:id/basses", handlers.GetBassesFromManufacturer)
	// router.GET("/manufacturers/:id", handlers.GetManufacturerByID)
	// router.POST("/manufacturers", handlers.PostManufacturers)
	// router.PUT("/manufacturers/:id", handlers.UpdateManufacturerByID)
	// router.DELETE("/manufacturers/:id", handlers.DeleteManufacturerByID)
	
	http.ListenAndServe(":8080", router)
}