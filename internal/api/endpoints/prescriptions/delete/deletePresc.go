package delete

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func DeletePrescription(w http.ResponseWriter, r *http.Request) {
	var prescId = mux.Vars(r)["id"]
	for i, singlePresc := range model.Prescs {
		if singlePresc.PreId == prescId {
			model.Prescs = append(model.Prescs[:i], model.Prescs[i+1:]...)
			fmt.Fprintf(w, "The prescription with ID %v has been deleted successfully", prescId)
		}
	}
}
