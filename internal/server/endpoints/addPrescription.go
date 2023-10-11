package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func (h *handlers) AddPrescription(w http.ResponseWriter, r *http.Request) {
	var newPrescription model.Prescription

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &newPrescription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := h.db.Create(&newPrescription)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	for _, singlePrescription := range model.Prescriptions {
		fmt.Println(singlePrescription)
		if singlePrescription.PreID == newPrescription.PreID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Drug %s already exist", newPrescription.PreID)})
			return
		}
	}

	model.Prescriptions = append(model.Prescriptions, newPrescription)

	fmt.Printf("added new drug %+v\n", newPrescription)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPrescription)
}
