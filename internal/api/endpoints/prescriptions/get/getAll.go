package get

import (
	"encoding/json"
	"github.com/szmulinho/prescription/database"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func GetAllPrescriptions(w http.ResponseWriter, r *http.Request) {
	var prescriptions []model.CreatePrescInput
	if err := database.DB.Find(&prescriptions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prescriptions)
}
