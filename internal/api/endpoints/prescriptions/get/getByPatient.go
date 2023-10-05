package get

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/database"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func GetPrescriptionsForPatient(w http.ResponseWriter, r *http.Request) {
	patientID := mux.Vars(r)["patient"]
	var prescriptionsForPatient []model.CreatePrescInput

	if err := database.DB.Where("patient = ?", patientID).Find(&prescriptionsForPatient).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prescriptionsForPatient)
}
