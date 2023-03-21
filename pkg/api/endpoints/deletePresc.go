package endpoints

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func DeletePrescription(w http.ResponseWriter, r *http.Request) {
	prescId := mux.Vars(r)["id"]

	prescs := []model.Presc{}
	for i, singlePresc := range prescs {
		if singlePresc.PreId == prescId {
			prescs = append(prescs[:i], prescs[i+1:]...)
			fmt.Fprintf(w, "The prescription with ID %v has been deleted successfully", prescId)
		}
	}
}
