package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/om00/menagerie/models"
)

func (app App) GetPetsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	species := r.URL.Query().Get("species")

	filter := models.GetFilter{Species: species}

	pets, err := app.Db.GetPetsList(filter)
	if err != nil {
		log.Printf("Error while fetching data from Pet table , Error: %v", err)
		app.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(pets) == 0 {
		app.respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"message": "No pets found",
			"pets":    []string{},
		})
		return
	}

	for ind, pet := range pets {
		if pet.BirthStr != nil {
			t, _ := time.Parse("2006-01-02 15:04:05", *pet.BirthStr)
			pets[ind].Birth = &t
		}

		if pet.DeathStr != nil {
			t, _ := time.Parse("2006-01-02 15:04:05", *pet.DeathStr)
			pets[ind].Death = &t
		}
	}

	app.respondWithJSON(w, http.StatusOK, pets)
}
