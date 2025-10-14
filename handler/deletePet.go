package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app App) DeletePet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	petIdStr := vars["id"]

	if petIdStr == "" {
		app.respondWithError(w, http.StatusBadRequest, "Pet id is required")
		return
	}

	petID, err := strconv.ParseInt(petIdStr, 10, 64)
	if err != nil {
		log.Printf("ERROR: Invalid pet ID - PetID: %s, Error: %v", petIdStr, err)
		app.respondWithError(w, http.StatusBadRequest, "id should be and interger")
		return
	}

	if err := app.Db.DeletePet(petID); err != nil {
		log.Printf("Error while delete pet from table:PetID: %d Error: %v", petID, err)
		app.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	payload := map[string]interface{}{
		"success": true,
		"message": "Pet deleted successfully",
		"id":      petID,
	}
	app.respondWithJSON(w, http.StatusOK, payload)
}
