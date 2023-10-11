package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *handlers) UpdatePrescription(w http.ResponseWriter, r *http.Request) {
	prescIDStr := mux.Vars(r)["id"]
	prescID, err := strconv.ParseInt(prescIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedPresc model.Prescription
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Kindly enter data with the drug name and price only in order to update", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &updatedPresc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var existingPresc model.Prescription
	result := h.db.First(&existingPresc, prescID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	existingPresc.Patient = updatedPresc.Patient
	existingPresc.Drugs = updatedPresc.Drugs
	existingPresc.Expiration = updatedPresc.Expiration

	result = h.db.Save(&existingPresc)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(existingPresc)
}
