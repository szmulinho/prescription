package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"io/ioutil"
	"net/http"
)

func UpdatePrescription(w http.ResponseWriter, r *http.Request) {
	prescPreId := mux.Vars(r)["id"]
	var updatedPresc model.Presc

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the task title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedPresc)

	prescs := []model.Presc{}
	for i, singlePresc := range prescs {
		if singlePresc.PreId == prescPreId {
			singlePresc.Drugs = updatedPresc.Drugs
			singlePresc.Expiration = updatedPresc.Expiration
			prescs = append(prescs[:i], singlePresc)
			json.NewEncoder(w).Encode(singlePresc)
		}
	}
}
