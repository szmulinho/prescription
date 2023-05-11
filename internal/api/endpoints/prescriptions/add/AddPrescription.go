package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/prescription/database"
	"github.com/szmulinho/prescription/internal/model"
	"io/ioutil"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func CreatePrescription(w http.ResponseWriter, r *http.Request) {

	var newPresc model.CreatePrescInput
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &newPresc)
	if err != nil {
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
		if singlePresc.PreId == model.Prescription.PreId {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Prescription %model.already exist", model.Prescription.PreId)})
			return
		}
	}

	model.Prescs = append(model.Prescs, model.Prescription)

	fmt.Printf("created new prescription %+v\n", model.Prescription)
	log.Printf("%+v", model.Prescription)
	w.WriteHeader(http.StatusCreated)

	client := &http.Client{}
	reqBody, _ = json.Marshal(model.Prescription)
	req, _ := http.NewRequest("POST", "http://localhost:8081/presc", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)

	json.NewEncoder(w).Encode(model.Prescription)
}
