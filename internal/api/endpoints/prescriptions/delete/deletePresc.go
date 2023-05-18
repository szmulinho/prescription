package delete

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
	"strconv"
)

func DeletePrescription(w http.ResponseWriter, r *http.Request) {
	prescIDStr := mux.Vars(r)["id"]
	prescID, err := strconv.ParseInt(prescIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, singlePresc := range model.Prescs {
		if singlePresc.PreID == prescID {
			model.Prescs = append(model.Prescs[:i], model.Prescs[i+1:]...)
			fmt.Fprintf(w, "The prescription with ID %v has been deleted successfully", prescID)
		}
	}
}
