package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *handlers) UpdatePrescription(w http.ResponseWriter, r *http.Request) {
	prescIDStr := mux.Vars(r)["id"]
	prescID, err := strconv.ParseInt(prescIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedPresc model.CreatePrescInput

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the task title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedPresc)

	for i, singlePresc := range model.Prescs {
		if singlePresc.PreID == prescID {
			singlePresc.Drugs = updatedPresc.Drugs
			singlePresc.Expiration = updatedPresc.Expiration
			model.Prescs = append(model.Prescs[:i], singlePresc)
			json.NewEncoder(w).Encode(singlePresc)
		}
	}
}
