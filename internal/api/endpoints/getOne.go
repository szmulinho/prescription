package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func GetOnePrescription(w http.ResponseWriter, r *http.Request) {
	prescPreId := mux.Vars(r)["id"]
	prescs := []model.Presc{}
	for _, singlePresc := range prescs {
		if singlePresc.PreId == prescPreId {
			json.NewEncoder(w).Encode(singlePresc)
		}
	}
}
