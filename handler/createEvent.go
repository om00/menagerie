package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/om00/menagerie/models"
)

func (app App) CreatePetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	petIdStr := vars["id"]

	if petIdStr == "" {
		log.Printf("Pet Id is required")
		app.respondWithError(w, http.StatusBadRequest, "Pet id is required")
		return
	}

	petID, err := strconv.ParseInt(petIdStr, 10, 64)
	if err != nil {
		log.Printf("ERROR: Invalid pet ID - PetID: %s, Error: %v", petIdStr, err)
		app.respondWithError(w, http.StatusBadRequest, "id should be and interger")
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Printf("ERROR: Invalid json request - PetID: %d, Error: %v", petID, err)
		app.respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	defer r.Body.Close()

	if err := app.Validator.Struct(event); err != nil {
		log.Printf("ERROR: request fiedls are missing - PetID: %d, Error: %v", petID, err)
		app.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	exist, err := app.Db.IsPetExist(petID)
	if err != nil {
		log.Printf("Error while checking Pet in table - PetID: %d, Error: %v", petID, err)
		app.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !exist {
		log.Printf("This PetId does not exist - PetID: %d", petID)
		app.respondWithError(w, http.StatusBadRequest, fmt.Sprintf("This PetId does not exist - PetID: %d", petID))
		return
	}

	event.PetID = petID
	eventId, err := app.Db.CreateEvent(event)
	if err != nil {
		log.Printf("Error while Insreting event for Pet - PetID: %d, Error: %v", petID, err)
		app.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	event.ID = eventId
	app.respondWithJSON(w, http.StatusCreated, event)
}
