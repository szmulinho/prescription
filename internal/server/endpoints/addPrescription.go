package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/prescription/internal/model"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func (h *handlers) CreatePrescription(w http.ResponseWriter, r *http.Request) {

	var newPresc model.Prescription

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
	err = json.Unmarshal(buf.Bytes(), &newPresc)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	result := h.db.Create(&newPresc)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	for _, singlePresc := range model.Prescriptions {
		fmt.Println(singlePresc)
		if singlePresc.PreID == model.Presc.PreID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Prescription %model.already exist", model.Presc.PreID)})
			return
		}
	}

	fmt.Printf("created new prescription %+v\n", model.Presc)
	log.Printf("%+v", model.Presc)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.Presc)
}
