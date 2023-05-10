package get

import (
	"encoding/json"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func GetAllPrescriptions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Prescs)

}
