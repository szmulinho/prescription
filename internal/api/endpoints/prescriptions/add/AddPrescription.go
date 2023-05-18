package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/prescription/internal/database"
	"github.com/szmulinho/prescription/internal/model"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func CreatePrescription(w http.ResponseWriter, r *http.Request) {

	var newPresc model.CreatePrescInput

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
	err = json.Unmarshal(buf.Bytes(), &newPresc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
		log.Printf("Invalid body")
	}

	result := database.DB.Create(&newPresc)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	for _, singlePresc := range model.Prescs {
		fmt.Println(singlePresc)
		if singlePresc.PreID == model.Prescription.PreID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Prescription %model.already exist", model.Prescription.PreID)})
			return
		}
	}

	fmt.Printf("created new prescription %+v\n", model.Prescription)
	log.Printf("%+v", model.Prescription)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.Prescription)
}
