package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/om00/menagerie/models"
)

func (app App) GetPetEvents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petIdStr := vars["id"]

	if petIdStr == "" {
		app.respondWithError(w, http.StatusBadRequest, "Pet id is required")
		return
	}

	petID, err := strconv.ParseInt(petIdStr, 10, 64)
	if err != nil {
		log.Printf("Error while parsing the PetId ,PetId: %s, Error: %v", petIdStr, err)
		app.respondWithError(w, http.StatusBadRequest, "id should be and interger")
		return
	}

	events, err := app.Db.GetPetEvents(petID)
	if err != nil {
		log.Printf("Error while getting the Pet Events Details ,PetId: %v, Error: %v", petID, err)
		app.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(events) == 0 {
		log.Printf("Pet does not exist with this PetId ,PetId: %v", petID)
		app.respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Pet does not exist with this PetId ,PetId: %v", petID))
		return
	}

	for ind, event := range events {
		if event.DateStr != nil {
			t, _ := time.Parse("2006-01-02 15:04:05", *event.DateStr)
			events[ind].Date = t
		}
	}

	PetEvents := models.PetEvents{ID: events[0].PetID, Name: events[0].Name, Event: events}

	app.respondWithJSON(w, http.StatusOK, PetEvents)
}
