package get

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/database"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
	"strconv"
)

func GetOnePrescription(w http.ResponseWriter, r *http.Request) {
	prescIDStr := mux.Vars(r)["id"]
	PreID, err := strconv.ParseInt(prescIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var prescription model.CreatePrescInput
	if err := database.DB.First(&prescription, PreID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(prescription)
}
