package get

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func GetPrescriptionsForPatient(w http.ResponseWriter, r *http.Request) {
	patientID := mux.Vars(r)["patient"]

	for _, singlePresc := range model.Prescs {
		if singlePresc.Patient == patientID {
			model.Prescs = append(model.Prescs, singlePresc)
		}
	}

	json.NewEncoder(w).Encode(model.Prescs)
}
