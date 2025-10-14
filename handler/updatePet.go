package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/om00/menagerie/models"
)

func (app App) UpdatePet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		log.Printf("Pet Id is Misssing")
		app.respondWithError(w, http.StatusBadRequest, "Pet id is required")
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error in parsinge the PetId ,PetID-%s, Error: %v", id, err)
		app.respondWithError(w, http.StatusBadRequest, "id should be and interger")
		return
	}

	var updateReq models.UpdatePetReq
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		log.Printf("InValid Update Request, Error: %v", err)
		app.respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if _, err := app.Db.UpdatePet(idInt, updateReq); err != nil {
		log.Printf("Error while updating  the Pet details ,PetID-%v, Error: %v", idInt, err)
		app.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	filter := models.GetFilter{ID: idInt}
	pets, err := app.Db.GetPetsList(filter)
	if err != nil {
		log.Printf("Error while getting  the Pet details ,PetID-%v, Error: %v", idInt, err)
		app.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	app.respondWithJSON(w, http.StatusOK, pets[0])
}
