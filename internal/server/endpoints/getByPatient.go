package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func (h *handlers) GetPrescriptionsForPatient(w http.ResponseWriter, r *http.Request) {
	patient := mux.Vars(r)["patient"]

	if err := h.db.Where("patient = ?", patient).Find(&model.Prescriptions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Prescriptions)
}
