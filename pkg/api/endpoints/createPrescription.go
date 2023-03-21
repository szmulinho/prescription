package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/szmulinho/prescription/internal/model"
	"io/ioutil"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func CreatePrescription(w http.ResponseWriter, r *http.Request) {
	var in model.CreatePrescInput
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &in)
	if err != nil {
		panic(err)
		log.Printf("Invalid body")
	}
	prescs := []model.Presc{}
	for _, singlePresc := range prescs {
		fmt.Println(singlePresc)
		if singlePresc.PreId == in.PreId {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Prescription %model.already exist", in.PreId)})
			return
		}
	}
	drugs := make([]model.Drug, len(in.Drugs))
	for i, d := range in.Drugs {
		exist := false
		for _, existingDrugs := range model.Drugs {
			if d == existingDrugs.DrugID {
				drugs[i] = existingDrugs
				exist = true
			}
		}
		if !exist {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("drug %model.not exist", d)})
			return
		}
	}
	newPresc := model.Presc{
		PreId:      in.PreId,
		Drugs:      drugs,
		Expiration: in.Expiration,
	}

	prescs = append(prescs, newPresc)

	fmt.Printf("created new prescription %+v\n", in)
	log.Printf("%+v", in)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPresc)
}
