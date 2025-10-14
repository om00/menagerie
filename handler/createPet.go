package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/om00/menagerie/models"
)

func (app App) CreatePet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		log.Printf("ERROR: Invalid json request , Error: %v", err)
		app.respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	defer r.Body.Close()

	if err := app.Validator.Struct(pet); err != nil {
		log.Printf("ERROR: Request fields are missing , Error: %v", err)
		app.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := app.Db.CreatePet(pet)
	if err != nil {
		log.Printf("Error while inserting data into Pet table , Error: %v", err)
		app.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pet.ID = id
	app.respondWithJSON(w, http.StatusCreated, pet)
}
