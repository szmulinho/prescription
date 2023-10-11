package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
	"strconv"
)

func (h *handlers) GetOnePrescription(w http.ResponseWriter, r *http.Request) {
	prescIDStr := mux.Vars(r)["id"]
	PreID, err := strconv.ParseInt(prescIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var prescription model.CreatePrescInput
	if err := h.db.First(&prescription, PreID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(prescription)
}
