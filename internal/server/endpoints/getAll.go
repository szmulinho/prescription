package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func (h *handlers) GetAllPrescriptions(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Find(&model.Prescriptions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Prescriptions)
}
