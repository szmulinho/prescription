package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func (h *handlers) GetPrescriptionsForPatient(w http.ResponseWriter, r *http.Request) {
	patientID := mux.Vars(r)["patient"]
	var prescriptionsForPatient []model.CreatePrescInput

	if err := h.db.Where("patient = ?", patientID).Find(&prescriptionsForPatient).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prescriptionsForPatient)
}
