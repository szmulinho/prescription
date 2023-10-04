package delete

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/database"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
	"strconv"
)

func DeletePrescription(w http.ResponseWriter, r *http.Request) {
	prescIDStr := mux.Vars(r)["id"]
	prescID, err := strconv.ParseInt(prescIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingPrescription model.CreatePrescInput
	result := database.DB.First(&existingPrescription, prescID)
	if result.Error != nil {
		http.Error(w, "Prescription not found", http.StatusNotFound)
		return
	}

	result = database.DB.Delete(&existingPrescription)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "The prescription with ID %v has been deleted successfully", prescID)
}
