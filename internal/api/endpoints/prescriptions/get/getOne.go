package get

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
	"strconv"
)

func GetOnePrescription(w http.ResponseWriter, r *http.Request) {
	prescIDStr := mux.Vars(r)["id"]
	PreID, err := strconv.ParseInt(prescIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, singlePresc := range model.Prescs {
		if singlePresc.PreID == PreID {
			json.NewEncoder(w).Encode(singlePresc)
		}
	}
}
