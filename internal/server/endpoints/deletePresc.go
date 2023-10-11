package endpoints

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
	"strconv"
)

func (h *handlers) DeletePrescription(w http.ResponseWriter, r *http.Request) {
	prescIDStr := mux.Vars(r)["id"]
	prescID, err := strconv.ParseInt(prescIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingPrescription model.CreatePrescInput
	result := h.db.First(&existingPrescription, prescID)
	if result.Error != nil {
		http.Error(w, "Prescription not found", http.StatusNotFound)
		return
	}

	result = h.db.Delete(&existingPrescription)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "The prescription with ID %v has been deleted successfully", prescID)
}
