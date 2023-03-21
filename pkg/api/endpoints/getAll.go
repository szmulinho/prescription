package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/prescription/pkg/model"
	"net/http"
)

func GetAllPrescriptions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Prescs)

}
