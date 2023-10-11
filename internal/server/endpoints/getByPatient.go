package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/common/model"
	"net/http"
)

func (h *handlers) GetPrescriptionsForPatient(w http.ResponseWriter, r *http.Request) {
	patientID := mux.Vars(r)["patient"]

	if err := h.db.Where("patient = ?", patientID).Find(&model.Prescriptions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Prescriptions)
}
