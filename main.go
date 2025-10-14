package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/om00/menagerie/database"
	"github.com/om00/menagerie/handler"
)

func main() {

	appPort := os.Getenv("APP_PORT_INSIDE")
	ExtAppPort := os.Getenv("APP_PORT")
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatal("error while connectin to the database", err)
	}

	defer db.MainDB.Close()
	database.RunMigrations()

	app := handler.App{Db: &db, Validator: validator.New()}
	r := mux.NewRouter()

	r.HandleFunc("/pets", app.CreatePet).Methods("POST")
	r.HandleFunc("/pets", app.GetPetsList).Methods("GET")
	r.HandleFunc("/pets/{id}", app.GetPetEvents).Methods("GET")
	r.HandleFunc("/pets/{id}", app.UpdatePet).Methods("PUT")
	r.HandleFunc("/pets/{id}", app.DeletePet).Methods("DELETE")
	r.HandleFunc("/pets/{id}", app.CreatePetEvent).Methods("POST")

	fmt.Printf("Server running on extrenal port:%s and inside contianer port:%s\n", ExtAppPort, appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}
